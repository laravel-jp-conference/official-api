package action

import (
	"github.com/goadesign/goa"
	"github.com/ytake/laravel-jp-conference-api/app"
	"github.com/ytake/laravel-jp-conference-api/foundation"
	"net/url"
)

// ConferenceController implements the conference resource.
type ConferenceController struct {
	*goa.Controller
}

// NewConferenceController creates a conference controller.
func NewConferenceController(service *goa.Service) *ConferenceController {
	return &ConferenceController{Controller: service.NewController("ConferenceController")}
}

// Info runs the info action.
func (c *ConferenceController) Info(ctx *app.InfoConferenceContext) error {
	// ConferenceController_Info: start_implement

	// Put your logic here
	em := foundation.ErrorMessage
	ec := 44
	rp := resolveRequestURI(ctx).String()
	return ctx.TimetableNotFound(&app.Error{
		Links: &app.VndLinkType{
			About: &app.HalLinkUnitType{
				Href: "https://conference2019.laravel.jp/",
			},
		},
		Logref:  &ec,
		Message: &em,
		Path:    &rp,
	})
	// ConferenceController_Info: end_implement
}

func resolveRequestURI(ctx *app.InfoConferenceContext) (uv *url.URL) {
	uv, _ = url.Parse(ctx.URL.String())
	uv.Scheme = string("https")
	uv.Path = ctx.URL.Path
	uv.Host = ctx.Host
	uv.RawQuery = ""
	return
}
