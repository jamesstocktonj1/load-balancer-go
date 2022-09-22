package node

import (
	"io"
	"strings"
	"testing"
)

func TestProcessData(t *testing.T) {
	testData := "{\"test\": \"value\"}"
	expectedData := map[string]string{"test": "value"}

	testReader := strings.NewReader(testData)
	testReadClose := io.NopCloser(testReader)
	value := ProcessData(testReadClose)

	if value["test"] != expectedData["test"] {
		t.Errorf("expected ProcessedData to equal %s but got %s", expectedData["test"], value["test"])
	}
}

func TestContainesKey(t *testing.T) {

	t.Run("present Key", func(t *testing.T) {
		testNode := Node{}
		testNode.Keys = append(testNode.Keys, "test")

		present := testNode.ContainesKey("test")

		if present != true {
			t.Errorf("expected Node Keys to contain \"test\" but presence was false")
		}
	})

	t.Run("not present Key", func(t *testing.T) {
		testNode := Node{}

		present := testNode.ContainesKey("test")

		if present != false {
			t.Errorf("expected Node Keys to not contain \"test\" but presence was true")
		}
	})
}