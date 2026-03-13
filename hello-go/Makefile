IMAGE_MULTI := hello-go:multi
IMAGE_SIMPLE := hello-go:simple

.PHONY: build build-multi build-simple clean

build: build-multi build-simple

build-multi:
	docker build -f Dockerfile -t $(IMAGE_MULTI) .

build-simple:
	docker build -f Simple.dockerfile -t $(IMAGE_SIMPLE) .

clean:
	docker rmi $(IMAGE_MULTI) $(IMAGE_SIMPLE)
