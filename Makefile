run:
	go run cmd/main.go

swag-gen:
	swag init -g api/router.go -o api/docs

docker-build:
	sudo docker build -t api .

docker-run:
	sudo docker run --name iman_api -p 8080:8080 api