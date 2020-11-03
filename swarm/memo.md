# Playground for Docker Swarm

## Build and run container

```
koketani: ~/g/g/k/p/swarm (swarm ?)$ docker build -t swarm-test .
```

1. publish port

```
koketani: ~/g/g/k/p/swarm (swarm ?)$ docker run --rm --name swarm-test -p 8080:8080 swarm-test
ping from 172.17.0.1:35526
ping from 172.17.0.1:35530
```

```
koketani: ~/g/g/k/p/swarm (swarm ?)$ docker inspect swarm-test
[
    {
        "Id": "ebd36447be7ae0059a7395fa4528cb16131c58b8a42b397ba823be166cdf6228",
        "Created": "2020-11-02T17:30:58.663768Z",
        "Path": "/bin/swarm-test",
        "Args": [],
        "State": {
            "Status": "running",
            "Running": true,
            "Paused": false,
            "Restarting": false,
            "OOMKilled": false,
            "Dead": false,
            "Pid": 96765,
            "ExitCode": 0,
            "Error": "",
            "StartedAt": "2020-11-02T17:30:58.9545962Z",
            "FinishedAt": "0001-01-01T00:00:00Z"
        },
        "Image": "sha256:6ae7f577836e88abd79a56415237479d520b209d650fce2849907216b7352f14",
        "ResolvConfPath": "/var/lib/docker/containers/ebd36447be7ae0059a7395fa4528cb16131c58b8a42b397ba823be166cdf6228/resolv.conf",
        "HostnamePath": "/var/lib/docker/containers/ebd36447be7ae0059a7395fa4528cb16131c58b8a42b397ba823be166cdf6228/hostname",
        "HostsPath": "/var/lib/docker/containers/ebd36447be7ae0059a7395fa4528cb16131c58b8a42b397ba823be166cdf6228/hosts",
        "LogPath": "/var/lib/docker/containers/ebd36447be7ae0059a7395fa4528cb16131c58b8a42b397ba823be166cdf6228/ebd36447be7ae0059a7395fa4528cb16131c58b8a42b397ba823be166cdf6228-json.log",
        "Name": "/swarm-test",
        "RestartCount": 0,
        "Driver": "overlay2",
        "Platform": "linux",
        "MountLabel": "",
        "ProcessLabel": "",
        "AppArmorProfile": "",
        "ExecIDs": null,
        "HostConfig": {
            "Binds": null,
            "ContainerIDFile": "",
            "LogConfig": {
                "Type": "json-file",
                "Config": {}
            },
            "NetworkMode": "default",
            "PortBindings": {
                "8080/tcp": [
                    {
                        "HostIp": "",
                        "HostPort": "8080"
                    }
                ]
            },
            "RestartPolicy": {
                "Name": "no",
                "MaximumRetryCount": 0
            },
            "AutoRemove": false,
            "VolumeDriver": "",
            "VolumesFrom": null,
            "CapAdd": null,
            "CapDrop": null,
            "Capabilities": null,
            "Dns": [],
            "DnsOptions": [],
            "DnsSearch": [],
            "ExtraHosts": null,
            "GroupAdd": null,
            "IpcMode": "private",
            "Cgroup": "",
            "Links": null,
            "OomScoreAdj": 0,
            "PidMode": "",
            "Privileged": false,
            "PublishAllPorts": false,
            "ReadonlyRootfs": false,
            "SecurityOpt": null,
            "UTSMode": "",
            "UsernsMode": "",
            "ShmSize": 67108864,
            "Runtime": "runc",
            "ConsoleSize": [
                0,
                0
            ],
            "Isolation": "",
            "CpuShares": 0,
            "Memory": 0,
            "NanoCpus": 0,
            "CgroupParent": "",
            "BlkioWeight": 0,
            "BlkioWeightDevice": [],
            "BlkioDeviceReadBps": null,
            "BlkioDeviceWriteBps": null,
            "BlkioDeviceReadIOps": null,
            "BlkioDeviceWriteIOps": null,
            "CpuPeriod": 0,
            "CpuQuota": 0,
            "CpuRealtimePeriod": 0,
            "CpuRealtimeRuntime": 0,
            "CpusetCpus": "",
            "CpusetMems": "",
            "Devices": [],
            "DeviceCgroupRules": null,
            "DeviceRequests": null,
            "KernelMemory": 0,
            "KernelMemoryTCP": 0,
            "MemoryReservation": 0,
            "MemorySwap": 0,
            "MemorySwappiness": null,
            "OomKillDisable": false,
            "PidsLimit": null,
            "Ulimits": null,
            "CpuCount": 0,
            "CpuPercent": 0,
            "IOMaximumIOps": 0,
            "IOMaximumBandwidth": 0,
            "MaskedPaths": [
                "/proc/asound",
                "/proc/acpi",
                "/proc/kcore",
                "/proc/keys",
                "/proc/latency_stats",
                "/proc/timer_list",
                "/proc/timer_stats",
                "/proc/sched_debug",
                "/proc/scsi",
                "/sys/firmware"
            ],
            "ReadonlyPaths": [
                "/proc/bus",
                "/proc/fs",
                "/proc/irq",
                "/proc/sys",
                "/proc/sysrq-trigger"
            ]
        },
        "GraphDriver": {
            "Data": {
                "LowerDir": "/var/lib/docker/overlay2/d6579f8a55cc061035811ab49df8e6356df075e9bfa81e9db0a026289fc64704-init/diff:/var/lib/docker/overlay2/2e5dba12f8c4bfafca82c6c5ca38aea1626d2e66a3a7067840becddc88f21e72/diff:/var/lib/docker/overlay2/78fb31d6eaf68e0adf8d5b1f059d143ee113a8d6152b367864c5c9d464f2e0d9/diff",
                "MergedDir": "/var/lib/docker/overlay2/d6579f8a55cc061035811ab49df8e6356df075e9bfa81e9db0a026289fc64704/merged",
                "UpperDir": "/var/lib/docker/overlay2/d6579f8a55cc061035811ab49df8e6356df075e9bfa81e9db0a026289fc64704/diff",
                "WorkDir": "/var/lib/docker/overlay2/d6579f8a55cc061035811ab49df8e6356df075e9bfa81e9db0a026289fc64704/work"
            },
            "Name": "overlay2"
        },
        "Mounts": [],
        "Config": {
            "Hostname": "ebd36447be7a",
            "Domainname": "",
            "User": "",
            "AttachStdin": false,
            "AttachStdout": true,
            "AttachStderr": true,
            "ExposedPorts": {
                "8080/tcp": {}
            },
            "Tty": false,
            "OpenStdin": false,
            "StdinOnce": false,
            "Env": [
                "PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin"
            ],
            "Cmd": null,
            "Image": "swarm-test",
            "Volumes": null,
            "WorkingDir": "",
            "Entrypoint": [
                "/bin/swarm-test"
            ],
            "OnBuild": null,
            "Labels": {}
        },
        "NetworkSettings": {
            "Bridge": "",
            "SandboxID": "cd524371ea3c5a7e8148721b35b837bcb9f387658b5826932a23d648d8fce6a5",
            "HairpinMode": false,
            "LinkLocalIPv6Address": "",
            "LinkLocalIPv6PrefixLen": 0,
            "Ports": {
                "8080/tcp": [
                    {
                        "HostIp": "0.0.0.0",
                        "HostPort": "8080"
                    }
                ]
            },
            "SandboxKey": "/var/run/docker/netns/cd524371ea3c",
            "SecondaryIPAddresses": null,
            "SecondaryIPv6Addresses": null,
            "EndpointID": "fdd665dd8d8484eb824a88af61878979113d6e97dd5033569fa5c7419c58db66",
            "Gateway": "172.17.0.1",
            "GlobalIPv6Address": "",
            "GlobalIPv6PrefixLen": 0,
            "IPAddress": "172.17.0.2",
            "IPPrefixLen": 16,
            "IPv6Gateway": "",
            "MacAddress": "02:42:ac:11:00:02",
            "Networks": {
                "bridge": {
                    "IPAMConfig": null,
                    "Links": null,
                    "Aliases": null,
                    "NetworkID": "74469048a1c85db6709b827490807151eaaf87dfa218e7552ab0e4ae72b81499",
                    "EndpointID": "fdd665dd8d8484eb824a88af61878979113d6e97dd5033569fa5c7419c58db66",
                    "Gateway": "172.17.0.1",
                    "IPAddress": "172.17.0.2",
                    "IPPrefixLen": 16,
                    "IPv6Gateway": "",
                    "GlobalIPv6Address": "",
                    "GlobalIPv6PrefixLen": 0,
                    "MacAddress": "02:42:ac:11:00:02",
                    "DriverOpts": null
                }
            }
        }
    }
]
```

2. host

```
koketani: ~/g/g/k/p/swarm (swarm ?)$ docker run --name swarm-test --network host swarm-test
```

```
koketani: ~/g/g/k/p/swarm (swarm ?)$ docker inspect swarm-test
[
    {
        "Id": "9ac66f2864ce42067aa3dc6d567b2e57b00bcdaafc26d54ee8494399fde16fcb",
        "Created": "2020-11-02T17:33:49.5032477Z",
        "Path": "/bin/swarm-test",
        "Args": [],
        "State": {
            "Status": "running",
            "Running": true,
            "Paused": false,
            "Restarting": false,
            "OOMKilled": false,
            "Dead": false,
            "Pid": 96968,
            "ExitCode": 0,
            "Error": "",
            "StartedAt": "2020-11-02T17:33:49.7105152Z",
            "FinishedAt": "0001-01-01T00:00:00Z"
        },
        "Image": "sha256:6ae7f577836e88abd79a56415237479d520b209d650fce2849907216b7352f14",
        "ResolvConfPath": "/var/lib/docker/containers/9ac66f2864ce42067aa3dc6d567b2e57b00bcdaafc26d54ee8494399fde16fcb/resolv.conf",
        "HostnamePath": "/var/lib/docker/containers/9ac66f2864ce42067aa3dc6d567b2e57b00bcdaafc26d54ee8494399fde16fcb/hostname",
        "HostsPath": "/var/lib/docker/containers/9ac66f2864ce42067aa3dc6d567b2e57b00bcdaafc26d54ee8494399fde16fcb/hosts",
        "LogPath": "/var/lib/docker/containers/9ac66f2864ce42067aa3dc6d567b2e57b00bcdaafc26d54ee8494399fde16fcb/9ac66f2864ce42067aa3dc6d567b2e57b00bcdaafc26d54ee8494399fde16fcb-json.log",
        "Name": "/swarm-test",
        "RestartCount": 0,
        "Driver": "overlay2",
        "Platform": "linux",
        "MountLabel": "",
        "ProcessLabel": "",
        "AppArmorProfile": "",
        "ExecIDs": null,
        "HostConfig": {
            "Binds": null,
            "ContainerIDFile": "",
            "LogConfig": {
                "Type": "json-file",
                "Config": {}
            },
            "NetworkMode": "host",
            "PortBindings": {},
            "RestartPolicy": {
                "Name": "no",
                "MaximumRetryCount": 0
            },
            "AutoRemove": false,
            "VolumeDriver": "",
            "VolumesFrom": null,
            "CapAdd": null,
            "CapDrop": null,
            "Capabilities": null,
            "Dns": [],
            "DnsOptions": [],
            "DnsSearch": [],
            "ExtraHosts": null,
            "GroupAdd": null,
            "IpcMode": "private",
            "Cgroup": "",
            "Links": null,
            "OomScoreAdj": 0,
            "PidMode": "",
            "Privileged": false,
            "PublishAllPorts": false,
            "ReadonlyRootfs": false,
            "SecurityOpt": null,
            "UTSMode": "",
            "UsernsMode": "",
            "ShmSize": 67108864,
            "Runtime": "runc",
            "ConsoleSize": [
                0,
                0
            ],
            "Isolation": "",
            "CpuShares": 0,
            "Memory": 0,
            "NanoCpus": 0,
            "CgroupParent": "",
            "BlkioWeight": 0,
            "BlkioWeightDevice": [],
            "BlkioDeviceReadBps": null,
            "BlkioDeviceWriteBps": null,
            "BlkioDeviceReadIOps": null,
            "BlkioDeviceWriteIOps": null,
            "CpuPeriod": 0,
            "CpuQuota": 0,
            "CpuRealtimePeriod": 0,
            "CpuRealtimeRuntime": 0,
            "CpusetCpus": "",
            "CpusetMems": "",
            "Devices": [],
            "DeviceCgroupRules": null,
            "DeviceRequests": null,
            "KernelMemory": 0,
            "KernelMemoryTCP": 0,
            "MemoryReservation": 0,
            "MemorySwap": 0,
            "MemorySwappiness": null,
            "OomKillDisable": false,
            "PidsLimit": null,
            "Ulimits": null,
            "CpuCount": 0,
            "CpuPercent": 0,
            "IOMaximumIOps": 0,
            "IOMaximumBandwidth": 0,
            "MaskedPaths": [
                "/proc/asound",
                "/proc/acpi",
                "/proc/kcore",
                "/proc/keys",
                "/proc/latency_stats",
                "/proc/timer_list",
                "/proc/timer_stats",
                "/proc/sched_debug",
                "/proc/scsi",
                "/sys/firmware"
            ],
            "ReadonlyPaths": [
                "/proc/bus",
                "/proc/fs",
                "/proc/irq",
                "/proc/sys",
                "/proc/sysrq-trigger"
            ]
        },
        "GraphDriver": {
            "Data": {
                "LowerDir": "/var/lib/docker/overlay2/c560aeef521e91844bed051970e8ead3b070a78df3ee9de6a7e39d96b06f7ba9-init/diff:/var/lib/docker/overlay2/2e5dba12f8c4bfafca82c6c5ca38aea1626d2e66a3a7067840becddc88f21e72/diff:/var/lib/docker/overlay2/78fb31d6eaf68e0adf8d5b1f059d143ee113a8d6152b367864c5c9d464f2e0d9/diff",
                "MergedDir": "/var/lib/docker/overlay2/c560aeef521e91844bed051970e8ead3b070a78df3ee9de6a7e39d96b06f7ba9/merged",
                "UpperDir": "/var/lib/docker/overlay2/c560aeef521e91844bed051970e8ead3b070a78df3ee9de6a7e39d96b06f7ba9/diff",
                "WorkDir": "/var/lib/docker/overlay2/c560aeef521e91844bed051970e8ead3b070a78df3ee9de6a7e39d96b06f7ba9/work"
            },
            "Name": "overlay2"
        },
        "Mounts": [],
        "Config": {
            "Hostname": "docker-desktop",
            "Domainname": "",
            "User": "",
            "AttachStdin": false,
            "AttachStdout": true,
            "AttachStderr": true,
            "Tty": false,
            "OpenStdin": false,
            "StdinOnce": false,
            "Env": [
                "PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin"
            ],
            "Cmd": null,
            "Image": "swarm-test",
            "Volumes": null,
            "WorkingDir": "",
            "Entrypoint": [
                "/bin/swarm-test"
            ],
            "OnBuild": null,
            "Labels": {}
        },
        "NetworkSettings": {
            "Bridge": "",
            "SandboxID": "40917d1633792e319ac41d233e7c8c7a7ac0696eac945e3fcb3110f76c5fd600",
            "HairpinMode": false,
            "LinkLocalIPv6Address": "",
            "LinkLocalIPv6PrefixLen": 0,
            "Ports": {},
            "SandboxKey": "/var/run/docker/netns/default",
            "SecondaryIPAddresses": null,
            "SecondaryIPv6Addresses": null,
            "EndpointID": "",
            "Gateway": "",
            "GlobalIPv6Address": "",
            "GlobalIPv6PrefixLen": 0,
            "IPAddress": "",
            "IPPrefixLen": 0,
            "IPv6Gateway": "",
            "MacAddress": "",
            "Networks": {
                "host": {
                    "IPAMConfig": null,
                    "Links": null,
                    "Aliases": null,
                    "NetworkID": "3b444893da3a579ec9ac3d409b6e93fa4fe9ed6939384098a6bbf4ee50dfc21c",
                    "EndpointID": "1c4d51c443eb4d303491f9832dabee8486ee0b6bfd7d973f6be01bc54ca506f2",
                    "Gateway": "",
                    "IPAddress": "",
                    "IPPrefixLen": 0,
                    "IPv6Gateway": "",
                    "GlobalIPv6Address": "",
                    "GlobalIPv6PrefixLen": 0,
                    "MacAddress": "",
                    "DriverOpts": null
                }
            }
        }
    }
]
```

## Create service

- create service and publish port afterwards

```
koketani: ~/g/g/k/p/swarm (swarm ?)$ docker service create --name swarm-test swarm-test
image swarm-test:latest could not be accessed on a registry to record
its digest. Each node will access swarm-test:latest independently,
possibly leading to different nodes running different
versions of the image.

ck648ikmfeggre3cj0mgn1zfu
overall progress: 1 out of 1 tasks
1/1: running   [==================================================>]
verify: Service converged
koketani: ~/g/g/k/p/swarm (swarm ?)$ docker service update --publish-add 8080:8080 swarm-test
swarm-test
overall progress: 1 out of 1 tasks
1/1: running   [==================================================>]
verify: Service converged
```

- create service with published port

ingress forwards traffic with changing source address

```
koketani: ~/g/g/k/p/swarm (swarm ?)$ docker service create -p 8080:8080 --name swarm-test swarm-test
image swarm-test:latest could not be accessed on a registry to record
its digest. Each node will access swarm-test:latest independently,
possibly leading to different nodes running different
versions of the image.

6ozvmu3z6hpklajyewzpxzr5s
overall progress: 1 out of 1 tasks
1/1: running   [==================================================>]
verify: Service converged
```

```
koketani: ~/g/g/k/p/swarm (swarm ?)$ docker inspect 99ov5zxl4lzf
[
    {
        "ID": "99ov5zxl4lzf13retxk73xnmv",
        "Version": {
            "Index": 39045
        },
        "CreatedAt": "2020-11-03T00:14:46.089289Z",
        "UpdatedAt": "2020-11-03T00:14:49.2852496Z",
        "Labels": {},
        "Spec": {
            "ContainerSpec": {
                "Image": "swarm-test:latest",
                "Init": false,
                "DNSConfig": {},
                "Isolation": "default"
            },
            "Resources": {
                "Limits": {},
                "Reservations": {}
            },
            "Placement": {},
            "ForceUpdate": 0
        },
        "ServiceID": "rnabnvai5v8d2rfy5ohxhg1i6",
        "Slot": 1,
        "NodeID": "hf2ff90i4qcu3a6qi7dilk5xg",
        "Status": {
            "Timestamp": "2020-11-03T00:14:49.2486644Z",
            "State": "running",
            "Message": "started",
            "ContainerStatus": {
                "ContainerID": "ab964981c3a3281cb55a3e1ad4614e13709e9c23035daf0882cd12331c95ddb6",
                "PID": 998,
                "ExitCode": 0
            },
            "PortStatus": {}
        },
        "DesiredState": "running",
        "NetworksAttachments": [
            {
                "Network": {
                    "ID": "ulmzb879tmdd10cxe5wwci7ts",
                    "Version": {
                        "Index": 38548
                    },
                    "CreatedAt": "2020-08-08T14:54:53.9948781Z",
                    "UpdatedAt": "2020-10-26T00:49:18.018155588Z",
                    "Spec": {
                        "Name": "ingress",
                        "Labels": {},
                        "DriverConfiguration": {},
                        "Ingress": true,
                        "IPAMOptions": {
                            "Driver": {}
                        },
                        "Scope": "swarm"
                    },
                    "DriverState": {
                        "Name": "overlay",
                        "Options": {
                            "com.docker.network.driver.overlay.vxlanid_list": "4096"
                        }
                    },
                    "IPAMOptions": {
                        "Driver": {
                            "Name": "default"
                        },
                        "Configs": [
                            {
                                "Subnet": "10.0.0.0/24",
                                "Gateway": "10.0.0.1"
                            }
                        ]
                    }
                },
                "Addresses": [
                    "10.0.0.70/24"
                ]
            }
        ]
    }
]
```

```
koketani: ~/g/g/k/p/swarm (swarm +)$ docker service logs -f swarm-test
swarm-test.1.pw1erwipcp2n@docker-desktop    | ping from 10.0.0.2:35544
swarm-test.1.pw1erwipcp2n@docker-desktop    | ping from 10.0.0.2:35546
```

- create service with bridge

```
koketani: ~/g/g/k/p/swarm (swarm +!)$ docker service create -p 8080:8080 --network bridge --name swarm-test swarm-test
image swarm-test:latest could not be accessed on a registry to record
its digest. Each node will access swarm-test:latest independently,
possibly leading to different nodes running different
versions of the image.

s9p5dibzt631ji3lyh4asa4b1
overall progress: 1 out of 1 tasks
1/1: running   [==================================================>]
verify: Service converged
```

```
koketani: ~/g/g/k/p/swarm (swarm +!)$ docker service ps swarm-test
ID                  NAME                IMAGE               NODE                DESIRED STATE       CURRENT STATE                ERROR               PORTS
hnl71t8f53e0        swarm-test.1        swarm-test:latest   docker-desktop      Running             Running about a minute ago
```

```
koketani: ~/g/g/k/p/swarm (swarm +!)$ docker inspect hnl71t8f53e0
[
    {
        "ID": "hnl71t8f53e07xbhzhtf0n228",
        "Version": {
            "Index": 39367
        },
        "CreatedAt": "2020-11-03T00:40:41.890576Z",
        "UpdatedAt": "2020-11-03T00:40:45.0211487Z",
        "Labels": {},
        "Spec": {
            "ContainerSpec": {
                "Image": "swarm-test:latest",
                "Init": false,
                "DNSConfig": {},
                "Isolation": "default"
            },
            "Resources": {
                "Limits": {},
                "Reservations": {}
            },
            "Placement": {},
            "Networks": [
                {
                    "Target": "aov6gi6zq1iidweh1re1nwh29"
                }
            ],
            "ForceUpdate": 0
        },
        "ServiceID": "s9p5dibzt631ji3lyh4asa4b1",
        "Slot": 1,
        "NodeID": "hf2ff90i4qcu3a6qi7dilk5xg",
        "Status": {
            "Timestamp": "2020-11-03T00:40:44.9743937Z",
            "State": "running",
            "Message": "started",
            "ContainerStatus": {
                "ContainerID": "dedd08106670a37b4c6382304722536d6a995f6724853a35ceb11e64255fd61a",
                "PID": 2527,
                "ExitCode": 0
            },
            "PortStatus": {}
        },
        "DesiredState": "running",
        "NetworksAttachments": [
            {
                "Network": {
                    "ID": "ulmzb879tmdd10cxe5wwci7ts",
                    "Version": {
                        "Index": 38548
                    },
                    "CreatedAt": "2020-08-08T14:54:53.9948781Z",
                    "UpdatedAt": "2020-10-26T00:49:18.018155588Z",
                    "Spec": {
                        "Name": "ingress",
                        "Labels": {},
                        "DriverConfiguration": {},
                        "Ingress": true,
                        "IPAMOptions": {
                            "Driver": {}
                        },
                        "Scope": "swarm"
                    },
                    "DriverState": {
                        "Name": "overlay",
                        "Options": {
                            "com.docker.network.driver.overlay.vxlanid_list": "4096"
                        }
                    },
                    "IPAMOptions": {
                        "Driver": {
                            "Name": "default"
                        },
                        "Configs": [
                            {
                                "Subnet": "10.0.0.0/24",
                                "Gateway": "10.0.0.1"
                            }
                        ]
                    }
                },
                "Addresses": [
                    "10.0.0.117/24"
                ]
            },
            {
                "Network": {
                    "ID": "aov6gi6zq1iidweh1re1nwh29",
                    "Version": {
                        "Index": 38549
                    },
                    "CreatedAt": "2020-08-08T14:54:53.9949416Z",
                    "UpdatedAt": "2020-10-26T00:49:18.018905025Z",
                    "Spec": {
                        "Name": "bridge",
                        "Labels": {
                            "com.docker.swarm.predefined": "true"
                        },
                        "DriverConfiguration": {
                            "Name": "bridge"
                        },
                        "Scope": "swarm"
                    },
                    "DriverState": {
                        "Name": "bridge"
                    },
                    "IPAMOptions": {
                        "Driver": {}
                    }
                }
            }
        ]
    }
]
```

```
koketani: ~/g/g/k/p/swarm (swarm +!)$ docker service logs -f swarm-test
swarm-test.1.hnl71t8f53e0@docker-desktop    | ping from 10.0.0.2:35562
swarm-test.1.hnl71t8f53e0@docker-desktop    | ping from 10.0.0.2:35564
```

- create service with host network

```
koketani: ~/g/g/k/p/swarm (swarm ?)$ docker service create -p 8080:8080 --network host --name swarm-test swarm-test
image swarm-test:latest could not be accessed on a registry to record
its digest. Each node will access swarm-test:latest independently,
possibly leading to different nodes running different
versions of the image.

p1pn0blvk222q3pjiha7qfk1x
overall progress: 0 out of 1 tasks
1/1: container cannot be disconnected from host network or connected to host ne…
```

- expose container to external network

1. `ip link set dev ens1 promisc on`
2. `docker network create --config-only --subnet 172.80.0.0/16 --gateway 172.80.0.1 -o parent=ens1 --ip-range 10.90.1.0/24 external_net`
3. `docker network create -d macvlan --scope swarm --config-from external_net_config external_net`

https://gist.github.com/thaJeztah/83e7469c85bac28ae90b5178a4919301

- swarm with docker stack

```
koketani: ~/g/g/k/p/swarm (swarm +?)$ cat docker-compose.yml
---

version: '3.8'

services:
  ping:
    build: .
    image: koketani/ping
    ports:
      - '8080:8080'
koketani: ~/g/g/k/p/swarm (swarm +?)$ docker-compose build
Building ping
Step 1/7 : FROM golang:1.15.3-alpine as builder
 ---> d099254f5fc3
Step 2/7 : WORKDIR /go/src/swarm-test
 ---> Using cache
 ---> 43148bc01c8f
Step 3/7 : COPY main.go .
 ---> 52e6e786737f
Step 4/7 : RUN go install
 ---> Running in 6b45c3769c71
Removing intermediate container 6b45c3769c71
 ---> 54c80ac580ba

Step 5/7 : FROM alpine:3.12
 ---> d6e46aa2470d
Step 6/7 : COPY --from=builder /go/bin/swarm-test /bin/
 ---> Using cache
 ---> 33211e06a84d
Step 7/7 : ENTRYPOINT ["/bin/swarm-test"]
 ---> Using cache
 ---> 6ae7f577836e

Successfully built 6ae7f577836e
Successfully tagged koketani/ping:latest
koketani: ~/g/g/k/p/swarm (swarm +?)$ docker stack deploy --compose-file docker-compose.yml swarm-test
Ignoring unsupported options: build

Creating network swarm-test_default
Creating service swarm-test_ping
koketani: ~/g/g/k/p/swarm (swarm +?)$ docker service logs -f swarm-test_ping
swarm-test_ping.1.m12dgswv5zq2@docker-desktop    | ping from 10.0.0.2:35634

```

```
koketani: ~/g/g/k/p/swarm (swarm +?)$ docker stack services swarm-test
ID                  NAME                MODE                REPLICAS            IMAGE                  PORTS
v8m4ygjic0hd        swarm-test_ping     replicated          1/1                 koketani/ping:latest   *:8080->8080/tcp
koketani: ~/g/g/k/p/swarm (swarm +!?)$ docker stack ps swarm-test
ID                  NAME                IMAGE                  NODE                DESIRED STATE       CURRENT STATE           ERROR               PORTS
m12dgswv5zq2        swarm-test_ping.1   koketani/ping:latest   docker-desktop      Running             Running 2 minutes ago
koketani: ~/g/g/k/p/swarm (swarm +!?)$ curl localhost:8080
pong
```

```
koketani: ~/g/g/k/p/swarm (swarm +!?)$ docker inspect m12dgswv5zq2
[
    {
        "ID": "m12dgswv5zq2lhxppk7gyeq5e",
        "Version": {
            "Index": 39435
        },
        "CreatedAt": "2020-11-03T00:50:42.467708Z",
        "UpdatedAt": "2020-11-03T00:50:46.2879866Z",
        "Labels": {},
        "Spec": {
            "ContainerSpec": {
                "Image": "koketani/ping:latest",
                "Labels": {
                    "com.docker.stack.namespace": "swarm-test"
                },
                "Privileges": {
                    "CredentialSpec": null,
                    "SELinuxContext": null
                },
                "Isolation": "default"
            },
            "Resources": {},
            "Placement": {},
            "Networks": [
                {
                    "Target": "l7ojvahwjpctof1qy2zsmo6iu",
                    "Aliases": [
                        "ping"
                    ]
                }
            ],
            "ForceUpdate": 0
        },
        "ServiceID": "v8m4ygjic0hd7maqgfzpdxcv8",
        "Slot": 1,
        "NodeID": "hf2ff90i4qcu3a6qi7dilk5xg",
        "Status": {
            "Timestamp": "2020-11-03T00:50:46.2271106Z",
            "State": "running",
            "Message": "started",
            "ContainerStatus": {
                "ContainerID": "64579e431bd0194e5139ca3bbaafd71818e5d8c8f2abcc3ae1d6fccb45d082d5",
                "PID": 3769,
                "ExitCode": 0
            },
            "PortStatus": {}
        },
        "DesiredState": "running",
        "NetworksAttachments": [
            {
                "Network": {
                    "ID": "ulmzb879tmdd10cxe5wwci7ts",
                    "Version": {
                        "Index": 38548
                    },
                    "CreatedAt": "2020-08-08T14:54:53.9948781Z",
                    "UpdatedAt": "2020-10-26T00:49:18.018155588Z",
                    "Spec": {
                        "Name": "ingress",
                        "Labels": {},
                        "DriverConfiguration": {},
                        "Ingress": true,
                        "IPAMOptions": {
                            "Driver": {}
                        },
                        "Scope": "swarm"
                    },
                    "DriverState": {
                        "Name": "overlay",
                        "Options": {
                            "com.docker.network.driver.overlay.vxlanid_list": "4096"
                        }
                    },
                    "IPAMOptions": {
                        "Driver": {
                            "Name": "default"
                        },
                        "Configs": [
                            {
                                "Subnet": "10.0.0.0/24",
                                "Gateway": "10.0.0.1"
                            }
                        ]
                    }
                },
                "Addresses": [
                    "10.0.0.125/24"
                ]
            },
            {
                "Network": {
                    "ID": "l7ojvahwjpctof1qy2zsmo6iu",
                    "Version": {
                        "Index": 39426
                    },
                    "CreatedAt": "2020-11-03T00:50:39.2655052Z",
                    "UpdatedAt": "2020-11-03T00:50:39.268749Z",
                    "Spec": {
                        "Name": "swarm-test_default",
                        "Labels": {
                            "com.docker.stack.namespace": "swarm-test"
                        },
                        "DriverConfiguration": {
                            "Name": "overlay"
                        },
                        "Scope": "swarm"
                    },
                    "DriverState": {
                        "Name": "overlay",
                        "Options": {
                            "com.docker.network.driver.overlay.vxlanid_list": "4098"
                        }
                    },
                    "IPAMOptions": {
                        "Driver": {
                            "Name": "default"
                        },
                        "Configs": [
                            {
                                "Subnet": "10.0.2.0/24",
                                "Gateway": "10.0.2.1"
                            }
                        ]
                    }
                },
                "Addresses": [
                    "10.0.2.3/24"
                ]
            }
        ]
    }
]
```
