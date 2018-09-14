package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
	"github.com/ytake/laravel-jp-conference-api/design/medias"
	"net/http"
)

var _ = API("adder", func() {
	Title("The adder API")
	Description("A teaser for goa")
	Host("localhost:8080")
	Scheme("http")
})

var _ = Resource("conference", func() {
	Origin("*", func() {
		Methods("GET")
		Headers("Origin", "X-Requested-With", "Content-Type", "Accept")
	})
	DefaultMedia(medias.HalMedia)
	Action("info", func() {
		Routing(GET("info"))
		Description(" the response body")
		Response(OK, func(){
			Description("This is the success response")
		})

		Response("TimetableNotFound", func() {
			Description("Not found response")
			Media(medias.NotFoundMedia)
			Status(http.StatusNotFound)
		})
	})
})
