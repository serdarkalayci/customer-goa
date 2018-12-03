package main

import (
	"github.com/goadesign/goa"
	"github.com/gostore/src/customer/app"
)

// CustomerController implements the customer resource.
type CustomerController struct {
	*goa.Controller
}

// NewCustomerController creates a customer controller.
func NewCustomerController(service *goa.Service) *CustomerController {
	return &CustomerController{Controller: service.NewController("CustomerController")}
}

// List runs the list action.
func (c *CustomerController) List(ctx *app.ListCustomerContext) error {
	// CustomerController_List: start_implement

	// Put your logic here

	res := &app.GoaExampleCustomer{}
	return ctx.OK(res)
	// CustomerController_List: end_implement
}

// Show runs the show action.
func (c *CustomerController) Show(ctx *app.ShowCustomerContext) error {
	// CustomerController_Show: start_implement

	// Put your logic here

	res := app.GoaExampleCustomer{
		ID:   1,
		Name: "Mahmut",
	}
	return ctx.OK(&res)
	// CustomerController_Show: end_implement
}
