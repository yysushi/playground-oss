# Minimum example

## Log

- run envoy

```shell-session
koketani: ~/g/g/k/p/e/logging (envoy ?)$ docker run --rm -i -t \
      -v $(pwd)/access.log:/access.log -v $(pwd)/envoy.yaml:/envoy.yaml \
      -p 9901:9901 \
      -p 10000:10000 \
      --name envoy \
      envoyproxy/envoy-dev:latest \
          -c /envoy.yaml
```

- request

```shell-session
koketani: ~/g/g/k/p/e/minimum (envoy ?)$ curl localhost:10000
ʕ◔ϖ◔ʔ% 
```

```shell-session
koketani: ~/g/g/k/p/e/logging (envoy ?)$ tail -f access.log
[2020-12-14T16:46:08.303Z] "GET / HTTP/1.1" 200 - 0 12 0 - "-" "curl/7.64.1" "484f5870-8d59-42b1-b799-b510289af64a" "localhost:10000" "-"
```
