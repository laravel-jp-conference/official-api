#! /usr/bin/make

.PHONY: code-gen
code-gen:
	goagen app -d github.com/ytake/laravel-jp-conference-api/design
	goagen controller -d github.com/ytake/laravel-jp-conference-api/design -o action

.PHONY: prod-build
prod-build: code-gen
	GOOS=linux GOARCH=amd64 go build -o=conference-api

.PHONY: swagger-gen
swagger-gen:
	goagen swagger --design github.com/ytake/laravel-jp-conference-api/design
