# API - Investio

Run a dev server
```bash
$ go run main.go
```

To build a docker image
```bash
$ docker build . -t dewkul/inv-api
$ docker tag dewkul/inv-api dewkul/inv-api:[version]
$ docker push dewkul/inv-api:[version]
```