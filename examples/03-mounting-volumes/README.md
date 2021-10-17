# Mounting volumes

TODO

## Building the image

```bash
docker build -t mounting-volumes .
```

## Running a container

```bash
$ docker run --rm -v $PWD/the-file.txt:/file.txt:ro mounting-volumes
This is the file in the host system
```

With the `--v` we mount the host file `the-file.txt` into the container, with the name `/file.txt`. With the `:ro` we mount it in read-only model, the container process can't modify the host file.
