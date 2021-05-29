# Minimum example

## Log

- run envoy

```shell-session
koketani: ~/g/g/k/p/e/minimum (envoy ?)$ docker run --rm -i -t \
      -v $(pwd)/envoy.yaml:/envoy.yaml \
      -p 9901:9901 \
      -p 10000:10000 \
      --name envoy \
      envoyproxy/envoy-dev:latest \
          -c /envoy.yaml
[2020-12-14 14:24:43.264][7][info][main] [source/server/server.cc:305] initializing epoch 0 (base id=0, hot restart version=11.104)
[2020-12-14 14:24:43.265][7][info][main] [source/server/server.cc:307] statically linked extensions:
[2020-12-14 14:24:43.265][7][info][main] [source/server/server.cc:309]   envoy.bootstrap: envoy.bootstrap.wasm, envoy.extensions.network.socket_interface.default_socket_interface
[2020-12-14 14:24:43.265][7][info][main] [source/server/server.cc:309]   envoy.guarddog_actions: envoy.watchdog.abort_action, envoy.watchdog.profile_action
[2020-12-14 14:24:43.265][7][info][main] [source/server/server.cc:309]   envoy.compression.decompressor: envoy.compression.gzip.decompressor
[2020-12-14 14:24:43.265][7][info][main] [source/server/server.cc:309]   envoy.quic_client_codec: quiche
[2020-12-14 14:24:43.265][7][info][main] [source/server/server.cc:309]   envoy.udp_packet_writers: udp_default_writer, udp_gso_batch_writer
[2020-12-14 14:24:43.265][7][info][main] [source/server/server.cc:309]   envoy.dubbo_proxy.serializers: dubbo.hessian2
[2020-12-14 14:24:43.266][7][info][main] [source/server/server.cc:309]   envoy.http.cache: envoy.extensions.http.cache.simple
[2020-12-14 14:24:43.266][7][info][main] [source/server/server.cc:309]   envoy.compression.compressor: envoy.compression.gzip.compressor
[2020-12-14 14:24:43.266][7][info][main] [source/server/server.cc:309]   envoy.filters.listener: envoy.filters.listener.http_inspector, envoy.filters.listener.original_dst, envoy.filters.listener.original_src, envoy.filters.listener.proxy_protocol, envoy.filters.listener.tls_inspector, envoy.listener.http_inspector, envoy.listener.original_dst, envoy.listener.original_src, envoy.listener.proxy_protocol, envoy.listener.tls_inspector
[2020-12-14 14:24:43.266][7][info][main] [source/server/server.cc:309]   envoy.thrift_proxy.protocols: auto, binary, binary/non-strict, compact, twitter
[2020-12-14 14:24:43.267][7][info][main] [source/server/server.cc:309]   envoy.filters.network: envoy.client_ssl_auth, envoy.echo, envoy.ext_authz, envoy.filters.network.client_ssl_auth, envoy.filters.network.direct_response, envoy.filters.network.dubbo_proxy, envoy.filters.network.echo, envoy.filters.network.ext_authz, envoy.filters.network.http_connection_manager, envoy.filters.network.kafka_broker, envoy.filters.network.local_ratelimit, envoy.filters.network.mongo_proxy, envoy.filters.network.mysql_proxy, envoy.filters.network.postgres_proxy, envoy.filters.network.ratelimit, envoy.filters.network.rbac, envoy.filters.network.redis_proxy, envoy.filters.network.rocketmq_proxy, envoy.filters.network.sni_cluster, envoy.filters.network.sni_dynamic_forward_proxy, envoy.filters.network.tcp_proxy, envoy.filters.network.thrift_proxy, envoy.filters.network.wasm, envoy.filters.network.zookeeper_proxy, envoy.http_connection_manager, envoy.mongo_proxy, envoy.ratelimit, envoy.redis_proxy, envoy.tcp_proxy
[2020-12-14 14:24:43.267][7][info][main] [source/server/server.cc:309]   envoy.retry_host_predicates: envoy.retry_host_predicates.omit_canary_hosts, envoy.retry_host_predicates.omit_host_metadata, envoy.retry_host_predicates.previous_hosts
[2020-12-14 14:24:43.267][7][info][main] [source/server/server.cc:309]   envoy.internal_redirect_predicates: envoy.internal_redirect_predicates.allow_listed_routes, envoy.internal_redirect_predicates.previous_routes, envoy.internal_redirect_predicates.safe_cross_scheme
[2020-12-14 14:24:43.267][7][info][main] [source/server/server.cc:309]   envoy.dubbo_proxy.route_matchers: default
[2020-12-14 14:24:43.268][7][info][main] [source/server/server.cc:309]   envoy.dubbo_proxy.filters: envoy.filters.dubbo.router
[2020-12-14 14:24:43.275][7][info][main] [source/server/server.cc:309]   envoy.retry_priorities: envoy.retry_priorities.previous_priorities
[2020-12-14 14:24:43.276][7][info][main] [source/server/server.cc:309]   envoy.clusters: envoy.cluster.eds, envoy.cluster.logical_dns, envoy.cluster.original_dst, envoy.cluster.static, envoy.cluster.strict_dns, envoy.clusters.aggregate, envoy.clusters.dynamic_forward_proxy, envoy.clusters.redis
[2020-12-14 14:24:43.276][7][info][main] [source/server/server.cc:309]   envoy.quic_server_codec: quiche
[2020-12-14 14:24:43.276][7][info][main] [source/server/server.cc:309]   envoy.resource_monitors: envoy.resource_monitors.fixed_heap, envoy.resource_monitors.injected_resource
[2020-12-14 14:24:43.276][7][info][main] [source/server/server.cc:309]   envoy.filters.udp_listener: envoy.filters.udp.dns_filter, envoy.filters.udp_listener.udp_proxy
[2020-12-14 14:24:43.276][7][info][main] [source/server/server.cc:309]   envoy.transport_sockets.downstream: envoy.transport_sockets.alts, envoy.transport_sockets.quic, envoy.transport_sockets.raw_buffer, envoy.transport_sockets.tap, envoy.transport_sockets.tls, raw_buffer, tls
[2020-12-14 14:24:43.276][7][info][main] [source/server/server.cc:309]   envoy.stats_sinks: envoy.dog_statsd, envoy.metrics_service, envoy.stat_sinks.dog_statsd, envoy.stat_sinks.hystrix, envoy.stat_sinks.metrics_service, envoy.stat_sinks.statsd, envoy.stat_sinks.wasm, envoy.statsd
[2020-12-14 14:24:43.276][7][info][main] [source/server/server.cc:309]   envoy.tracers: envoy.dynamic.ot, envoy.lightstep, envoy.tracers.datadog, envoy.tracers.dynamic_ot, envoy.tracers.lightstep, envoy.tracers.opencensus, envoy.tracers.xray, envoy.tracers.zipkin, envoy.zipkin
[2020-12-14 14:24:43.278][7][info][main] [source/server/server.cc:309]   envoy.filters.http: envoy.buffer, envoy.cors, envoy.csrf, envoy.ext_authz, envoy.fault, envoy.filters.http.adaptive_concurrency, envoy.filters.http.admission_control, envoy.filters.http.aws_lambda, envoy.filters.http.aws_request_signing, envoy.filters.http.buffer, envoy.filters.http.cache, envoy.filters.http.cdn_loop, envoy.filters.http.compressor, envoy.filters.http.cors, envoy.filters.http.csrf, envoy.filters.http.decompressor, envoy.filters.http.dynamic_forward_proxy, envoy.filters.http.dynamo, envoy.filters.http.ext_authz, envoy.filters.http.fault, envoy.filters.http.grpc_http1_bridge, envoy.filters.http.grpc_http1_reverse_bridge, envoy.filters.http.grpc_json_transcoder, envoy.filters.http.grpc_stats, envoy.filters.http.grpc_web, envoy.filters.http.gzip, envoy.filters.http.header_to_metadata, envoy.filters.http.health_check, envoy.filters.http.ip_tagging, envoy.filters.http.jwt_authn, envoy.filters.http.local_ratelimit, envoy.filters.http.lua, envoy.filters.http.oauth2, envoy.filters.http.on_demand, envoy.filters.http.original_src, envoy.filters.http.ratelimit, envoy.filters.http.rbac, envoy.filters.http.router, envoy.filters.http.squash, envoy.filters.http.tap, envoy.filters.http.wasm, envoy.grpc_http1_bridge, envoy.grpc_json_transcoder, envoy.grpc_web, envoy.gzip, envoy.health_check, envoy.http_dynamo_filter, envoy.ip_tagging, envoy.local_rate_limit, envoy.lua, envoy.rate_limit, envoy.router, envoy.squash
[2020-12-14 14:24:43.278][7][info][main] [source/server/server.cc:309]   envoy.dubbo_proxy.protocols: dubbo
[2020-12-14 14:24:43.278][7][info][main] [source/server/server.cc:309]   envoy.udp_listeners: quiche_quic_listener, raw_udp_listener
[2020-12-14 14:24:43.278][7][info][main] [source/server/server.cc:309]   envoy.thrift_proxy.transports: auto, framed, header, unframed
[2020-12-14 14:24:43.278][7][info][main] [source/server/server.cc:309]   envoy.upstreams: envoy.filters.connection_pools.http.generic, envoy.filters.connection_pools.http.http, envoy.filters.connection_pools.http.tcp
[2020-12-14 14:24:43.279][7][info][main] [source/server/server.cc:309]   envoy.thrift_proxy.filters: envoy.filters.thrift.rate_limit, envoy.filters.thrift.router
[2020-12-14 14:24:43.279][7][info][main] [source/server/server.cc:309]   envoy.grpc_credentials: envoy.grpc_credentials.aws_iam, envoy.grpc_credentials.default, envoy.grpc_credentials.file_based_metadata
[2020-12-14 14:24:43.279][7][info][main] [source/server/server.cc:309]   envoy.resolvers: envoy.ip
[2020-12-14 14:24:43.279][7][info][main] [source/server/server.cc:309]   envoy.access_loggers: envoy.access_loggers.file, envoy.access_loggers.http_grpc, envoy.access_loggers.tcp_grpc, envoy.access_loggers.wasm, envoy.file_access_log, envoy.http_grpc_access_log, envoy.tcp_grpc_access_log, envoy.wasm_access_log
[2020-12-14 14:24:43.279][7][info][main] [source/server/server.cc:309]   envoy.health_checkers: envoy.health_checkers.redis
[2020-12-14 14:24:43.279][7][info][main] [source/server/server.cc:309]   envoy.transport_sockets.upstream: envoy.transport_sockets.alts, envoy.transport_sockets.quic, envoy.transport_sockets.raw_buffer, envoy.transport_sockets.tap, envoy.transport_sockets.tls, envoy.transport_sockets.upstream_proxy_protocol, raw_buffer, tls
[2020-12-14 14:24:43.325][7][info][main] [source/server/server.cc:325] HTTP header map info:
[2020-12-14 14:24:43.328][7][warning][runtime] [source/common/runtime/runtime_features.cc:31] Unable to use runtime singleton for feature envoy.http.headermap.lazy_map_min_size
[2020-12-14 14:24:43.328][7][warning][runtime] [source/common/runtime/runtime_features.cc:31] Unable to use runtime singleton for feature envoy.http.headermap.lazy_map_min_size
[2020-12-14 14:24:43.331][7][warning][runtime] [source/common/runtime/runtime_features.cc:31] Unable to use runtime singleton for feature envoy.http.headermap.lazy_map_min_size
[2020-12-14 14:24:43.331][7][warning][runtime] [source/common/runtime/runtime_features.cc:31] Unable to use runtime singleton for feature envoy.http.headermap.lazy_map_min_size
[2020-12-14 14:24:43.331][7][info][main] [source/server/server.cc:328]   request header map: 608 bytes: :authority,:method,:path,:protocol,:scheme,accept,accept-encoding,access-control-request-method,authorization,cache-control,cdn-loop,connection,content-encoding,content-length,content-type,expect,grpc-accept-encoding,grpc-timeout,if-match,if-modified-since,if-none-match,if-range,if-unmodified-since,keep-alive,origin,pragma,proxy-connection,referer,te,transfer-encoding,upgrade,user-agent,via,x-client-trace-id,x-envoy-attempt-count,x-envoy-decorator-operation,x-envoy-downstream-service-cluster,x-envoy-downstream-service-node,x-envoy-expected-rq-timeout-ms,x-envoy-external-address,x-envoy-force-trace,x-envoy-hedge-on-per-try-timeout,x-envoy-internal,x-envoy-ip-tags,x-envoy-max-retries,x-envoy-original-path,x-envoy-original-url,x-envoy-retriable-header-names,x-envoy-retriable-status-codes,x-envoy-retry-grpc-on,x-envoy-retry-on,x-envoy-upstream-alt-stat-name,x-envoy-upstream-rq-per-try-timeout-ms,x-envoy-upstream-rq-timeout-alt-response,x-envoy-upstream-rq-timeout-ms,x-forwarded-client-cert,x-forwarded-for,x-forwarded-proto,x-ot-span-context,x-request-id
[2020-12-14 14:24:43.332][7][info][main] [source/server/server.cc:328]   request trailer map: 128 bytes:
[2020-12-14 14:24:43.332][7][info][main] [source/server/server.cc:328]   response header map: 424 bytes: :status,access-control-allow-credentials,access-control-allow-headers,access-control-allow-methods,access-control-allow-origin,access-control-expose-headers,access-control-max-age,age,cache-control,connection,content-encoding,content-length,content-type,date,etag,expires,grpc-message,grpc-status,keep-alive,last-modified,location,proxy-connection,server,transfer-encoding,upgrade,vary,via,x-envoy-attempt-count,x-envoy-decorator-operation,x-envoy-degraded,x-envoy-immediate-health-check-fail,x-envoy-ratelimited,x-envoy-upstream-canary,x-envoy-upstream-healthchecked-cluster,x-envoy-upstream-service-time,x-request-id
[2020-12-14 14:24:43.332][7][info][main] [source/server/server.cc:328]   response trailer map: 152 bytes: grpc-message,grpc-status
[2020-12-14 14:24:43.342][7][info][main] [source/server/server.cc:448] admin address: 0.0.0.0:9901
[2020-12-14 14:24:43.346][7][info][main] [source/server/server.cc:583] runtime: layers:
  - name: base
    static_layer:
      {}
  - name: admin
    admin_layer:
      {}
[2020-12-14 14:24:43.348][7][info][config] [source/server/configuration_impl.cc:95] loading tracing configuration
[2020-12-14 14:24:43.348][7][info][config] [source/server/configuration_impl.cc:70] loading 0 static secret(s)
[2020-12-14 14:24:43.348][7][info][config] [source/server/configuration_impl.cc:76] loading 0 cluster(s)
[2020-12-14 14:24:43.351][7][info][config] [source/server/configuration_impl.cc:80] loading 1 listener(s)
[2020-12-14 14:24:43.368][7][info][config] [source/server/configuration_impl.cc:121] loading stats sink configuration
[2020-12-14 14:24:43.369][7][info][runtime] [source/common/runtime/runtime_impl.cc:425] RTDS has finished initialization
[2020-12-14 14:24:43.369][7][info][upstream] [source/common/upstream/cluster_manager_impl.cc:178] cm init: all clusters initialized
[2020-12-14 14:24:43.369][7][warning][main] [source/server/server.cc:565] there is no configured limit to the number of allowed active connections. Set a limit via the runtime key overload.global_downstream_max_connections
[2020-12-14 14:24:43.370][7][info][main] [source/server/server.cc:660] all clusters initialized. initializing init manager
[2020-12-14 14:24:43.371][7][info][config] [source/server/listener_manager_impl.cc:888] all dependencies initialized. starting workers
[2020-12-14 14:24:43.372][7][info][main] [source/server/server.cc:679] starting main dispatch loop
```

- request

```shell-session
koketani: ~/g/g/k/p/e/minimum (envoy ?)$ curl localhost:10000
ʕ◔ϖ◔ʔ% 
```
