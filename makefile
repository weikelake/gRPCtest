clean:
	rm pkg/api/*.go
generate:
	protoc -I . --go_grpc_out ./pkg\
    --go_grpc_opt paths=source_relative \
    proto/rusprofile.proto
	protoc -I . --go_out ./pkg\
    --go_opt paths=source_relative \
    proto/rusprofile.proto
	protoc -I . --grpc-gateway_out ./pkg\
    --grpc-gateway_opt logtostderr=true \
    --grpc-gateway_opt paths=source_relative \
    proto/rusprofile.proto
	protoc -I . --openapiv2_out ./pkg \
    --openapiv2_opt logtostderr=true \
    proto/rusprofile.proto


