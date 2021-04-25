# TODO

- dynamic forward

  - dynamic forward proxy
    <https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/http/http_proxy#arch-overview-http-dynamic-forward-proxy>
  - header to metadata
    <https://www.envoyproxy.io/docs/envoy/latest/configuration/http/http_filters/header_to_metadata_filter>

    `%DYNAMIC_METADATA([“namespace”, “key”, …])%`

    <https://www.envoyproxy.io/docs/envoy/latest/configuration/http/http_conn_man/headers.html?highlight=dynamic_metadata>

    `host_rewrite_header`

    <https://www.envoyproxy.io/docs/envoy/latest/api-v3/config/route/v3/route_components.proto#envoy-v3-api-field-config-route-v3-routeaction-host-rewrite-header>

    `request_headers_to_add`

    `regex_rewrite` or `host_rewrite_path_regex`

- request flow

  <https://www.envoyproxy.io/docs/envoy/latest/intro/life_of_a_request>

  ```
  The request first passes through CustomFilter which may read and modify the request. The most important HTTP filter is the router filter which sits at the end of the HTTP filter chain. When decodeHeaders is invoked on the router filter, the route is selected and a cluster is picked. The request headers on the stream are forwarded to an upstream endpoint in that cluster.
  ```

- ext authz

- tracing

- fault injection

- websocket

- udp

- dynamic configuration

https://github.com/envoyproxy/envoy/tree/main/examples

- lua
