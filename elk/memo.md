# Notes for elk

## References

<https://elk-docker.readthedocs.io/>

## Elastic Stack

- elk
  - logstash; import/export/parse data
  - elastic search; store/search data
  - kibana; visualize data

- x-pack; security, alert, monitoring, reporting and graph
- beat;

## Elastic Search

### technical terms

<https://noti.st/johtani/6bX2ZY/elastic-stack-for#s9B7fDm>

- index; logical data group like database in RDB
- replication; to avoid SPOF and improve read performance
- sharding; split data to multiple hosts, improve write performance and control write flow

- full text search; search given strings over from multiple documents

- index; place to search data, which is used by search engine
- document; data stored in search engine
- field; attribute included in docuemnt
- query; search condition and serach formula
- schema; definition of document structure
- termn and token
  - keyword which added by index
  - words in a scentence
  - including word offset in the scentence

- one log is one document
- one log is indexed
- one log consists of some words
- some words in one log are indexed

- by specifying specific word, we can find some logs

### example

- curl `http://localhost:9200`

```json
{
  "name" : "SxqhSvG",
  "cluster_name" : "elasticsearch",
  "cluster_uuid" : "dV4h-VoHRlefw_iZsjBDOg",
  "version" : {
    "number" : "6.5.1",
    "build_flavor" : "default",
    "build_type" : "tar",
    "build_hash" : "8c58350",
    "build_date" : "2018-11-16T02:22:42.182257Z",
    "build_snapshot" : false,
    "lucene_version" : "7.5.0",
    "minimum_wire_compatibility_version" : "5.6.0",
    "minimum_index_compatibility_version" : "5.0.0"
  },
  "tagline" : "You Know, for Search"
}
```

- curl `http://localhost:9200/\_search?pretty`

```json
{
  "took" : 1,
  "timed_out" : false,
  "_shards" : {
    "total" : 6,
    "successful" : 6,
    "skipped" : 0,
    "failed" : 0
  },
  "hits" : {
    "total" : 3,
    "max_score" : 1.0,
    "hits" : [
      {
        "_index" : ".kibana_1",
        "_type" : "doc",
        "_id" : "space:default",
        "_score" : 1.0,
        "_source" : {
          "space" : {
            "name" : "Default",
            "description" : "This is your default space!",
            "color" : "#00bfb3",
            "_reserved" : true
          },
          "type" : "space",
          "updated_at" : "2019-01-06T08:35:38.205Z"
        }
      },
      {
        "_index" : ".kibana_1",
        "_type" : "doc",
        "_id" : "config:6.5.1",
        "_score" : 1.0,
        "_source" : {
          "config" : {
            "buildNum" : 18763
          },
          "type" : "config",
          "updated_at" : "2019-01-06T08:45:57.962Z"
        }
      },
      {
        "_index" : "logstash-2019.01.06",
        "_type" : "doc",
        "_id" : "gFxoImgBFJ9iT7SyVuJg",
        "_score" : 1.0,
        "_source" : {
          "@timestamp" : "2019-01-06T09:04:49.225Z",
          "host" : "148417cdb3ec",
          "message" : "this is a dummy entry",
          "@version" : "1"
        }
      }
    ]
  }
}
```

- CRUD

```json
```

## Logstash

- input
- filter
- output
