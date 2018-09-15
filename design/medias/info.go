package medias

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
	"github.com/ytake/laravel-jp-conference-api/design/types"
)

var HalInfoMedia = MediaType("HalInfoMedia", func() {
	Description("for hypermedia application language")
	ContentType("application/laravel-jp-conference.hal+json; charset=utf-8") // Override default Content-Type header value

	Reference(types.PlaceInfoType)
	Required("_links", "_embedded", "name", "date")
	Attributes(func() {
		Attribute("name", String, "カンファレンス名")
		Attribute("date", String, "開催日時")
		Attribute("_links")
		Attribute("_embedded")
	})
	View("default", func() {
		Attribute("name", String, "カンファレンス名")
		Attribute("date", String, "開催日時")
		Attribute("_links")
		Attribute("_embedded")
	})
})
