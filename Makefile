NAME=quay.io/tomkral/guestbook-demo-backend
VERSION=v1

all: image push

build: server.go
	go build server.go

image:
	docker build -t $(NAME):$(VERSION) .

push:
	docker push $(NAME):$(VERSION)

deploy:
	kedge apply -f Kedge/