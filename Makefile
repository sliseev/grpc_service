build:
	protoc --go_out=. --go-grpc_out=. --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative ./api/api.proto
	go build -o gserver cmd/server/main.go
	go build -o gclient cmd/client/main.go

clean:
	rm -f api/*.pb.go
	rm -f gserver
	rm -f gclient
