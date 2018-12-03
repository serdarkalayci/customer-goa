// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "customer": Application Media Types
//
// Command:
// $ goagen
// --design=github.com/gostore/src/customer/design
// --out=$(GOPATH)/src/github.com/gostore/src/customer
// --version=v1.3.1

package client

import (
	"github.com/goadesign/goa"
	"net/http"
)

// A customer of Westeros Shop (default view)
//
// Identifier: application/vnd.goa.example.customer+json; view=default
type GoaExampleCustomer struct {
	// Unique customer ID
	ID int `form:"id" json:"id" yaml:"id" xml:"id"`
	// Name of wine
	Name string `form:"name" json:"name" yaml:"name" xml:"name"`
}

// Validate validates the GoaExampleCustomer media type instance.
func (mt *GoaExampleCustomer) Validate() (err error) {

	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	return
}

// DecodeGoaExampleCustomer decodes the GoaExampleCustomer instance encoded in resp body.
func (c *Client) DecodeGoaExampleCustomer(resp *http.Response) (*GoaExampleCustomer, error) {
	var decoded GoaExampleCustomer
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}
