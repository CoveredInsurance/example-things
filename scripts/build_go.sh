#!/bin/sh

set -e -x

export GOPATH=`pwd`
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOBIN

go get github.com/mholt/archiver/cmd/archiver
wget https://github.com/google/protobuf/releases/download/v3.4.0/protoc-3.4.0-linux-x86_64.zip -O protoc.zip

archiver open protoc.zip protoc

export PATH=$PATH:`pwd`/protoc/bin
protoc --version

git clone things-source $GOPATH/src/github.com/CoveredInsurance/example-things
git clone grpc-go-source $GOPATH/src/google.golang.org/grpc

rm -rf $GOPATH/src/github.com/CoveredInsurance/example-things/**/*.pb.go
rm -rf $GOPATH/src/github.com/CoveredInsurance/example-things/**/*.pb.gw.go

go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
go get -u github.com/golang/protobuf/...
go get -u github.com/square/goprotowrap/cmd/protowrap

cd things-source/things
$GOPATH/bin/protowrap \
    -I. \
    -I/usr/local/include \
    -I$GOPATH/src \
    -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
    --proto_path=. \
    --go_out=plugins=grpc:$GOPATH/src \
    ./**/*.proto

# Swagger directory must exist before using it for output
mkdir -p $GOPATH/src/github.com/CoveredInsurance/example-things/swagger

$GOPATH/bin/protowrap \
    -I. \
    -I/usr/local/include \
    -I$GOPATH/src \
    -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
    --proto_path=. \
    --grpc-gateway_out=logtostderr=true:$GOPATH/src \
    --swagger_out=logtostderr=true:$GOPATH/src/github.com/CoveredInsurance/example-things/swagger \
    ./**/*.proto

cd $GOPATH/src/github.com/CoveredInsurance/example-things/swagger
rm -f *.swagger.json

cd $GOPATH/src/github.com/CoveredInsurance/example-things
git config --global user.email "devops@itscovered.com"
git config --global user.name "Concourse CICD"
git add .
git commit --allow-empty -m "CICD publish"
