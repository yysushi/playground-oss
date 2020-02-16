# Note for docker multi-stage

- log

```shell-session
koketani:multi-stage (master>)$ docker build .
Sending build context to Docker daemon  3.072kB
Step 1/6 : FROM golang:alpine AS builder
 ---> e1fd9820be16
Step 2/6 : WORKDIR $GOPATH/src/somewhere/multi-stage
 ---> Using cache
 ---> ef2aedf85f4f
Step 3/6 : COPY . .
 ---> Using cache
 ---> 58df15a50c30
Step 4/6 : RUN go install
 ---> Using cache
 ---> 3e966fc52451
Step 5/6 : FROM alpine:latest
 ---> e7d92cdc71fe
Step 6/6 : COPY --from=builder /go/bin/multi-stage /usr/local/bin/
 ---> Using cache
 ---> a606fbcc7651
Successfully built a606fbcc7651
koketani:multi-stage (master>)$ docker image inspect a606fbcc7651
[
    {
        "Id": "sha256:a606fbcc76517c77a3283d47ffc04b5dbb2623adbb64c486fe3c562ecc3190ae",
        "RepoTags": [
            "hoge:latest"
        ],
        "RepoDigests": [],
        "Parent": "sha256:e7d92cdc71feacf90708cb59182d0df1b911f8ae022d29e8e95d75ca6a99776a",
        "Comment": "",
        "Created": "2020-02-16T10:29:41.8379409Z",
        "Container": "",
        "ContainerConfig": {
            "Hostname": "",
            "Domainname": "",
            "User": "",
            "AttachStdin": false,
            "AttachStdout": false,
            "AttachStderr": false,
            "Tty": false,
            "OpenStdin": false,
            "StdinOnce": false,
            "Env": [
                "PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin"
            ],
            "Cmd": [
                "/bin/sh",
                "-c",
                "#(nop) COPY file:f9e0376c5638a44ef80f635b87ece945834977380121b97165d7641fa7d5b378 in /usr/local/bin/ "
            ],
            "ArgsEscaped": true,
            "Image": "sha256:e7d92cdc71feacf90708cb59182d0df1b911f8ae022d29e8e95d75ca6a99776a",
            "Volumes": null,
            "WorkingDir": "",
            "Entrypoint": null,
            "OnBuild": null,
            "Labels": null
        },
        "DockerVersion": "19.03.5",
        "Author": "",
        "Config": {
            "Hostname": "",
            "Domainname": "",
            "User": "",
            "AttachStdin": false,
            "AttachStdout": false,
            "AttachStderr": false,
            "Tty": false,
            "OpenStdin": false,
            "StdinOnce": false,
            "Env": [
                "PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin"
            ],
            "Cmd": [
                "/bin/sh"
            ],
            "ArgsEscaped": true,
            "Image": "sha256:e7d92cdc71feacf90708cb59182d0df1b911f8ae022d29e8e95d75ca6a99776a",
            "Volumes": null,
            "WorkingDir": "",
            "Entrypoint": null,
            "OnBuild": null,
            "Labels": null
        },
        "Architecture": "amd64",
        "Os": "linux",
        "Size": 7600101,
        "VirtualSize": 7600101,
        "GraphDriver": {
            "Data": {
                "LowerDir": "/var/lib/docker/overlay2/0fa7e0ff0b3508e525c137073ed8bfc4644626c7accad3ef9c0a7e6f61d5c48b/diff",
                "MergedDir": "/var/lib/docker/overlay2/4beadcc67a761cb3635a5c4f48d53e7abc3552f787b0a83baccb8f049cabe457/merged",
                "UpperDir": "/var/lib/docker/overlay2/4beadcc67a761cb3635a5c4f48d53e7abc3552f787b0a83baccb8f049cabe457/diff",
                "WorkDir": "/var/lib/docker/overlay2/4beadcc67a761cb3635a5c4f48d53e7abc3552f787b0a83baccb8f049cabe457/work"
            },
            "Name": "overlay2"
        },
        "RootFS": {
            "Type": "layers",
            "Layers": [
                "sha256:5216338b40a7b96416b8b9858974bbe4acc3096ee60acbc4dfb1ee02aecceb10",
                "sha256:c6bddf7d8cf6d75ab5a78bc65c012ae0e4cb78fff3b8027887f362c67d7e1fef"
            ]
        },
        "Metadata": {
            "LastTagTime": "2020-02-16T10:29:41.8838082Z"
        }
    }
]
```
