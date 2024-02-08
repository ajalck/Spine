BINARY_NAME=spine

build:
	go build -o ${BINARY_NAME} cmd/main.go

run: build
	./${BINARY_NAME}

wire: 
	wire ./pkg/di