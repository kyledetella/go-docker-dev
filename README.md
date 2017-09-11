# go-docker-dev

A simple setup to quickly prototype and develop Go applications inside of a Docker container.

This is based on [this blog post](https://medium.com/statuscode/golang-docker-for-development-and-production-ce3ad4e69673) and a more full-featured version can be found in [this repo](https://github.com/McMenemy/GoDoRP).

## Usage

### Build container

```bash
docker build -t kyledetella/go-docker-dev .
```

### Start container

```bash
docker run -i -t -p 8080:8080 \
  -v `pwd`/app:go/src/github.com/kyledetella/go-docker-dev/app \
  kyledetella/go-docker-dev
```

and then `curl localhost:8080`
