.PHONY: build lint

clean:
	rm build/app && rm build/app-linux

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

lint: ## Run all the linters
	docker run --rm -v $(PWD):/app -w /app golangci/golangci-lint:v1.59.1 golangci-lint run ./...

run-linux:
	make build-linux && cd ./build && ./app-linux -p 3000 -t ../static -c ../config/config.yaml

run-local:
	make build && cd ./build && ./app -p 3000 -t ../static -c ../config/config.yaml


#run app and db in docker
compose-up:
	cd deployment &&  docker rmi deployment-app -f  && docker-compose -f compose.yaml up -d

compose-down:
	cd deployment && docker-compose -f compose.yaml down


postgres-up:
	cd deployment && docker-compose -f postgres.yaml up -d

postgres-down:
	cd deployment && docker-compose -f postgres.yaml down
