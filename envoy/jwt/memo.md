# JWT example

## Log

- run envoy

```shell-session
koketani: ~/g/g/k/p/e/jwt (envoy ?)$ docker run --rm -i -t \
      -v $(pwd)/envoy.yaml:/envoy.yaml \
      -p 9901:9901 \
      -p 10000:10000 \
      --name envoy \
      envoyproxy/envoy-dev:latest \
          -c /envoy.yaml
```

- request

```shell-session
koketani: ~/g/g/k/p/e/jwt (envoy ?)$ curl -i localhost:10000
HTTP/1.1 401 Unauthorized
content-length: 14
content-type: text/plain
date: Mon, 14 Dec 2020 14:43:23 GMT
server: envoy

Jwt is missing% 
```

```shell-session
koketani: ~/g/g/k/p/e/jwt (envoy ?)$ curl -i -H 'Authorization: Bearer 123' localhost:10000
HTTP/1.1 401 Unauthorized
content-length: 79
content-type: text/plain
date: Mon, 14 Dec 2020 14:47:32 GMT
server: envoy

Jwt is not in the form of Header.Payload.Signature with two dots and 3 sections% 
```

```shell-session
koketani: ~/g/g/k/p/e/jwt (envoy ?)$ export ID_TOKEN=$(gcloud config config-helper --format 'value(credential.id_token)')
koketani: ~/g/g/k/p/e/jwt (envoy ?)$ # OR gcloud auth print-identity-token
koketani: ~/g/g/k/p/e/jwt (envoy ?)$ curl -i -H "Authorization: Bearer $ID_TOKEN" localhost:10000
HTTP/1.1 200 OK
content-length: 12
content-type: text/plain
date: Mon, 14 Dec 2020 14:48:33 GMT
server: envoy

ʕ◔ϖ◔ʔ% 
```
