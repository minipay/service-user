BINARY=engine
test:
	go test -count=1 -v -cover -covermode=atomic ./...

download:
	go mod download

vendor:
	go mod vendor

docker:
	docker build -t cleanbase .

run:
	docker-compose up -d

stop:
	docker-compose down