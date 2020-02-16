# Note for docker ignore

dockerignore specifies to ignore files.
- build cache mechanism (even if file changes, keep using cache)
- copy mechanism (ignore the file)

```shell-session
koketani:ignore (master %=)$ ls
Dockerfile
koketani:ignore (master %=)$ cat .dockerignore
memo.md
gomi
koketani:ignore (master %=)$ touch gomi2
koketani:ignore (master %=)$ docker build -t ignore .
Sending build context to Docker daemon  3.584kB
Step 1/2 : FROM alpine
 ---> e7d92cdc71fe
Step 2/2 : COPY . .
 ---> 348533a4828c
Successfully built 348533a4828c
Successfully tagged ignore:latest
koketani:ignore (master %=)$ touch gomi
koketani:ignore (master %=)$ docker build -t ignore .
Sending build context to Docker daemon  3.584kB
Step 1/2 : FROM alpine
 ---> e7d92cdc71fe
Step 2/2 : COPY . .
 ---> Using cache
 ---> 348533a4828c
Successfully built 348533a4828c
Successfully tagged ignore:latest

ketani:ignore (master %=)$ docker run ignore ls | grep gomi
gomi2
```
