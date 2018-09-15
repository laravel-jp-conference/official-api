package action

import (
	"github.com/goadesign/goa"
	"github.com/ytake/laravel-jp-conference-api/app"
	"github.com/ytake/laravel-jp-conference-api/foundation"
)

// InfoController implements the info resource.
type InfoController struct {
	*goa.Controller
}

// NewInfoController creates a info controller.
func NewInfoController(service *goa.Service) *InfoController {
	return &InfoController{Controller: service.NewController("InfoController")}
}

// Info runs the info action.
func (c *InfoController) Info(ctx *app.InfoInfoContext) error {
	// InfoController_Info: start_implement

	// Put your logic here
	res := &app.Halinfomedia{
		Date:  foundation.ConferenceDate,
		Name:  foundation.ConferenceName,
		Links: &app.HalLinkType{Self: &app.HalLinkUnitType{Href: foundation.ConferenceURL}},
		Embedded: &app.EmbeddedPlaceResponseType{
			Places: &app.PlaceCollectType{
				Name:      foundation.HallName,
				Address:   foundation.HallAddress,
				Note:      foundation.FindHallNote(),
				Latitude:  foundation.HallLatitude,
				Longitude: foundation.HallLongitude,
				Links:     &app.HalLinkType{Self: &app.HalLinkUnitType{Href: foundation.HallURL}},
			},
		},
	}
	return ctx.OK(res)
	// InfoController_Info: end_implement
}
