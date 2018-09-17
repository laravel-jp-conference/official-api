package action

import (
	"github.com/goadesign/goa"
	"github.com/ytake/laravel-jp-conference-api/app"
	"github.com/ytake/laravel-jp-conference-api/foundation"
)

// TimetableController implements the timetable resource.
type TimetableController struct {
	*goa.Controller
}

// NewTimetableController creates a timetable controller.
func NewTimetableController(service *goa.Service) *TimetableController {
	return &TimetableController{Controller: service.NewController("TimetableController")}
}

// Tables タイムテーブル情報
func (c *TimetableController) Tables(ctx *app.TablesTimetableContext) error {
	// TimetableController_Table: start_implement
	a := &foundation.RequestAction{Request: ctx.Request}
	// Put your logic here
	return ctx.TimetableNotFound(&app.Error{
		Links: &app.VndLinkType{
			About: &app.HalLinkUnitType{
				Href: foundation.ConferenceURL,
			},
		},
		Logref:  44,
		Message: foundation.ErrorMessage,
		Path:    a.ResolveRequestURI().String(),
	})
	// TimetableController_Table: end_implement
}
