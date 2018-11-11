#!/bin/bash

set -ex


CURDIR=$PWD

docker run --rm -e GOBIN=/go/bin/ -v "$PWD"/bin:/go/bin/ -v "$PWD":/go/src/github.com/fananchong/go-redis-orm.v2 -w /go/src/github.com/fananchong/go-redis-orm.v2 golang go install ./tools/...

./bin/redis2go --input_dir=./example/redis_def --output_dir=./example --package=main

docker run --rm -e GOBIN=/go/bin/ -v "$PWD"/bin:/go/bin/ -v "$PWD":/go/src/github.com/fananchong/go-redis-orm.v2 -w /go/src/github.com/fananchong/go-redis-orm.v2 golang go install ./example/...


if [ "$1" = "publish" ]; then
	docker build -t redis2go .
	docker tag redis2go:latest fananchong/redis2go:latest
	docker push fananchong/redis2go:latest
fi