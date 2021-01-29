# Quick start

## Run envoy

- pull

```shell-session
koketani: ~/g/g/k/p/envoy (envoy !?)$ docker pull envoyproxy/envoy-dev:latest
latest: Pulling from envoyproxy/envoy-dev
171857c49d0f: Pull complete
419640447d26: Pull complete
61e52f862619: Pull complete
bd8bd40cd5f4: Pull complete
fdc0ee30b80f: Pull complete
bc5decbb4064: Pull complete
3b6c0d66e8ba: Pull complete
7a7d62d81809: Pull complete
bc63ad63dd92: Pull complete
5e8d4e3fc45a: Pull complete
Digest: sha256:15190c8f3829213813dbe360e07540f0b52c0f71e4388df8f5d12c9606eb9094
Status: Downloaded newer image for envoyproxy/envoy-dev:latest
docker.io/envoyproxy/envoy-dev:latest
```

- version and help

```
koketani: ~/g/g/k/p/envoy (envoy !?)$ docker run --rm -i -t envoyproxy/envoy-dev bash
envoy@0494903365f5:/$ envoy --version

envoy  version: bd73f3c4da0efffb2593d7c9ecf87788856dc052/1.17.0-dev/Clean/RELEASE/BoringSSL

envoy@0494903365f5:/$ envoy --help

USAGE:

   envoy  [--socket-mode <string>] [--socket-path <string>]
          [--disable-extensions <string>] [--use-fake-symbol-table <bool>]
          [--cpuset-threads] [--enable-mutex-tracing]
          [--disable-hot-restart] [--mode <string>]
          [--parent-shutdown-time-s <uint32_t>] [--drain-strategy <string>]
          [--drain-time-s <uint32_t>] [--file-flush-interval-msec
          <uint32_t>] [--service-zone <string>] [--service-node <string>]
          [--service-cluster <string>] [--hot-restart-version]
          [--restart-epoch <uint32_t>] [--log-path <string>]
          [--log-format-prefix-with-location <bool>]
          [--enable-fine-grain-logging] [--log-format-escaped]
          [--log-format <string>] [--component-log-level <string>] [-l
          <string>] [--local-address-ip-version <string>]
          [--admin-address-path <string>] [--ignore-unknown-dynamic-fields]
          [--reject-unknown-dynamic-fields] [--allow-unknown-static-fields]
          [--allow-unknown-fields] [--bootstrap-version <string>]
          [--config-yaml <string>] [-c <string>] [--concurrency <uint32_t>]
          [--base-id-path <string>] [--use-dynamic-base-id] [--base-id
          <uint32_t>] [--] [--version] [-h]


Where:

   --socket-mode <string>
     Socket file permission

   --socket-path <string>
     Path to hot restart socket file

   --disable-extensions <string>
     Comma-separated list of extensions to disable

   --use-fake-symbol-table <bool>
     Use fake symbol table implementation

   --cpuset-threads
     Get the default # of worker threads from cpuset size

   --enable-mutex-tracing
     Enable mutex contention tracing functionality

   --disable-hot-restart
     Disable hot restart functionality

   --mode <string>
     One of 'serve' (default; validate configs and then serve traffic
     normally) or 'validate' (validate configs and exit).

   --parent-shutdown-time-s <uint32_t>
     Hot restart parent shutdown time in seconds

   --drain-strategy <string>
     Hot restart drain sequence behaviour, one of 'gradual' (default) or
     'immediate'.

   --drain-time-s <uint32_t>
     Hot restart and LDS removal drain time in seconds

   --file-flush-interval-msec <uint32_t>
     Interval for log flushing in msec

   --service-zone <string>
     Zone name

   --service-node <string>
     Node name

   --service-cluster <string>
     Cluster name

   --hot-restart-version
     hot restart compatibility version

   --restart-epoch <uint32_t>
     hot restart epoch #

   --log-path <string>
     Path to logfile

   --log-format-prefix-with-location <bool>
     Prefix all occurrences of '%v' in log format with with '[%g:%#] '
     ('[path/to/file.cc:99] ').

   --enable-fine-grain-logging
     Logger mode: enable file level log control(Fancy Logger)or not

   --log-format-escaped
     Escape c-style escape sequences in the application logs

   --log-format <string>
     Log message format in spdlog syntax (see
     https://github.com/gabime/spdlog/wiki/3.-Custom-formatting)

     Default is "[%Y-%m-%d %T.%e][%t][%l][%n] [%g:%#] %v"

   --component-log-level <string>
     Comma separated list of component log levels. For example
     upstream:debug,config:trace

   -l <string>,  --log-level <string>
     Log levels: [trace][debug][info][warning
     |warn][error][critical][off]

     Default is [info]

   --local-address-ip-version <string>
     The local IP address version (v4 or v6).

   --admin-address-path <string>
     Admin address path

   --ignore-unknown-dynamic-fields
     ignore unknown fields in dynamic configuration

   --reject-unknown-dynamic-fields
     reject unknown fields in dynamic configuration

   --allow-unknown-static-fields
     allow unknown fields in static configuration

   --allow-unknown-fields
     allow unknown fields in static configuration (DEPRECATED)

   --bootstrap-version <string>
     API version to parse the bootstrap config as (e.g. 3). If unset, all
     known versions will be attempted

   --config-yaml <string>
     Inline YAML configuration, merges with the contents of --config-path

   -c <string>,  --config-path <string>
     Path to configuration file

   --concurrency <uint32_t>
     # of worker threads to run

   --base-id-path <string>
     path to which the base ID is written

   --use-dynamic-base-id
     the server chooses a base ID dynamically. Supersedes a static base ID.
     May not be used when the restart epoch is non-zero.

   --base-id <uint32_t>
     base ID so that multiple envoys can run on the same host if needed

   --,  --ignore_rest
     Ignores the rest of the labeled arguments following this flag.

   --version
     Displays version information and exits.

   -h,  --help
     Displays usage information and exits.


   envoy


koketani: ~/g/g/k/p/envoy (envoy !?)$ docker run --rm envoyproxy/envoy-dev --version

envoy  version: bd73f3c4da0efffb2593d7c9ecf87788856dc052/1.17.0-dev/Clean/RELEASE/BoringSSL
```

- run with demo configuration

```
koketani: ~/g/g/k/p/envoy (envoy !?)$ docker run --rm --name envoy -p 9901:9901 -p 10000:10000 -v $(pwd)/reverse:/reverse -i -t envoyproxy/envoy-dev -c /reverse/envoy-demo.yaml
```
