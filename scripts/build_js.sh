#!/bin/sh

set -e -x

npm install google-protobuf
npm install grpc-tools

./node_modules/grpc-tools/bin/protoc --version

git clone things-source updated-things-source

rm -rf updated-things-source/**/*_pb.js

mkdir -p updated-things-source/js

ls -la .

./node_modules/grpc-tools/bin/protoc \
    --proto_path=grpc-gateway/third_party/googleapis/ \
    --grpc_out=updated-things-source/js \
    --plugin=protoc-gen-grpc=./node_modules/grpc-tools/bin/grpc_node_plugin \
    --js_out=import_style=commonjs,binary:updated-things-source \
    ./grpc-gateway/third_party/googleapis/google/api/annotations.proto \
    ./grpc-gateway/third_party/googleapis/google/api/http.proto

./node_modules/grpc-tools/bin/protoc \
    -Igrpc-gateway/third_party/googleapis \
    --proto_path=things-source/things \
    --grpc_out=updated-things-source/js \
    --plugin=protoc-gen-grpc=./node_modules/grpc-tools/bin/grpc_node_plugin \
    --js_out=import_style=commonjs,binary:updated-things-source/js \
    ./things-source/things/*.proto

cd updated-things-source/js
sed -i 's/require(\([[:punct:]][[:punct:]]\+.*\));/require(require("path").join(__dirname, \1));/' ./*.js
sed -i 's/require(\([[:punct:]][[:punct:]]\+.*\));/require(require("path").join(__dirname, \1));/' ./google/api/annotations_pb.js

git config --global user.email "devops@itscovered.com"
git config --global user.name "Concourse CICD"

git add .
git commit --allow-empty -m "CICD publish js"
