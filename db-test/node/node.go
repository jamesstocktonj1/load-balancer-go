package node

import (
	"net/http"
	"io"
	"encoding/json"
	"errors"
)

type Node struct {
	Address		string
	Port		string
	Status		string
	Keys		[]string
}

var Nodes = []Node{}

func ProcessData(data io.ReadCloser) map[string]string {
	defer data.Close()
	
	buf := map[string]string {}
	temp, _ := io.ReadAll(data)
	
	json.Unmarshal(temp, &buf)
	
	return buf
}

func (n *Node) ContainesKey(key string) bool {

	for _, k := range n.Keys {

		if k == key {
			return true
		}
	}
	return false
}

func (n *Node) UpdateKeys() error {

	resp, err := http.Get("http://" + n.Address + ":" + n.Port + "/dump")
	if err != nil {
		return err
	}

	data := ProcessData(resp.Body)
	
	n.Keys = []string {}
	for key, _ := range data {
		n.Keys = append(n.Keys, key)
	}

	return nil
}

func (n *Node) Ping() {
	resp, err := http.Get("http://" + n.Address + ":" + n.Port + "/ping")

	if err != nil {
		n.Status = "Offline"
	} else if resp.StatusCode == 200 {
		n.Status = "Online"
	} else {
		n.Status = "Error"
	}
}


func (n *Node) CreateValue(key string, value string) error {
	req, err := http.NewRequest("POST", "http://" + n.Address + ":" + n.Port + "/data/" + key + "?value=" + value, nil)
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	} else if resp.StatusCode != http.StatusCreated {
		return errors.New(resp.Status)
	}

	n.Keys = append(n.Keys, key)
	return nil
}

func (n *Node) SetValue(key string, value string) error {
	req, err := http.NewRequest("PUT", "http://" + n.Address + ":" + n.Port + "/data/" + key + "?value=" + value, nil)
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	} else if resp.StatusCode != http.StatusOK {
		return errors.New(resp.Status)
	}

	return nil
}

func (n *Node) GetValue(key string) (string, error) {

	resp, err := http.Get("http://" + n.Address + ":" + n.Port + "/data/" + key)
	if err != nil {
		return "", err
	}

	data := ProcessData(resp.Body)
	return data["value"], nil
}

func (n *Node) DeleteValue(key string) error {
	req, err := http.NewRequest("DELETE", "http://" + n.Address + ":" + n.Port + "/data/" + key, nil)
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	} else if resp.StatusCode != http.StatusOK {
		return errors.New(resp.Status)
	}

	return nil
}
