# Bandwidth

<https://github.com/envoyproxy/envoy/pull/16358/files>

<https://www.envoyproxy.io/docs/envoy/latest/api-v3/extensions/filters/http/bandwidth_limit/v3alpha/bandwidth_limit.proto>

<https://www.envoyproxy.io/docs/envoy/latest/configuration/http/http_filters/bandwidth_limit_filter>

## Step

```shell-session
$ mkfile 100m image.bin
```

```shell-session
$ docker-compose up
```

```shell-session
$ docker build -t bandwidth/hey github.com/rakyll/hey
$ docker run --net bandwidth_default -v $(pwd)/image.bin:image.bin --rm bandwidth/hey -m PUT -D image.bin -t 0 -n 3 -c 3 http://envoy:10000/prefix
```
