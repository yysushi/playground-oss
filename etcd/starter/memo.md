# Getting Started with ETCD

## Run daemon in single node

<https://github.com/etcd-io/etcd/blob/master/Documentation/op-guide/container.md#running-a-single-node-etcd-1>

Set node ip and configure a Docker volume to store etcd data.

Run docker single cluster. The version is v3.3.25.

```
export NODE1=127.0.0.1
docker volume create --name etcd-data
export DATA_DIR="etcd-data"
REGISTRY=quay.io/coreos/etcd
# available from v3.2.5
REGISTRY=gcr.io/etcd-development/etcd

docker run \
  -p 2379:2379 \
  -p 2380:2380 \
  --volume=${DATA_DIR}:/etcd-data \
  --name etcd ${REGISTRY}:v3.3.25 \
  /usr/local/bin/etcd \
  --data-dir=/etcd-data --name node1 \
  --initial-advertise-peer-urls http://${NODE1}:2380 --listen-peer-urls http://0.0.0.0:2380 \
  --advertise-client-urls http://${NODE1}:2379 --listen-client-urls http://0.0.0.0:2379 \
  --initial-cluster node1=http://${NODE1}:2380
```

## Walk through APIs

<https://github.com/etcd-io/etcd/blob/master/Documentation/rfc/v3api.md#design>

- Flatten binary key-value space
- Keep the event history until compaction
- Support range query
- Replace TTL key with Lease
- Replace CAS/CAD with multi-object Txn
- Support efficient watching with multiple ranges
- RPC API supports the completed set of APIs.
- HTTP API supports a subset of APIs.

## ETCD client

<https://github.com/etcd-io/etcd/blob/master/Documentation/dev-guide/interacting_v3.md>

<https://github.com/etcd-io/etcd/blob/master/etcdctl/README.md>

```
koketani: ~/g/g/k/p/e/starter (etcd x?)$ docker exec -e ETCDCTL_API=3 -i -t etcd etcdctl --help
NAME:
        etcdctl - A simple command line client for etcd3.

USAGE:
        etcdctl

VERSION:
        3.3.25

API VERSION:
        3.3


COMMANDS:
        get                     Gets the key or a range of keys
        put                     Puts the given key into the store
        del                     Removes the specified key or range of keys [key, range_end)
        txn                     Txn processes all the requests in one transaction
        compaction              Compacts the event history in etcd
        alarm disarm            Disarms all alarms
        alarm list              Lists all alarms
        defrag                  Defragments the storage of the etcd members with given endpoints
        endpoint health         Checks the healthiness of endpoints specified in `--endpoints` flag
        endpoint status         Prints out the status of endpoints specified in `--endpoints` flag
        endpoint hashkv         Prints the KV history hash for each endpoint in --endpoints
        move-leader             Transfers leadership to another etcd cluster member.
        watch                   Watches events stream on keys or prefixes
        version                 Prints the version of etcdctl
        lease grant             Creates leases
        lease revoke            Revokes leases
        lease timetolive        Get lease information
        lease list              List all active leases
        lease keep-alive        Keeps leases alive (renew)
        member add              Adds a member into the cluster
        member remove           Removes a member from the cluster
        member update           Updates a member in the cluster
        member list             Lists all members in the cluster
        snapshot save           Stores an etcd node backend snapshot to a given file
        snapshot restore        Restores an etcd member snapshot to an etcd directory
        snapshot status         Gets backend snapshot status of a given file
        make-mirror             Makes a mirror at the destination etcd cluster
        migrate                 Migrates keys in a v2 store to a mvcc store
        lock                    Acquires a named lock
        elect                   Observes and participates in leader election
        auth enable             Enables authentication
        auth disable            Disables authentication
        user add                Adds a new user
        user delete             Deletes a user
        user get                Gets detailed information of a user
        user list               Lists all users
        user passwd             Changes password of user
        user grant-role         Grants a role to a user
        user revoke-role        Revokes a role from a user
        role add                Adds a new role
        role delete             Deletes a role
        role get                Gets detailed information of a role
        role list               Lists all roles
        role grant-permission   Grants a key to a role
        role revoke-permission  Revokes a key from a role
        check perf              Check the performance of the etcd cluster
        help                    Help about any command

OPTIONS:
      --cacert=""                               verify certificates of TLS-enabled secure servers using this CA bundle
      --cert=""                                 identify secure client using this TLS certificate file
      --command-timeout=5s                      timeout for short running command (excluding dial timeout)
      --debug[=false]                           enable client-side debug logging
      --dial-timeout=2s                         dial timeout for client connections
  -d, --discovery-srv=""                        domain name to query for SRV records describing cluster endpoints
      --endpoints=[127.0.0.1:2379]              gRPC endpoints
      --hex[=false]                             print byte strings as hex encoded strings
      --insecure-discovery[=true]               accept insecure SRV records describing cluster endpoints
      --insecure-skip-tls-verify[=false]        skip server certificate verification (CAUTION: this option should be enabled only for testing purposes)
      --insecure-transport[=true]               disable transport security for client connections
      --keepalive-time=2s                       keepalive time for client connections
      --keepalive-timeout=6s                    keepalive timeout for client connections
      --key=""                                  identify secure client using this TLS key file
      --user=""                                 username[:password] for authentication (prompt if password is not supplied)
  -w, --write-out="simple"                      set the output format (fields, json, protobuf, simple, table)
```

- version

```
koketani: ~/g/g/k/p/e/starter (etcd x?)$ docker exec -i -t etcd etcdctl -v
etcdctl version: 3.3.25
API version: 2
koketani: ~/g/g/k/p/e/starter (etcd x?)$ docker exec -e ETCDCTL_API=3 -i -t etcd etcdctl version
etcdctl version: 3.3.25
API version: 3.3
```

- write and read a key

Read with a revision allows us to check old data.

```
koketani: ~/g/g/k/p/e/starter (etcd x?)$ docker exec -e ETCDCTL_API=3 -i -t etcd etcdctl put foo bar
OK
koketani: ~/g/g/k/p/e/starter (etcd x?)$ docker exec -e ETCDCTL_API=3 -i -t etcd etcdctl get --prefix ""
foo
bar
```

- index in quoram

```
koketani: ~/g/g/k/p/e/starter (etcd x?)$ docker exec -e ETCDCTL_API=3 -i -t etcd etcdctl get foo -w json
{"header":{"cluster_id":12743156124158006367,"member_id":16536637930688536249,"revision":6,"raft_term":5},"kvs":[{"key":"Zm9v","create_revision":6,"mod_revision":6,"version":1,"value":"YmFy"}],"count":1}
```

- read past data

```
koketani: ~/g/g/k/p/e/starter (etcd x?)$ docker exec -e ETCDCTL_API=3 -i -t etcd etcdctl put foo bar2
koketani: ~/g/g/k/p/e/starter (etcd x?)$ docker exec -e ETCDCTL_API=3 -i -t etcd etcdctl get foo --rev 6
foo
bar
```

- compact

```
koketani: ~/g/g/k/p/e/starter (etcd x?)$ docker exec -e ETCDCTL_API=3 -i -t etcd etcdctl compact 7
compacted revision 7
koketani: ~/g/g/k/p/e/starter (etcd x?)$ docker exec -e ETCDCTL_API=3 -i -t etcd etcdctl get foo --rev 6
{"level":"warn","ts":"2020-11-01T05:09:28.845Z","caller":"clientv3/retry_interceptor.go:62","msg":"retrying of unary invoker failed","target":"endpoint://client-5a05e377-6391-42a3-bb45-0ddcff8c3d2e/127.0.0.1:2379","attempt":0,"error":"rpc error: code = OutOfRange desc = etcdserver: mvcc: required revision has been compacted"}
Error: etcdserver: mvcc: required revision has been compacted
```

- lease

leased key will be deleted with timeout

keep-alive allows us to clear ttl

```
koketani: ~/g/g/k/p/e/starter (etcd x?)$ docker exec -e ETCDCTL_API=3 -i -t etcd etcdctl lease grant 300
lease 5eb97581b4711e40 granted with TTL(300s)
koketani: ~/g/g/k/p/e/starter (etcd x?)$ docker exec -e ETCDCTL_API=3 -i -t etcd etcdctl put --lease=5eb97581b4711e40 foo3 bar3
OK
koketani: ~/g/g/k/p/e/starter (etcd x?)$ docker exec -e ETCDCTL_API=3 -i -t etcd etcdctl lease timetolive --keys 5eb97581b4711e40
lease 5eb97581b4711e40 granted with TTL(300s), remaining(281s), attached keys([foo3])
koketani: ~/g/g/k/p/e/starter (etcd x?)$

docker exec -e ETCDCTL_API=3 -i -t etcd etcdctl get foo -w json
koketani: ~/g/g/k/p/e/starter (etcd x?)$ docker exec -e ETCDCTL_API=3 -i -t etcd etcdctl get foo -w json
koketani: ~/g/g/k/p/e/starter (etcd x?)$ docker exec -e ETCDCTL_API=3 -i -t etcd etcdctl get foo -w json
koketani: ~/g/g/k/p/e/starter (etcd x?)$ docker exec -e ETCDCTL_API=3 -i -t etcd etcdctl get foo3 -w json
{"header":{"cluster_id":12743156124158006367,"member_id":16536637930688536249,"revision":10,"raft_term":5},"kvs":[{"key":"Zm9vMw==","create_revision":10,"mod_revision":10,"version":1,"value":"YmFyMw==","lease":6825615910195240512}],"count":1}
koketani: ~/g/g/k/p/e/starter (etcd x?)$ docker exec -e ETCDCTL_API=3 -i -t etcd etcdctl lease keep-alive 5eb97581b4711e40
lease 5eb97581b4711e40 keepalived with TTL(300)
^C%                                                                                                                                                                                              koketani: ~/g/g/k/p/e/starter (etcd x?)$ docker exec -e ETCDCTL_API=3 -i -t etcd etcdctl lease timetolive --keys 5eb97581b4711e40
lease 5eb97581b4711e40 granted with TTL(300s), remaining(293s), attached keys([foo3])
koketani: ~/g/g/k/p/e/starter (etcd x?)$ docker exec -e ETCDCTL_API=3 -i -t etcd etcdctl lease revoke 5eb97581b4711e40
lease 5eb97581b4711e40 revoked
koketani: ~/g/g/k/p/e/starter (etcd x?)$ docker exec -e ETCDCTL_API=3 -i -t etcd etcdctl lease timetolive --keys 5eb97581b4711e40
lease 5eb97581b4711e40 already expired
koketani: ~/g/g/k/p/e/starter (etcd x?)$ docker exec -e ETCDCTL_API=3 -i -t etcd etcdctl get foo3 -w json
{"header":{"cluster_id":12743156124158006367,"member_id":16536637930688536249,"revision":11,"raft_term":5}}
```

- watch

```
koketani: ~/g/g/k/p/e/starter (etcd x?)$ date; docker exec -e ETCDCTL_API=3 -i -t etcd etcdctl watch foo1 foo9
Sun Nov  1 14:36:10 JST 2020
PUT
foo4
bar4
^C% 
```

```
koketani: ~/g/g/k/p/e/starter (etcd x?)$ docker exec -e ETCDCTL_API=3 -i -t etcd etcdctl put foo4 bar4
OK
```

## Terminology

<https://github.com/etcd-io/etcd/blob/master/Documentation/learning/glossary.md>

## Check data to be persisted

```
koketani: ~/g/g/k/p/e/starter (etcd x?)$ docker volume inspect etcd-data
[
    {
        "CreatedAt": "2020-10-31T15:35:03Z",
        "Driver": "local",
        "Labels": {},
        "Mountpoint": "/var/lib/docker/volumes/etcd-data/_data",
        "Name": "etcd-data",
        "Options": {},
        "Scope": "local"
    }
]
koketani: ~/g/g/k/p/e/starter (etcd x?)$ docker run -v etcd-data:/etcd-data -i -t ubuntu:20.04 bash
root@aaeb5303e1be:/# ls -R /etcd-data
/etcd-data:
member

/etcd-data/member:
snap  wal

/etcd-data/member/snap:
db

/etcd-data/member/wal:
0.tmp  0000000000000000-0000000000000000.wal
```

## TODO

liner

faq https://github.com/etcd-io/etcd/blob/master/Documentation/faq.md

consistent https://github.com/etcd-io/etcd/blob/master/Documentation/learning/api_guarantees.md#isolation-level-and-consistency-of-replicas

https://github.com/etcd-io/etcd/blob/master/Documentation/dev-guide/api_reference_v3.md#message-rangerequest-apietcdserverpbrpcproto

https://coreos.com/blog/etcd-3-1-announcement.html

シリアライズの意味 https://ja.wikipedia.org/wiki/%E3%82%B7%E3%83%AA%E3%82%A2%E3%83%A9%E3%82%A4%E3%82%BA

lock https://github.com/etcd-io/etcd/tree/master/Documentation/learning/lock
