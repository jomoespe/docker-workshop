# Long running example

This is an example of a long-running container. It will start an HTTP server

## Building the image

```bash
docker build -t long-running-container .
```

## Running a container

```bash
# No args
docker run --rm -p 8080:80 long-running-container
```

> Mapping the ports with `-p 8080:80` argument.

Then you can access to root resource:

```bash
$ curl localhost:8080
server instance name: unamed
Hello !

$ curl localhost:8080?name=Demo
server instance name: unamed
Hello Demo!
```

## Running two instances

Here you can see how two containers are opening same port, **80**, but *networking namespaces* allow it with no conflicts. Then this port is mapped to a host port, and this must be unique.

```bash
docker run --detach --rm -p 8080:80 long-running-container instance-1

docker run --detach --rm -p 8081:80 long-running-container instance-2
```

> The `--detach` argument detached execution from terminal. 

And then you access each instance by its port

```bash
$ curl localhost:8080
server instance name: instance-1
Hello !

$ curl localhost:8081
server instance name: instance-2
Hello !
```

To see the containers running use the `docker ps` command.

```bash
docker ps
CONTAINER ID   IMAGE                    COMMAND                  CREATED          STATUS          PORTS                                   NAMES
847c432a266c   long-running-container   "/helloworldd instan…"   5 seconds ago    Up 2 seconds    0.0.0.0:8081->80/tcp, :::8081->80/tcp   inspiring_torvalds
d593d606ce35   long-running-container   "/helloworldd instan…"   12 seconds ago   Up 10 seconds   0.0.0.0:8080->80/tcp, :::8080->80/tcp   admiring_maxwell
```

To stop the containers `docker stop CONTAINER_ID`
