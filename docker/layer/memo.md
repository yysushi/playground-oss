# Note for docker layer

We should consider docker image layer and pack multiple commands to one with using "&&" if needed.

- image sizes

```shell-session
koketani:layer (master %=)$ docker images | grep -e big -e layer
layer                                   latest              78ac3a0fab3c        4 seconds ago        64.2MB
big2                                    latest              0bda207fc332        About a minute ago   92.1MB
big                                     latest              93e935353c00        3 minutes ago        92.1MB
```

- history

```shell-session
koketani:layer (master %=)$ diff <(docker history big2) <(docker history big)
2d1
< 0bda207fc332        6 minutes ago       /bin/sh -c apt clean && rm -rf /var/lib/apt/…   0B
```

```shell-session
koketani:layer (master %=)$ diff <(docker history big2) <(docker history layer)
2,3c2
< 0bda207fc332        7 minutes ago       /bin/sh -c apt clean && rm -rf /var/lib/apt/…   0B
< 93e935353c00        9 minutes ago       /bin/sh -c apt update                           27.9MB
---
> 78ac3a0fab3c        6 minutes ago       /bin/sh -c apt update && apt clean && rm -rf…   0B
```

- log

```shell-session
koketani:layer (master %=)$ docker build -t layer .
Sending build context to Docker daemon  4.096kB
Step 1/2 : FROM ubuntu:18.04
 ---> ccc6e87d482b
Step 2/2 : RUN apt update && apt clean && rm -rf /var/lib/apt/lists/*
 ---> Running in 2fef21791843

WARNING: apt does not have a stable CLI interface. Use with caution in scripts.

Get:1 http://security.ubuntu.com/ubuntu bionic-security InRelease [88.7 kB]
Get:2 http://archive.ubuntu.com/ubuntu bionic InRelease [242 kB]
Get:3 http://security.ubuntu.com/ubuntu bionic-security/multiverse amd64 Packages [7064 B]
Get:4 http://archive.ubuntu.com/ubuntu bionic-updates InRelease [88.7 kB]
Get:5 http://security.ubuntu.com/ubuntu bionic-security/main amd64 Packages [817 kB]
Get:6 http://archive.ubuntu.com/ubuntu bionic-backports InRelease [74.6 kB]
Get:7 http://archive.ubuntu.com/ubuntu bionic/multiverse amd64 Packages [186 kB]
Get:8 http://security.ubuntu.com/ubuntu bionic-security/restricted amd64 Packages [27.5 kB]
Get:9 http://security.ubuntu.com/ubuntu bionic-security/universe amd64 Packages [818 kB]
Get:10 http://archive.ubuntu.com/ubuntu bionic/restricted amd64 Packages [13.5 kB]
Get:11 http://archive.ubuntu.com/ubuntu bionic/main amd64 Packages [1344 kB]
Get:12 http://archive.ubuntu.com/ubuntu bionic/universe amd64 Packages [11.3 MB]
Get:13 http://archive.ubuntu.com/ubuntu bionic-updates/multiverse amd64 Packages [11.1 kB]
Get:14 http://archive.ubuntu.com/ubuntu bionic-updates/universe amd64 Packages [1345 kB]
Get:15 http://archive.ubuntu.com/ubuntu bionic-updates/main amd64 Packages [1104 kB]
Get:16 http://archive.ubuntu.com/ubuntu bionic-updates/restricted amd64 Packages [41.2 kB]
Get:17 http://archive.ubuntu.com/ubuntu bionic-backports/universe amd64 Packages [4252 B]
Get:18 http://archive.ubuntu.com/ubuntu bionic-backports/main amd64 Packages [2496 B]
Fetched 17.6 MB in 24s (736 kB/s)
Reading package lists...
Building dependency tree...
Reading state information...
18 packages can be upgraded. Run 'apt list --upgradable' to see them.

WARNING: apt does not have a stable CLI interface. Use with caution in scripts.

Removing intermediate container 2fef21791843
 ---> 78ac3a0fab3c
Successfully built 78ac3a0fab3c
Successfully tagged layer:latest
```

```shell-session
koketani:layer (master %=)$ docker build -t big2 -f Dockerfile.big2 .
Sending build context to Docker daemon  3.072kB
Step 1/3 : FROM ubuntu:18.04
 ---> ccc6e87d482b
Step 2/3 : RUN apt update
 ---> Using cache
 ---> 93e935353c00
Step 3/3 : RUN apt clean && rm -rf /var/lib/apt/lists/*
 ---> Running in a1395a975869

WARNING: apt does not have a stable CLI interface. Use with caution in scripts.

Removing intermediate container a1395a975869
 ---> 0bda207fc332
Successfully built 0bda207fc332
Successfully tagged big2:latest
```
