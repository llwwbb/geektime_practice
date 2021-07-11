init

```shell
docker network create redis-net
docker run -d --name redis-server --network redis-net redis:alpine
```

> 使用 redis benchmark 工具, 测试 10 20 50 100 200 1k 5k 字节 value 大小，redis get set 性能。

```shell
docker run -it --network redis-net --rm redis:alpine redis-benchmark -h redis-server -d 10 -t get,set
docker run -it --network redis-net --rm redis:alpine redis-benchmark -h redis-server -d 20 -t get,set
docker run -it --network redis-net --rm redis:alpine redis-benchmark -h redis-server -d 50 -t get,set
docker run -it --network redis-net --rm redis:alpine redis-benchmark -h redis-server -d 100 -t get,set
docker run -it --network redis-net --rm redis:alpine redis-benchmark -h redis-server -d 200 -t get,set
docker run -it --network redis-net --rm redis:alpine redis-benchmark -h redis-server -d 1000 -t get,set
docker run -it --network redis-net --rm redis:alpine redis-benchmark -h redis-server -d 5000 -t get,set
docker run -it --network redis-net --rm redis:alpine redis-benchmark -h redis-server -d 10000 -t get,set
```

- requests per second

| size  |   set    |    get    |
| :---: | :------: | :-------: |
|  10   | 77700.08 | 71787.51  |
|  20   | 72098.05 | 68775.79  |
|  50   | 81499.59 | 75757.57  |
|  100  | 74906.37 | 70972.32  |
|  200  | 75757.57 | 88495.58  |
| 1000  | 74128.98 | 92336.11  |
| 5000  | 36153.29 | 114547.53 |
| 10000 | 30609.12 | 54171.18  |

- set latency summary (msec)

| size  |  avg  |  min  |  p50  |  p95  |  p99  |  max  |
| :---: | :---: | :---: | :---: | :---: | :---: | :---: |
|  10   | 0.332 | 0.088 | 0.319 | 0.487 | 0.575 | 1.391 |
|  20   | 0.354 | 0.080 | 0.343 | 0.511 | 0.599 | 1.239 |
|  50   | 0.316 | 0.056 | 0.303 | 0.455 | 0.543 | 1.519 |
|  100  | 0.342 | 0.072 | 0.327 | 0.487 | 0.591 | 1.375 |
|  200  | 0.338 | 0.080 | 0.327 | 0.471 | 0.567 | 1.135 |
| 1000  | 0.345 | 0.072 | 0.335 | 0.503 | 0.599 | 1.359 |
| 5000  | 0.705 | 0.072 | 0.687 | 1.063 | 1.263 | 1.687 |
| 10000 | 0.833 | 0.072 | 0.823 | 1.175 | 1.359 | 1.991 |

- get latency summary (msec)

| size  |  avg  |  min  |  p50  |  p95  |  p99  |  max  |
| :---: | :---: | :---: | :---: | :---: | :---: | :---: |
|  10   | 0.355 | 0.064 | 0.351 | 0.487 | 0.567 | 1.215 |
|  20   | 0.370 | 0.064 | 0.367 | 0.543 | 0.623 | 1.503 |
|  50   | 0.339 | 0.072 | 0.327 | 0.487 | 0.607 | 1.839 |
|  100  | 0.359 | 0.096 | 0.351 | 0.503 | 0.599 | 1.463 |
|  200  | 0.293 | 0.048 | 0.287 | 0.471 | 0.567 | 1.431 |
| 1000  | 0.281 | 0.056 | 0.279 | 0.359 | 0.487 | 1.279 |
| 5000  | 0.260 | 0.096 | 0.231 | 0.479 | 0.631 | 1.271 |
| 10000 | 0.471 | 0.176 | 0.463 | 0.631 | 0.719 | 1.743 |

> 写入一定量的 kv 数据, 根据数据大小 1w-50w 自己评估, 结合写入前后的 info memory 信息 , 分析上述不同 value 大小下，平均每个 key 的占用内存空间。

```shell
docker build -t week8 .
docker run -it --network redis-net --rm week8
```

向redis插入200000条kv，key长度15，内存占用情况如下

memory uses | size | before | after | actual | per | about | | :---: | :----: | :--------: | :--------: | :---------:
| :---: | | 10 | 67864 | 8047352 | 7979488 | 39.89744 | 40 | | 20 | 47504 | 9647504 | 9600000 | 48 | 48 | | 50 | 47672 |
16047656 | 15999984 | 79.99992 | 80 | | 100 | 47824 | 27247808 | 27199984 | 135.99992 | 136 | | 200 | 47960 | 49647960 |
49600000 | 248 | 248 | | 1000 | 48128 | 209648112 | 209599984 | 1047.99992 | 1048 | | 5000 | 48280 | 1028848264 |
1028799984 | 5143.99992 | 5144 | | 10000 | 48432 | 2052848416 | 2052799984 | 10263.99992 | 10264 |
