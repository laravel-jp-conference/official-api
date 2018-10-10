package design

import (
	. "github.com/goadesign/goa/design/apidsl"
	_ "github.com/laravel-jp-conference/official-api/design/resources"
)

var _ = API("adder", func() {
	Title("The adder API")
	Description("fo Laravel JP Conference")
	Host("localhost:8080")
	Scheme("http")
})
