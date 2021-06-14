# Bandwidth

<https://github.com/envoyproxy/envoy/pull/16358/files>

<https://www.envoyproxy.io/docs/envoy/latest/api-v3/extensions/filters/http/bandwidth_limit/v3alpha/bandwidth_limit.proto>

<https://www.envoyproxy.io/docs/envoy/latest/configuration/http/http_filters/bandwidth_limit_filter>

## Step

```
docker run --rm -i -t \
      -v $(pwd)/envoy.yaml:/envoy.yaml \
      -p 8000:8000 \
      --name envoy \
      envoyproxy/envoy-dev:latest \
          -l debug \
          -c /envoy.yaml
```

```
$ go get github.com/rakyll/hey
```
