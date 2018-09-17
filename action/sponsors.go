package action

import (
	"github.com/goadesign/goa"
	"github.com/ytake/laravel-jp-conference-api/app"
	"github.com/ytake/laravel-jp-conference-api/foundation"
)

// SponsorsController implements the sponsors resource.
type SponsorsController struct {
	*goa.Controller
}

// NewSponsorsController creates a sponsors controller.
func NewSponsorsController(service *goa.Service) *SponsorsController {
	return &SponsorsController{Controller: service.NewController("SponsorsController")}
}

// Sponsors runs the sponsors action.
func (c *SponsorsController) Sponsors(ctx *app.SponsorsSponsorsContext) error {
	// SponsorsController_Sponsors: start_implement

	// Put your logic here
	a := &foundation.RequestAction{Request: ctx.Request}
	// Put your logic here
	return ctx.SponsorNotFound(&app.Error{
		Links: &app.VndLinkType{
			About: &app.HalLinkUnitType{
				Href: foundation.ConferenceURL,
			},
		},
		Logref:  44,
		Message: foundation.ErrorMessage,
		Path:    a.ResolveRequestURI().String(),
	})
	// SponsorsController_Sponsors: end_implement
}
