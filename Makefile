GoModule = github.com/aesoper101/kitextest

.PHONY: gen gen-server
gen:
	@echo "generate"
	@make gen-server
	@make gen-client-api
	@make gen-client-stream-api

.PHONY: gen-server
gen-server:
	kitex -module $(GoModule) -service test -no-fast-api   ./idl/trans.proto

.PHONY: gen-client-api
gen-client-api:
	cd client && \
	protoc --proto_path=./    --go_out=paths=source_relative:./api  --go-http_out=paths=source_relative:./api  --go-grpc_out=paths=source_relative:./api ./trans.proto


.PHONY: gen-client-stream-api
gen-client-stream-api:
	cd streamclient && \
	protoc --proto_path=./    --go_out=paths=source_relative:./api  --go-http_out=paths=source_relative:./api  --go-grpc_out=paths=source_relative:./api ./trans.proto
