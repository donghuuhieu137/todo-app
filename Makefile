gen-proto:
	protoc proto/todo/todo.proto --go_out=paths=source_relative,plugins=grpc:. --grpc-gateway_out=logtostderr=true,paths=source_relative:.
run-server:
	go run ./cmd/server/main.go
run-proxy:
	go run ./cmd/proxy/proxy.go