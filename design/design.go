package design // The convention consists of naming the design
// package "design"
import (
	. "github.com/goadesign/goa/design" // Use . imports to enable the DSL
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("customer", func() { // API defines the microservice endpoint and
	Title("The customer service of Westeros Shop")       // other global properties. There should be one
	Description("The customer service of Westeros Shop") // and exactly one API definition appearing in
	Scheme("http")                                       // the design.
	Host("localhost:6543")
})

var _ = Resource("customer", func() { // Resources group related API endpoints
	BasePath("/customers")      // together. They map to REST resources for REST
	DefaultMedia(CustomerMedia) // services.

	Action("list", func() { // Actions define a single API endpoint together
		Description("Lists all customers") // with its path, parameters (both path
		Routing(GET("/"))                  // parameters and querystring values) and payload
		Response(OK)                       // Responses define the shape and status code
		Response(NotFound)                 // of HTTP responses.
	})

	Action("show", func() { // Actions define a single API endpoint together
		Description("Get customer by id") // with its path, parameters (both path
		Routing(GET("/:customerID"))      // parameters and querystring values) and payload
		Params(func() {                   // (shape of the request body).
			Param("customerID", Integer, "customer ID")
		})
		Response(OK)       // Responses define the shape and status code
		Response(NotFound) // of HTTP responses.
	})
})

var _ = Resource("swagger", func() {
	Origin("*", func() {
		Methods("GET", "OPTIONS")
	})
	Files("/swagger/*filepath", "public/swagger/")
})

// customerMedia defines the media type used to render customers.
var CustomerMedia = MediaType("application/vnd.goa.example.customer+json", func() {
	Description("A customer of Westeros Shop")
	Attributes(func() { // Attributes define the media type shape.
		Attribute("id", Integer, "Unique customer ID")
		Attribute("name", String, "Name of wine")
		Required("id", "name")
	})
	View("default", func() { // View defines a rendering of the media type.
		Attribute("id") // Media types may have multiple views and must
		Attribute("name")
	})
})
