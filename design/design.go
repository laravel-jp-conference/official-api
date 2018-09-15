package design

import (
	. "github.com/goadesign/goa/design/apidsl"
	_ "github.com/ytake/laravel-jp-conference-api/design/resources"
)

var _ = API("adder", func() {
	Title("The adder API")
	Description("A teaser for goa")
	Host("localhost:8080")
	Scheme("http")
})
