package medias

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
	"github.com/laravel-jp-conference/official-api/design/types"
)

var HalTimetableMedia = MediaType("HalTimetableMedia", func() {
	Description("for hypermedia application language")
	ContentType("application/laravel-jp-conference.hal+json; charset=utf-8") // Override default Content-Type header value

	Reference(types.TimetableType)
	Required("_links", "_embedded", "title", "body", "from_to", "place")
	Attributes(func() {
		Attribute("title", String, "講演タイトル")
		Attribute("body", String, "講演内容")
		Attribute("from_to", String, "講演時間")
		Attribute("place", String, "ホール")
		Attribute("_links")
		Attribute("_embedded")
	})
	View("default", func() {
		Attribute("title", String, "講演タイトル")
		Attribute("body", String, "講演内容")
		Attribute("from_to", String, "講演時間")
		Attribute("place", String, "ホール")
		Attribute("_links")
		Attribute("_embedded")
	})
})

var NotFoundMedia = MediaType("application/vnd.error+json", func() {
	Description("for hypermedia application language")
	ContentType("application/vnd.error+json; charset=utf-8") // Override default Content-Type header value

	Reference(types.VndErrorType)
	Required("message", "path", "logref", "_links")
	Attributes(func() {
		Attribute("message", String, "error message")
		Attribute("path", String, "error request uri")
		Attribute("logref", Integer, "code")
		Attribute("_links")
	})
	View("default", func() {
		Attribute("message", String)
		Attribute("path", Integer)
		Attribute("logref", String)
		Attribute("_links")
	})
})
