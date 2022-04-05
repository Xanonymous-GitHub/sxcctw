# Simple URL Shortener on Xanonymous Core Cell (xcc).

It is An URL shortener.

BUT STILL DEVELOPED IN PROGRESS.
Please go back here to see the changes later.

## TODOS
- [ ] Add redis.
- [ ] Add openAPI3 doc.
- [ ] Add k8s yaml.
- [ ] Consider not use base62 (KGS, or else...??)
- [ ] Add tests.

## Get started

1. Ensure using go 1.18.

```shell
# go version go1.18
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

## How to use? (for this time)
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
