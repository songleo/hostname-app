# hostname-app

The hostname-app is a simple http server to show the hostname and version:

```
$ export VERSION=v1.0 && go run main.go
2021/12/23 16:33:07 hostname server running on 0.0.0.0:8080 ...
2021/12/23 16:33:10 method=GET url=/ protocol=HTTP/1.1 status_code=200
2021/12/23 16:33:10 method=GET url=/favicon.ico protocol=HTTP/1.1 status_code=200
```

go to http://localhost:8080/ :

```
hostname: host1
app version: v1.0
```
