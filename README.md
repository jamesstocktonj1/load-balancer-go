# Go Database Load Balancer

A simple distributed database using a [load-balancer](https://github.com/jamesstocktonj1/load-balancer-go/tree/main/load-balancer) node and a number of [database](https://github.com/jamesstocktonj1/load-balancer-go/tree/main/database) nodes. This is an introduction to docker and docker compose.


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
The performance can be measured using the [db-test](https://github.com/jamesstocktonj1/load-balancer-go/tree/main/db-test) program. Results are as followed:

```
Starting Test
Creating Test Data
Data Size: 800

Test Create Value...
Minimum Response: 2.997000 ms
Maximum Response: 8.395000 ms
Average Response: 3.966000 ms

Test Get Value...
Minimum Response: 1.519000 ms
Maximum Response: 29.694000 ms
Average Response: 3.053000 ms

Test Set Value
Minimum Response: 2.952000 ms
Maximum Response: 47.825001 ms
Average Response: 4.608000 ms

Test Verify Value...
Minimum Response: 1.585000 ms
Maximum Response: 35.428001 ms
Average Response: 3.096000 ms
```