# Envoy OPA

## Ref

- [ext_authz in sandboxes](https://www.envoyproxy.io/docs/envoy/latest/start/sandboxes/ext_authz)
- [ext_authz](https://www.envoyproxy.io/docs/envoy/latest/api-v3/extensions/filters/http/ext_authz/v3/ext_authz.proto#envoy-v3-api-msg-extensions-filters-http-ext-authz-v3-extauthz)

- [opa](https://www.openpolicyagent.org/docs/latest/)
- [opa-envoy-plugin](https://github.com/open-policy-agent/opa-envoy-plugin)

## Policy

```
koketani: ~/g/g/k/p/e/opa (envoy ?)$ cat policy.rego
package envoy.authz

import input.attributes.request.http as http_request

default allow = false

allow = response {
  http_request.method == "GET"
  response := {
    "allowed": true,
    "headers": {"x-current-user": "OPA"}
  }
}
```

## Log

```
koketani: ~/g/g/k/p/e/opa (envoy ?)$ curl localhost:8000 --verbose
*   Trying ::1...
* TCP_NODELAY set
* Connected to localhost (::1) port 8000 (#0)
> GET / HTTP/1.1
> Host: localhost:8000
> User-Agent: curl/7.64.1
> Accept: */*
>
< HTTP/1.1 404 Not Found
< content-type: text/html; charset=utf-8
< content-length: 232
< server: envoy
< date: Tue, 15 Dec 2020 05:51:23 GMT
< x-envoy-upstream-service-time: 3
<
<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 3.2 Final//EN">
<title>404 Not Found</title>
<h1>Not Found</h1>
<p>The requested URL was not found on the server. If you entered the URL manually please check your spelling and try again.</p>
* Connection #0 to host localhost left intact
* Closing connection 0
koketani: ~/g/g/k/p/e/opa (envoy ?)$ curl localhost:8000/service --verbose
*   Trying ::1...
* TCP_NODELAY set
* Connected to localhost (::1) port 8000 (#0)
> GET /service HTTP/1.1
> Host: localhost:8000
> User-Agent: curl/7.64.1
> Accept: */*
>
< HTTP/1.1 200 OK
< content-type: text/html; charset=utf-8
< content-length: 28
< server: envoy
< date: Tue, 15 Dec 2020 05:51:28 GMT
< x-envoy-upstream-service-time: 2
<
* Connection #0 to host localhost left intact
Hello OPA from behind Envoy!* Closing connection 0
koketani: ~/g/g/k/p/e/opa (envoy ?)$ curl -X POST -d hoge=fuaga localhost:8000/service --verbose
Note: Unnecessary use of -X or --request, POST is already inferred.
*   Trying ::1...
* TCP_NODELAY set
* Connected to localhost (::1) port 8000 (#0)
> POST /service HTTP/1.1
> Host: localhost:8000
> User-Agent: curl/7.64.1
> Accept: */*
> Content-Length: 10
> Content-Type: application/x-www-form-urlencoded
>
* upload completely sent off: 10 out of 10 bytes
< HTTP/1.1 405 Method Not Allowed
< content-type: text/html; charset=utf-8
< allow: OPTIONS, GET, HEAD
< content-length: 178
< server: envoy
< date: Tue, 15 Dec 2020 05:52:04 GMT
< x-envoy-upstream-service-time: 1
<
<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 3.2 Final//EN">
<title>405 Method Not Allowed</title>
<h1>Method Not Allowed</h1>
<p>The method is not allowed for the requested URL.</p>
* Connection #0 to host localhost left intact
* Closing connection 0
```
