#! /usr/bin/make

.PHONY: code-gen
code-gen:
	goagen app -d github.com/ytake/laravel-jp-conference-api/design
	goagen controller -d github.com/ytake/laravel-jp-conference-api/design -o action
