//go:generate goagen bootstrap -d github.com/ytake/laravel-jp-conference-api/design

package main

import (
	"flag"
	"fmt"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/ytake/laravel-jp-conference-api/action"
	"github.com/ytake/laravel-jp-conference-api/app"
	"github.com/ytake/laravel-jp-conference-api/foundation"
)

type AppConfig struct {
	Port *string
}

func main() {
	// Create service
	service := goa.New("laravel-jp-conference-api")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	app.MountTimetableController(service, action.NewTimetableController(service))
	app.MountInfoController(service, action.NewInfoController(service))
	app.MountSponsorsController(service, action.NewSponsorsController(service))
	ao := retrievePortOption()

	// Start service
	if err := service.ListenAndServe(fmt.Sprintf(":%s", *ao.Port)); err != nil {
		service.LogError("startup", "err", err)
	}
}

func retrievePortOption() *AppConfig {
	var (
		port string
	)
	flag.StringVar(&port, "port", foundation.ApplicationPort, "port number option")
	flag.Parse()
	return &AppConfig{&port}
}
