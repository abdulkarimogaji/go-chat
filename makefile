protoc:
	rm -rf ./pb && mkdir pb && protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	proto/*.proto

evans:
	evans --host localhost --port 8080 -r repl

db:
	docker run --name go-chat -p 3000:3306 -e MYSQL_ROOT_PASSWORD=password -e MYSQL_DATABASE=go_chat -d mysql:8.0.31

server:
	go run main.go

db_terminal:
	docker exec -it go-chat mysql -u root --protocol=tcp -p

.PHONY: protoc evans db server db_terminal