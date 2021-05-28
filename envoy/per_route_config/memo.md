# Per Route Configuration

## Command

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
curl localhost:8000/docs
```

## Description

<https://github.com/envoyproxy/envoy/blob/1d1b708c7bf6efa02c41d9ce22cbf1e4a1aeec2c/docs/root/intro/arch_overview/advanced/data_sharing_between_filters.rst#http-per-route-filter-configuration>

```
In HTTP routes, :ref:`typed_per_filter_config <envoy_v3_api_field_config.route.v3.VirtualHost.typed_per_filter_config>` allows HTTP filters to have virtualhost/route-specific configuration in addition to a global filter config common to all virtual hosts. This configuration is converted and embedded into the route table. It is up to the HTTP filter implementation to treat the route-specific filter config as a replacement to global config or an enhancement. For example, the HTTP fault filter uses this technique to provide per-route fault configuration.

typed_per_filter_config is a map<string, google.protobuf.Any>. The Connection manager iterates over this map and invokes the filter factory interface createRouteSpecificFilterConfigTyped to parse/validate the struct value and convert it into a typed class object that’s stored with the route itself. HTTP filters can then query the route-specific filter config during request processing.
```
