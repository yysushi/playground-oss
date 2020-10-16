# Envoy

## Log

```
docker pull envoyproxy/envoy-alpine
```

`envoyproxy/envoy-alpine`

## Documentation

### What is?

- Out of process architecture
- L3/L4 filter architecture
- HTTP L7 filter architecture
- First class HTTP/2 support
- HTTP L7 routing
- gRPC support
- Service discovery and dynamic configuration
- Health checking
- Advanced load balancing
- Front/edge proxy support
- Best in class observability

### Architecture

#### Terminology

#### Threading model

### Life of a Request

#### Terminology

- Cluster: a logical service with a set of endpoints that Envoy forwards requests to.
- Downstream: an entity connecting to Envoy. This may be a local application (in a sidecar model) or a network node. In non-sidecar models, this is a remote client.
- Endpoints: network nodes that implement a logical service. They are grouped into clusters. Endpoints in a cluster are upstream of an Envoy proxy.
- Filter: a module in the connection or request processing pipeline providing some aspect of request handling. An analogy from Unix is the composition of small utilities (filters) with Unix pipes (filter chains).
- Filter chain: a series of filters.
- Listeners: Envoy module responsible for binding to an IP/port, accepting new TCP connections (or UDP datagrams) and orchestrating the downstream facing aspects of request processing.
- Upstream: an endpoint (network node) that Envoy connects to when forwarding requests for a service. This may be a local application (in a sidecar model) or a network node. In non-sidecar models, this corresponds with a remote backend.

#### Network topology

#### Configuration

- L3/4 protocol, e.g. TCP, UDP, Unix domain sockets.
- L7 protocol, e.g. HTTP/1, HTTP/2, HTTP/3, gRPC, Thrift, Dubbo, Kafka, Redis and various databases.
- Transport socket, e.g. plain text, TLS, ALTS.
- Connection routing, e.g. PROXY protocol, original destination, dynamic forwarding.
- Authentication and authorization.
- Circuit breakers and outlier detection configuration and activation state.
- Many other configurations for networking, HTTP, listener, access logging, health checking, tracing and stats extensions.

#### High level architecture
