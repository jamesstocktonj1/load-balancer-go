# Go Database Load Balancer

A simple distributed database using a load-balancer node and a number of database nodes. This is an introduction to docker and docker compose.


To run the project simply run docker compose: 
```
docker compose up -d --build
```

The architecture is as followed:

| Docker Name | Port |
| --- | --- |
| load-balancer | 3000 |
| database1 | 3001  |
| database2 | 3002  |