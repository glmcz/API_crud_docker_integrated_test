.PHONY: build generate-proto lint



deps:
	go get ./...

test:
	go test -v ./...

build:
	CGO_ENABLED=0 \
		go build \
			-installsuffix cgo \
			-o ./build/app \
			-ldflags "-X main.Version=$(APP_VERSION)" \
			./cmd/*.go


build-linux:
	CGO_ENABLED=0 \
	GOOS=linux \
		go build \
			-installsuffix cgo \
			-o ./build/app-linux \
			-ldflags "-X main.Version=$(APP_VERSION)" \
			./cmd/*.go

fmt:
	gofmt -s -w . && gofumpt -w .

run:
	go run ./cmd/*.go


#run app and db in docker
compose-up:
	cd deployment &&  docker rmi deployment-app -f  && docker-compose -f compose.yaml up -d

compose-down:
	cd deployment && docker-compose -f compose.yaml down


postgres-up:
	cd deployment && docker-compose -f postgres.yaml up -d

postgres-down:
	cd deployment && docker-compose -f postgres.yaml down
