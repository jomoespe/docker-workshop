# Docker Workshop

[Link to presentation](presentation/)

To execute the presentation:

```bash
docker run \
  --rm \
  --publish 8080:80 \
  jomoespe/docker-workshop-presentation
```

Once the container is running, open your browser to [localhost:8080](http://localhost:8080)

## Examples

1. [Minimal image](examples/01-minimal-image)
2. [Container image from base image](examples/02-image-from-image)
3. [Mounting volumes](examples/03-mounting-volumes)
4. [Long running containers](examples/04-long-running-container)
5. [Multistage Dockerfile](examples/05-multistage-docker-file)
6. [Conteinize the presentation](examples/06-conteinize-the-presentation)

## References

* [Docker site](https://docker.com)
* [What even is a container: namespaces and cgroups](https://jvns.ca/blog/2016/10/10/what-even-is-a-container/)
* [Docker Internals. A Deep Dive Into Docker For Engineers Interested In The Gritty Detai](http://docker-saigon.github.io/post/Docker-Internals/)
