# Note for Docker Registry

## Commands

- run docker registry and daemon

```shell-session
$ docker-compose up
```

- [commit, tag and push](https://docs.docker.com/engine/reference/commandline/push/#examples)
  - target tag should include registry url.

```shell-session
$ docker commit [container-id] [image-name]
$ docker tag [source-tag] [target-tag]
$ docker push [target-tag]
$ docker images
```

```shell-session
$ docker commit c16378f943fe rhel-httpd
$ docker tag rhel-httpd registry-host:5000/myadmin/rhel-httpd
$ docker push registry-host:5000/myadmin/rhel-httpd
$ docker images
```


## memo

- what we can configure for registry
  - [allow-nondistributable-artifacts, registry-mirror, insecure-registry](https://github.com/moby/moby/blob/master/cmd/dockerd/config.go#L107-L109)
  ```golang
  flags.Var(ana, "allow-nondistributable-artifacts", "Allow push of nondistributable artifacts to registry")
  flags.Var(mirrors, "registry-mirror", "Preferred Docker registry mirror")
  flags.Var(insecureRegistries, "insecure-registry", "Enable insecure registry communication")
  ```
