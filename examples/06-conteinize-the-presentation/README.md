# Conteinize this presentation

In this example we will create a Docker container image with this presentation. When you run it, the container will serve it via HTTP.

## Building the image

In this example the build context will be the `presentation/` directory in the project root. We use the `--file` argument to set the Dockerfile to use, because it's not in the build context (the presentation directory).

```bash
docker build --tag docker-workshop-presentation --file ./Dockerfile ../../presentation
```

## Running a container

We create and start a container from the image, mapping the TCP/80 port to TCP/8080 port.

```bash
docker run --rm --publish 8080:80 --name amazing_docker_presentation docker-workshop-presentation
```

> We set the container name with the `--name` argument

Once the container is running yo can access to `http://localhost:8080` to see the presentation.
