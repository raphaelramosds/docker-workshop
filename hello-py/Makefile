IMAGE := hello-py:local

.PHONY: build clean

build:
	docker build -f Dockerfile -t $(IMAGE) .

clean:
	docker rmi $(IMAGE)
