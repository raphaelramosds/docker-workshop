# hello-go

Very basic Go example with a multi-stage Docker build.

## Build image

```bash
docker build -t hello-go .
```

## Run container

```bash
docker run --rm hello-go
```

Expected output:

```text
Hello from Go in Docker multi-stage build!
```
