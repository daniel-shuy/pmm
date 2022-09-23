// Code generated by go-swagger; DO NOT EDIT.

package external

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewAddExternalParams creates a new AddExternalParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddExternalParams() *AddExternalParams {
	return &AddExternalParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddExternalParamsWithTimeout creates a new AddExternalParams object
// with the ability to set a timeout on a request.
func NewAddExternalParamsWithTimeout(timeout time.Duration) *AddExternalParams {
	return &AddExternalParams{
		timeout: timeout,
	}
}

// NewAddExternalParamsWithContext creates a new AddExternalParams object
// with the ability to set a context for a request.
func NewAddExternalParamsWithContext(ctx context.Context) *AddExternalParams {
	return &AddExternalParams{
		Context: ctx,
	}
}

// NewAddExternalParamsWithHTTPClient creates a new AddExternalParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddExternalParamsWithHTTPClient(client *http.Client) *AddExternalParams {
	return &AddExternalParams{
		HTTPClient: client,
	}
}

/*
AddExternalParams contains all the parameters to send to the API endpoint

	for the add external operation.

	Typically these are written to a http.Request.
*/
type AddExternalParams struct {
	// Body.
	Body AddExternalBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add external params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddExternalParams) WithDefaults() *AddExternalParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add external params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddExternalParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the add external params
func (o *AddExternalParams) WithTimeout(timeout time.Duration) *AddExternalParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add external params
func (o *AddExternalParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add external params
func (o *AddExternalParams) WithContext(ctx context.Context) *AddExternalParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add external params
func (o *AddExternalParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add external params
func (o *AddExternalParams) WithHTTPClient(client *http.Client) *AddExternalParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add external params
func (o *AddExternalParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add external params
func (o *AddExternalParams) WithBody(body AddExternalBody) *AddExternalParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add external params
func (o *AddExternalParams) SetBody(body AddExternalBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AddExternalParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {
	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
