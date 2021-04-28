# Envoy

## Terminology

terminologies

- host
- downstream
- upstream
- listener
- cluster
- mesh
- runtime configuration

two high level layers

- `user <-> service`
- `user <-> proxy <-> service`: they communicate with each other as hosts

two directions

- downstream: `user <-> proxy`
- upstream: `proxy <-> service`

two subsystems

- listener: proxy provides listener, how user connects to it
- cluster: hosts in service are logically grouped as cluster

## Threading model

a single process with multiple threads architecture

- a single primary thread: controls various sporadic coordination tasks
- some number of worker threads: perform listening, filtering, and forwarding

## High level architecture

- Listener subsystem: handles downstream request processing and be responsible for managing the downstream request lifecycle and for the response path to the client
- Cluster subsystem: be responsible for selecting and configuring the upstream connection to an endpoint by being knowledgeable about cluster and endpoint health, load balancing and connection pooling exists

![architecture](https://www.envoyproxy.io/docs/envoy/latest/_images/lor-architecture.svg)

## Configuration overview

two types of configurations

static and dynamic

- example

```yaml
static_resources:
  listeners:
  # There is a single listener bound to port 443.
  - name: listener_https
    address:
      socket_address:
        protocol: TCP
        address: 0.0.0.0
        port_value: 443
    # A single listener filter exists for TLS inspector.
    listener_filters:
    - name: "envoy.filters.listener.tls_inspector"
      typed_config: {}
    # On the listener, there is a single filter chain that matches SNI for acme.com.
    filter_chains:
    - filter_chain_match:
        # This will match the SNI extracted by the TLS Inspector filter.
        server_names: ["acme.com"]
      # Downstream TLS configuration.
      transport_socket:
        name: envoy.transport_sockets.tls
        typed_config:
          "@type": type.googleapis.com/envoy.extensions.transport_sockets.tls.v3.DownstreamTlsContext
          common_tls_context:
            tls_certificates:
            - certificate_chain: { filename: "certs/servercert.pem" }
              private_key: { filename: "certs/serverkey.pem" }
      filters:
      # The HTTP connection manager is the only network filter.
      - name: envoy.filters.network.http_connection_manager
        typed_config:
          "@type": type.googleapis.com/envoy.extensions.filters.network.http_connection_manager.v3.HttpConnectionManager
          stat_prefix: ingress_http
          use_remote_address: true
          http2_protocol_options:
            max_concurrent_streams: 100
          # File system based access logging.
          access_log:
            - name: envoy.access_loggers.file
              typed_config:
                "@type": type.googleapis.com/envoy.extensions.access_loggers.file.v3.FileAccessLog
                path: "/var/log/envoy/access.log"
          # The route table, mapping /foo to some_service.
          route_config:
            name: local_route
            virtual_hosts:
            - name: local_service
              domains: ["acme.com"]
              routes:
              - match:
                  path: "/foo"
                route:
                  cluster: some_service
          # CustomFilter and the HTTP router filter are the HTTP filter chain.
          http_filters:
            # - name: some.customer.filter
            - name: envoy.filters.http.router
  clusters:
  - name: some_service
    connect_timeout: 5s
    # Upstream TLS configuration.
    transport_socket:
      name: envoy.transport_sockets.tls
      typed_config:
        "@type": type.googleapis.com/envoy.extensions.transport_sockets.tls.v3.UpstreamTlsContext
    load_assignment:
      cluster_name: some_service
      # Static endpoint assignment.
      endpoints:
      - lb_endpoints:
        - endpoint:
            address:
              socket_address:
                address: 10.1.2.10
                port_value: 10002
        - endpoint:
            address:
              socket_address:
                address: 10.1.2.11
                port_value: 10002
    http2_protocol_options:
      max_concurrent_streams: 100
  - name: some_statsd_sink
    connect_timeout: 5s
    # The rest of the configuration for statsd sink cluster.
# statsd sink.
stats_sinks:
  - name: envoy.stat_sinks.statsd
    typed_config:
      "@type": type.googleapis.com/envoy.config.metrics.v3.StatsdSink
      tcp_cluster_name: some_statsd_sink
```

- schema

```yaml
- static_resources
  - listeners
    - name
    - address
      - socket_address
        - protocol
        - address
        - port_value
    - listener_filters
      - name
      - typed_config
    - filter_chains
      - filter_chain_match
        - server_names
      - transport_socket
        - name
        - typed_config
          - @type
          - common_tls_context
            - tls_certificates
              - certificate_chain
                - filename
              - private_key
                - filename
      - filters
        - name
        - typed_config
          - @type
          - stat_prefix
          - use_remote_address
          - http2_protocol_options
            - max_concurrent_streams
          - access_log
            - name
            - typed_config
              - @type
              - path
          - route_config
            - name
            - virtual_hosts
              - name
              - domains
              - routes
                - match
                  - path
                - route
                  - cluster
          - http_filters
            - name
  - clusters
    - name
    - connect_timeout
    - transport_socket
      - name
      - typed_config
        - @type
    - load_assignment
      - cluster_name
      - endpoints
        - lb_endpoints
          - endpoint
            - address
              - socket_address
                - address
                - port_value
    - http2_protocol_options
      - max_concurrent_streams
- stats_sinks
  - name
  - typed_config
    - @type
    - tcp_cluster_name
```

## Network topology

1. service mesh sidecar proxy: gateway to the network

ingress: user -> envoy
egress: envoy <- service

2. load balancer in service mesh
3. ingress/egress proxy on the network edge
4. double proxy

## Deployment types

- Service to service only
- Service to service plus front proxy
- Service to service, front proxy, and double proxy

## Request flow

https://www.envoyproxy.io/docs/envoy/latest/intro/life_of_a_request#request-flow

```
6. For each HTTP stream, an HTTP filter chain is created and runs. The request first passes through CustomFilter which may read and modify the request. The most important HTTP filter is the router filter which sits at the end of the HTTP filter chain. When decodeHeaders is invoked on the router filter, the route is selected and a cluster is picked. The request headers on the stream are forwarded to an upstream endpoint in that cluster. The router filter obtains an HTTP connection pool from the cluster manager for the matched cluster to do this.
10. The request, consisting of headers, and optional body and trailers, is proxied upstream, and the response is proxied downstream. The response passes through the HTTP filters in the opposite order from the request, starting at the router filter and passing through CustomFilter, before being sent downstream.
```

https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/http/http_filters
