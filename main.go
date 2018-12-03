//go:generate goagen bootstrap -d github.com/gostore/src/customer/design

package main

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/gostore/src/customer/app"
)

func main() {
	// Create service
	service := goa.New("customer")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "customer" controller
	c := NewCustomerController(service)
	app.MountCustomerController(service, c)

	d := NewSwaggerController(service)
	app.MountSwaggerController(service, d)

	// Start service
	if err := service.ListenAndServe(":6543"); err != nil {
		service.LogError("startup", "err", err)
	}

}
