# Bandwidth

<https://github.com/envoyproxy/envoy/pull/16358/files>

<https://www.envoyproxy.io/docs/envoy/latest/api-v3/extensions/filters/http/bandwidth_limit/v3alpha/bandwidth_limit.proto>

<https://www.envoyproxy.io/docs/envoy/latest/configuration/http/http_filters/bandwidth_limit_filter>

## Step

```shell-session
$ mkfile 1g image.bin
```

```shell-session
$ docker-compose up
```
 
```shell-session
$ docker build -t bandwidth/hey github.com/rakyll/hey
```

- limited access

```shell-session
$ docker run --net bandwidth_default -v $(pwd)/image.bin:/image.bin --rm bandwidth/hey -m PUT -D /image.bin -t 0 -n 3 -c 3 http://envoy:10000/protected
```

- unlimited access

```shell-session
$ docker run --net bandwidth_default -v $(pwd)/image.bin:/image.bin --rm bandwidth/hey -m PUT -D /image.bin -t 0 -n 3 -c 3 http://envoy:10000/unprotected
```

## Logs

- limited (60mbps, 3 sessions)

hey

```
koketani: ~/g/g/k/p/e/bandwidth (envoy !)$ docker run --net bandwidth_default -v $(pwd)/image.bin:/image.bin --rm bandwidth/hey -m PUT -D /image.bin -t 0 -n 3 -c 3 http://envoy:10000/protected

Summary:
  Total:        53.2688 secs
  Slowest:      53.2648 secs
  Fastest:      51.8716 secs
  Average:      52.6026 secs
  Requests/sec: 0.0563


Response time histogram:
  51.872 [1]    |■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  52.011 [0]    |
  52.150 [0]    |
  52.290 [0]    |
  52.429 [0]    |
  52.568 [0]    |
  52.708 [1]    |■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  52.847 [0]    |
  52.986 [0]    |
  53.125 [0]    |
  53.265 [1]    |■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■


Latency distribution:
  10% in 52.6714 secs
  25% in 53.2648 secs
  0% in 0.0000 secs
  0% in 0.0000 secs
  0% in 0.0000 secs
  0% in 0.0000 secs
  0% in 0.0000 secs

Details (average, fastest, slowest):
  DNS+dialup:   0.0275 secs, 0.0261 secs, 0.0288 secs
  DNS-lookup:   0.0241 secs, 0.0237 secs, 0.0243 secs
  req write:    52.2422 secs, 51.2683 secs, 53.1378 secs
  resp wait:    0.3260 secs, 0.0936 secs, 0.5675 secs
  resp read:    0.0011 secs, 0.0001 secs, 0.0031 secs

Status code distribution:
  [200] 3 responses
```

server

```
server    | 2021/06/14 22:13:26 the number of available cpus is 4
server    | 2021/06/14 22:14:19 new session from 172.30.0.3:47222
server    | 2021/06/14 22:14:19 new session from 172.30.0.3:47220
server    | 2021/06/14 22:14:19 new session from 172.30.0.3:47224
server    | 2021/06/14 22:15:11 the total upload size is 1023mb from 172.30.0.3:47224
server    | 2021/06/14 22:15:12 the total upload size is 1023mb from 172.30.0.3:47220
server    | 2021/06/14 22:15:12 the total upload size is 1023mb from 172.30.0.3:47222
```

- limited (60mbps, 1 sessions)

hey

```
koketani: ~/g/g/k/p/e/bandwidth (envoy !)$ docker run --net bandwidth_default -v $(pwd)/image.bin:/image.bin --rm bandwidth/hey -m PUT -D /image.bin -t 0 -n 1 -c 1 http://envoy:10000/protected

Summary:
  Total:        22.6910 secs
  Slowest:      22.6868 secs
  Fastest:      22.6868 secs
  Average:      22.6868 secs
  Requests/sec: 0.0441


Response time histogram:
  22.687 [1]    |■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  22.687 [0]    |
  22.687 [0]    |
  22.687 [0]    |
  22.687 [0]    |
  22.687 [0]    |
  22.687 [0]    |
  22.687 [0]    |
  22.687 [0]    |
  22.687 [0]    |
  22.687 [0]    |


Latency distribution:
  0% in 0.0000 secs
  0% in 0.0000 secs
  0% in 0.0000 secs
  0% in 0.0000 secs
  0% in 0.0000 secs
  0% in 0.0000 secs
  0% in 0.0000 secs

Details (average, fastest, slowest):
  DNS+dialup:   0.0333 secs, 0.0333 secs, 0.0333 secs
  DNS-lookup:   0.0254 secs, 0.0254 secs, 0.0254 secs
  req write:    22.5375 secs, 22.5375 secs, 22.5375 secs
  resp wait:    0.1072 secs, 0.1072 secs, 0.1072 secs
  resp read:    0.0006 secs, 0.0006 secs, 0.0006 secs

Status code distribution:
  [200] 1 responses
```

server

```
server    | 2021/06/14 22:19:55 new session from 172.30.0.3:47220
server    | 2021/06/14 22:20:18 the total upload size is 1023mb from 172.30.0.3:47220
```


- unlimited (3 sessions)

hey

```
koketani: ~/g/g/k/p/e/bandwidth (envoy !)$ docker run --net bandwidth_default -v $(pwd)/image.bin:/image.bin --rm bandwidth/hey -m PUT -D /image.bin -t 0 -n 3 -c 3 http://envoy:10000/unprotected

Summary:
  Total:        36.9831 secs
  Slowest:      36.9760 secs
  Fastest:      35.7403 secs
  Average:      36.2992 secs
  Requests/sec: 0.0811


Response time histogram:
  35.740 [1]    |■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  35.864 [0]    |
  35.987 [0]    |
  36.111 [0]    |
  36.235 [1]    |■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  36.358 [0]    |
  36.482 [0]    |
  36.605 [0]    |
  36.729 [0]    |
  36.852 [0]    |
  36.976 [1]    |■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■


Latency distribution:
  10% in 36.1811 secs
  25% in 36.9760 secs
  0% in 0.0000 secs
  0% in 0.0000 secs
  0% in 0.0000 secs
  0% in 0.0000 secs
  0% in 0.0000 secs

Details (average, fastest, slowest):
  DNS+dialup:   0.0163 secs, 0.0135 secs, 0.0183 secs
  DNS-lookup:   0.0079 secs, 0.0021 secs, 0.0136 secs
  req write:    36.2585 secs, 35.6818 secs, 36.9437 secs
  resp wait:    0.0123 secs, 0.0028 secs, 0.0310 secs
  resp read:    0.0013 secs, 0.0002 secs, 0.0034 secs

Status code distribution:
  [200] 3 responses
```

server

```
server    | 2021/06/14 22:22:27 new session from 172.30.0.3:50368
server    | 2021/06/14 22:22:27 new session from 172.30.0.3:50372
server    | 2021/06/14 22:22:27 new session from 172.30.0.3:50992
server    | 2021/06/14 22:23:03 the total upload size is 1023mb from 172.30.0.3:50992
server    | 2021/06/14 22:23:03 the total upload size is 1023mb from 172.30.0.3:50368
server    | 2021/06/14 22:23:04 the total upload size is 1023mb from 172.30.0.3:50372
```
