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

## Performance
The performance can be measured using the db-test/ program. Results are as followed:

```
Starting Test
Creating Test Data
Data Size: 5000

Test Create Value...
Create Value Response Time: 4.010917 ms

Test Get Value...
Get Value Response Time: 15.587667 ms

Test Set Value
Set Value Response Time: 15.461544 ms

Test Verify Value...
Get Value Response Time: 15.599510 ms
```