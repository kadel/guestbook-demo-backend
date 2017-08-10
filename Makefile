NAME=tomaskral/kedge-demo-backend
VERSION=v1

all: image push

build: server.go
	go build server.go

image:
	docker build -t $(NAME):$(VERSION) .

push:
	docker push $(NAME):$(VERSION)
