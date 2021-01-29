# Quick start

## Run envoy

- run

```shell-session
docker run --rm -i -t \
      -v $(pwd)/envoy-demo.yaml:/envoy-demo.yaml \
      -p 9901:9901 \
      -p 10000:10000 \
      --name envoy \
      envoyproxy/envoy-dev:latest \
          -c /envoy-demo.yaml
```

- confirm

```shell-session
curl localhost:9901
curl localhost:10000
```

## Static configuration

need to specify listeners and clusters as static_resources

an admin section if you wish to monitor Envoy or retrieve stats

- static_resources

everything that is configured statically when Envoy starts, as opposed to dynamically at runtime

```
static_resources:
```

- listeners

listen on port 10000 and serves all paths are matched and routed to the `service_envoyproxy_io` cluster

```
static_resources:

  listeners:
  - name: listener_0
    address:
      socket_address:
        address: 0.0.0.0
        port_value: 10000
    filter_chains:
    - filters:
      - name: envoy.filters.network.http_connection_manager
        typed_config:
          "@type": type.googleapis.com/envoy.extensions.filters.network.http_connection_manager.v3.HttpConnectionManager
          stat_prefix: ingress_http
          http_filters:
          - name: envoy.filters.http.router
          route_config:
            name: local_route
            virtual_hosts:
            - name: local_service
              domains: ["*"]
              routes:
              - match:
                  prefix: "/"
                route:
                  host_rewrite_literal: www.envoyproxy.io
                  cluster: service_envoyproxy_io
```

- clusters

the cluster proxies over TLS to https://www.envoyproxy.io.

```
  clusters:
  - name: service_envoyproxy_io
    connect_timeout: 30s
    type: LOGICAL_DNS
    # Comment out the following line to test on v6 networks
    dns_lookup_family: V4_ONLY
    load_assignment:
      cluster_name: service_envoyproxy_io
      endpoints:
      - lb_endpoints:
        - endpoint:
            address:
              socket_address:
                address: www.envoyproxy.io
                port_value: 443
    transport_socket:
      name: envoy.transport_sockets.tls
      typed_config:
        "@type": type.googleapis.com/envoy.extensions.transport_sockets.tls.v3.UpstreamTlsContext
        sni: www.envoyproxy.io
```

- admin

required to enable and configure the administration server

```
admin:
  access_log_path: /dev/null
  address:
    socket_address:
      address: 0.0.0.0
      port_value: 9901
```

## Dynamic configuration from filesystem

## Dynamic configuration from control plane
