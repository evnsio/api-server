IMAGE := simple-server

build:
	docker build -t $(IMAGE) .

run:
	docker run -ti -p 8080:8080 $(IMAGE)

build-and-run: build run