package resources

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
	"github.com/laravel-jp-conference/official-api/design/medias"
	"net/http"
)

var _ = Resource("info", func() {
	Origin("*", func() {
		Methods("GET")
		Headers("Origin", "X-Requested-With", "Content-Type", "Accept")
	})
	DefaultMedia(medias.HalInfoMedia)
	Action("info", func() {
		Routing(GET("info"))
		Description("the response body")
		Response(OK, func() {
			Description("This is the success response")
		})

		Response(NotFound, func() {
			Description("Not found response")
			Status(http.StatusNotFound)
		})
	})
})
