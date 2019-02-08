#!/usr/bin/env bash
	yarn protoc-gen-grpc \
        --js_out="import_style=commonjs,binary:./generated" \
        --grpc_out="./generated" \
        --proto_path=../pkg/api/proto \
        ../pkg/api/proto/payment.proto

    yarn protoc-gen-grpc-ts \
        --ts_out=service=true:./generated \
        --proto_path=../pkg/api/proto \
        ../pkg/api/proto/payment.proto
