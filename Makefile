

gen:
	protoc --go_out=plugins=grpc:.. pkg/agent/proto/sandbox.proto
	protoc --go_out=plugins=grpc:.. pkg/api/proto/payment.proto

