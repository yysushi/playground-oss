# Logstash

- run

```nolang
\$ docker run -v $(pwd):/tmp/logstash -e ELASTICSEARCH_START=0 -e KIBANA_START=0 -it sebp/elk /bin/bash
```

<https://www.elastic.co/blog/a-practical-introduction-to-logstash>
