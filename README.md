# Laravel JP Conference 2019 API

required Go >= 1.11, [goa](https://github.com/goadesign/goa)


## Goa Install

```bash
$ go get -u github.com/goadesign/goa/...
```

## For Develop

```bash
$ $GOPATH/src/github.com/laravel-jp-conference
$ git clone https://github.com/laravel-jp-conference/official-api
```

### Generate Code

```bash
$ make code-gen
```

### Build

```bash
$ make prod-build
```

### Need Swagger?

```bash
$ make swagger-gen
$ go run main.go -needSwagger
```

[Swagger UI](http://localhost:8080/swagger/)
