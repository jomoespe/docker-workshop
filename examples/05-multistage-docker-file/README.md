# Build container from multi-stage Dockerfile

This example creates a container in two stages. First one compiles full independent binary program from a [Go program](hello-world.go). The second stage builds a container from a scrath with only this program.

> The example is the same than [the minimal container example](../01-minimal-container/).

## Building the image

```bash
docker build --tag hello-world-multistage-minimal .
```

## Running a container

```bash
# No args
docker run --rm hello-world-multistage-minimal

# With args
docker run --rm hello-world-multistage-minimal "Data Team"
```
