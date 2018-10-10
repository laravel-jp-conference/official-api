package resources

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
	"github.com/laravel-jp-conference/official-api/design/medias"
	"net/http"
)

var _ = Resource("sponsors", func() {
	Origin("*", func() {
		Methods("GET")
		Headers("Origin", "X-Requested-With", "Content-Type", "Accept")
	})
	DefaultMedia(medias.HalSponsorMedia)
	Action("sponsors", func() {
		Routing(GET("sponsors"))
		Description("the response body")
		Response(OK, func() {
			Description("This is the success response")
		})

		Response("SponsorNotFound", func() {
			Description("Not found response")
			Media(medias.NotFoundMedia)
			Status(http.StatusNotFound)
		})
	})
})
