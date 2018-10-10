package resources

import (
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("swagger", func() {
	Origin("*", func() {
		Methods("GET")
	})
	Files("/swagger/swagger.json", "public/swagger/swagger.json")
	Files("/swagger/*filepath", "public/dist")
})
