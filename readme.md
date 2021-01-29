# API - Investio

Run a dev server
```bash
$ go run main.go
```

To build a docker image
```bash
$ docker build . -t dewkul/inv-api
$ docker tag [imageTag] dewkul/inv-api:[version]
$ docker push [username]/inv-api:[version]
```