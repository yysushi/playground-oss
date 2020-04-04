# Note for Niginx

## Installation

```shell-session
koketani:nginx (master>)$ docker image ls | grep nginx
nginx                                               latest              ed21b7a8aee9        4 days ago          127MB
koketani:nginx (master %>)$ mkdir html
koketani:nginx (master %>)$ cat html/index.html
<html>
<header><title>This is title</title></header>
<body>
Hello world
</body>
</html>
koketani:nginx (master %>)$ docker run --name some-nginx -p 8080:80 -v $(pwd)/html:/usr/share/nginx/html:ro -d nginx
6fc0757c5ff7f762b40d1dbda598c6966c8eb9eda79b71bb4ab1752ebae0370d
koketani:nginx (master %>)$ curl localhost:8080
<html>
<header><title>This is title</title></header>
<body>
Hello world
</body>
</html>
```

```shell-session
koketani:nginx (master *>)$ docker cp some-nginx:/etc/nginx/. etc/nginx
```

## Logging

```shell-session
root@fc739aad7231:/# ls -l /var/log/nginx/
total 0
lrwxrwxrwx 1 root root 11 Mar 31 03:19 access.log -> /dev/stdout
lrwxrwxrwx 1 root root 11 Mar 31 03:19 error.log -> /dev/stderr
```

```shell-session
koketani:nginx (master>)$ docker logs fc739aad723155af83738f3c656d1fa31cfa74da8278e724f1347ddec6005b74
172.17.0.1 - - [04/Apr/2020:13:23:31 +0000] "GET /index.html HTTP/1.1" 200 88 "-" "curl/7.54.0" "-"
172.17.0.1 - - [04/Apr/2020:13:24:03 +0000] "GET /index.html HTTP/1.1" 200 88 "-" "curl/7.54.0" "-"
172.17.0.1 - - [04/Apr/2020:13:24:04 +0000] "GET /index.html HTTP/1.1" 200 88 "-" "curl/7.54.0" "-"
```

## Documentation

<https://nginx.org/en/docs/>

<https://nginx.org/en/docs/beginners_guide.html>
