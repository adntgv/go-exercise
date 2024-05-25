build:
	docker build -t service .

run:
	docker run -p 80:80 service 

test:
	go test ./...