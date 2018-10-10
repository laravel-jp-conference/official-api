//go:generate goagen bootstrap -d github.com/laravel-jp-conference/official-api/design

package main

import (
	"flag"
	"fmt"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/laravel-jp-conference/official-api/action"
	"github.com/laravel-jp-conference/official-api/app"
	"github.com/laravel-jp-conference/official-api/foundation"
)

type AppConfig struct {
	Port        *string
	NeedSwagger *bool
}

func main() {
	// Create service
	service := goa.New("official-api")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	app.MountTimetableController(service, action.NewTimetableController(service))
	app.MountInfoController(service, action.NewInfoController(service))
	app.MountSponsorsController(service, action.NewSponsorsController(service))

	ao := retrievePortOption()
	if *ao.NeedSwagger {
		app.MountSwaggerController(service, action.NewSwaggerController(service))
	}

	// Start service
	if err := service.ListenAndServe(fmt.Sprintf(":%s", *ao.Port)); err != nil {
		service.LogError("startup", "err", err)
	}
}

func retrievePortOption() *AppConfig {
	var (
		port string
		needSwagger  bool
	)
	flag.StringVar(&port, "port", foundation.ApplicationPort, "port number option")
	flag.BoolVar(&needSwagger, "needSwagger", false, "need swagger ui")
	flag.Parse()
	return &AppConfig{&port, &needSwagger}
}
