package action

import (
	"github.com/goadesign/goa"
	"github.com/ytake/laravel-jp-conference-api/app"
	"github.com/ytake/laravel-jp-conference-api/foundation"
	"net/url"
)

// TimetableController implements the timetable resource.
type TimetableController struct {
	*goa.Controller
}

// NewTimetableController creates a timetable controller.
func NewTimetableController(service *goa.Service) *TimetableController {
	return &TimetableController{Controller: service.NewController("TimetableController")}
}

// Table runs the table action.
func (c *TimetableController) Table(ctx *app.TableTimetableContext) error {
	// TimetableController_Table: start_implement

	// Put your logic here
	return ctx.TimetableNotFound(&app.Error{
		Links: &app.VndLinkType{
			About: &app.HalLinkUnitType{
				Href: foundation.ConferenceURL,
			},
		},
		Logref:  44,
		Message: foundation.ErrorMessage,
		Path:    resolveRequestURI(ctx).String(),
	})
	// TimetableController_Table: end_implement
}

func resolveRequestURI(ctx *app.TableTimetableContext) (uv *url.URL) {
	uv, _ = url.Parse(ctx.URL.String())
	uv.Scheme = string("https")
	uv.Path = ctx.URL.Path
	uv.Host = ctx.Host
	uv.RawQuery = ""
	return
}
