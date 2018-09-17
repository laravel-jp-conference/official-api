package medias

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
	"github.com/ytake/laravel-jp-conference-api/design/types"
)

var HalSponsorMedia = MediaType("HalSponsorMedia", func() {
	Description("for hypermedia application language")
	ContentType("application/laravel-jp-conference.hal+json; charset=utf-8") // Override default Content-Type header value

	Reference(types.SponsorResponseType)
	Required("_links", "_embedded", "count")
	Attributes(func() {
		Attribute("_links")
		Attribute("_embedded")
		Attribute("count", Integer, "スポンサー数")
	})
	View("default", func() {
		Attribute("_links")
		Attribute("_embedded")
		Attribute("count")
	})
})
