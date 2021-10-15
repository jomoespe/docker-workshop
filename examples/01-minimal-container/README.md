# Minimal Container

This example creates a minimun container image. It only contains a binary that say *Hello world!*. Also this example shows how to send arguments to conteinized process.

## Building the image

```bash
docker build -t hello-world-minimal .
```

## Running a container

```bash
# No args
$ docker run hello-world-minimal 
Hello world!

# With args
$ docker run hello-world-minimal "Data Team"
Hello Data Team!
```

Everytime you run the image it will create an container. You can see your computer containers with the `docker ps` command. This will show the running containers. To see also the stopped containers use the `--all` argument.

```bash
$ docker ps --all
CONTAINER ID   IMAGE                 COMMAND                  CREATED             STATUS                         PORTS     NAMES
6b87e3a29c94   hello-world-minimal   "/hello-world 'Data …"   2 minutes ago       Exited (0) 2 minutes ago                 mystifying_hamilton
5b343f7323ba   hello-world-minimal   "/hello-world"           2 minutes ago       Exited (0) 2 minutes ago                 jolly_chatterjee
4da69020286c   hello-world-minimal   "/hello-world pepe"      About an hour ago   Exited (0) About an hour ago             strange_bell
d7ce3018e968   hello-world-minimal   "/hello-world"           About an hour ago   Exited (0) About an hour ago             charming_curran
```
