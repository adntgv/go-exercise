build:
	docker build -t service .

run:
	docker run service

test:
	go test ./...