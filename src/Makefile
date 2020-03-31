.PHONY: proto data run

run:
	echo "hello"
	@echo "hello world"
	-echo "hello one"

build:
	go build -o goms1 main.go

proto:
	for f in **/proto/*.proto; do \
		protoc --go_out=plugins=grpc:. $$f; \
		echo compiled: $$f; \
	done

