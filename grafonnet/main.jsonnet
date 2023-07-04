local g = import 'g.libsonnet';

g.dashboard.new('A title of dashboard')
+ g.dashboard.withPanels([
  g.panel.timeSeries.new('A title of panel')
  + g.panel.timeSeries.queryOptions.withTargets([
    g.query.prometheus.new(
      'Prometheus',
      'go_memstats_heap_alloc_bytes',
    ),
  ]),
  g.panel.timeSeries.new('A title of panel')
  + g.panel.timeSeries.queryOptions.withTargets([
    g.query.prometheus.new(
      'Prometheus',
      'go_gc_duration_seconds_count',
    ),
  ]),
  g.panel.timeSeries.new('A title of panel')
  + g.panel.timeSeries.queryOptions.withTargets([
    g.query.prometheus.new(
      'Prometheus',
      'prometheus_http_request_duration_seconds_bucket{handler="/graph"}',
    ),
  ]),
  g.panel.timeSeries.new('A title of panel')
  + g.panel.timeSeries.queryOptions.withTargets([
    g.query.prometheus.new(
      'Prometheus',
      'histogram_quantile(0.9, rate(prometheus_http_request_duration_seconds_bucket{handler="/graph"}[5m]))',
    ),
  ]),
])
