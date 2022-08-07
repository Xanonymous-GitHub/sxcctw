# Simple URL Shortener on Xanonymous Core Cell.

It is An URL shortener.

[![Service Build](https://github.com/Xanonymous-GitHub/sxcctw/actions/workflows/go.yml/badge.svg)](https://github.com/Xanonymous-GitHub/sxcctw/actions/workflows/go.yml)
[![Web Build](https://github.com/Xanonymous-GitHub/sxcctw/actions/workflows/web.yml/badge.svg)](https://github.com/Xanonymous-GitHub/sxcctw/actions/workflows/web.yml)
[![API Image](https://github.com/Xanonymous-GitHub/sxcctw/actions/workflows/docker-publish-sxcctw-api.yml/badge.svg)](https://github.com/Xanonymous-GitHub/sxcctw/actions/workflows/docker-publish-sxcctw-api.yml)
[![DB Image](https://github.com/Xanonymous-GitHub/sxcctw/actions/workflows/docker-publish-sxcctw-db.yml/badge.svg)](https://github.com/Xanonymous-GitHub/sxcctw/actions/workflows/docker-publish-sxcctw-db.yml)

BUT STILL DEVELOPED IN PROGRESS. Please go back here to see the changes later.

## Try it now

1. Create your shortened URL.

```shell
curl --location --request POST 'https://s.xcc.tw/api/url' \
--header 'Content-Type: application/json' \
--data-raw '{
    "originUrl": "https://www.dcard.tw/f/mood/p/238545781?cid=306EA5A4-D67F-41A8-903F-1456E8FDE547",
    "expireAt": "2034-11-12T11:45:26.371Z"
}'
```

2. Then use it

```shell
# Go to this URL.
https://s.xcc.tw/v5oSt3VFD1l
```

## TODOS

- [ ] Add redis. (In progress...)
- [ ] Add openAPI3 doc.
- [ ] Add k8s yaml.
- [ ] Add LB, Grafana, ArgoCD
- [ ] Consider not use base62 (KGS, or else...??)
- [ ] Add tests.
- [ ] cronjob (maybe)

## Get started

1. Ensure using go 1.19.

```shell
# go version go1.19
go version
```

2. install dependencies.

```shell
go install github.com/silenceper/gowatch@latest
```

If you want to re-generate the gRPC files, please install these

```shell
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest \
&& go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest \
&& go install entgo.io/contrib/entproto/cmd/protoc-gen-entgrpc@latest \
```

And do this:

```shell
go generate ./...
```

3. start a MySQL DB.

```shell
# For example, use docker
docker run --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=root -e MYSQL_DATABASE=sxcctw -d mysql
```

4. start `sxcctw db` server.

```shell
./scripts/watchgo.sh sxcctw_db
```

5. start `sxcctw api` server.

```shell
./scripts/watchgo.sh sxcctw_api
```

## How to use? (for development)

(openAPI doc TBD)

#### Create new ID. (201 Created)

```shell
# Create a shortened id, which is from `https://ntut/club`,
# and set expire time at `2034-11-12T11:45:26.371Z`.
curl --location --request POST 'localhost:8080/api/url' \
--header 'Content-Type: application/json' \
--data-raw '{
    "originUrl": "https://ntut/club",
    "expireAt": "2034-11-12T11:45:26.371Z"
}'
```

#### Get URL by existed ID. (200 OK / 404 Not Found / 410 Gone)

```shell
# Get an existing shortened id (ZNkw23qpiss).
curl --location --request GET 'localhost:8080/api/url?id=ZNkw23qpiss'
```

## Why I choose (TBD)

- MySQL
- Gin
- Gorm
- gRPC
- Clean Architecture
