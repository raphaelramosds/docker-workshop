GO=go

.PHONY: consumer producer

EXECS = bin/consumer bin/producer

consumer:
	$(GO) build -o bin/consumer consumer/main.go

producer:
	$(GO) build -o bin/producer producer/main.go

all: consumer producer

clean:
	rm -rf $(EXECS) 