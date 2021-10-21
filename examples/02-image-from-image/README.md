# Container from image

This example creates an image from another image, **alpine:3.14.2**, that executes the command `echo "Hello from Alpine Linux"`.

## Building the image

```bash
docker build -t hello-from-image .
```

## Running a container

```bash
# No args
$ docker run --rm hello-from-image
Hello from Alpine Linux
```

> With the `--rm` argument the intermediate containers will be removed when container finish.
