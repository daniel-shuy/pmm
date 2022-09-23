// Code generated by go-swagger; DO NOT EDIT.

package pxc_clusters

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

// NewCreatePXCClusterParams creates a new CreatePXCClusterParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreatePXCClusterParams() *CreatePXCClusterParams {
	return &CreatePXCClusterParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreatePXCClusterParamsWithTimeout creates a new CreatePXCClusterParams object
// with the ability to set a timeout on a request.
func NewCreatePXCClusterParamsWithTimeout(timeout time.Duration) *CreatePXCClusterParams {
	return &CreatePXCClusterParams{
		timeout: timeout,
	}
}

// NewCreatePXCClusterParamsWithContext creates a new CreatePXCClusterParams object
// with the ability to set a context for a request.
func NewCreatePXCClusterParamsWithContext(ctx context.Context) *CreatePXCClusterParams {
	return &CreatePXCClusterParams{
		Context: ctx,
	}
}

// NewCreatePXCClusterParamsWithHTTPClient creates a new CreatePXCClusterParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreatePXCClusterParamsWithHTTPClient(client *http.Client) *CreatePXCClusterParams {
	return &CreatePXCClusterParams{
		HTTPClient: client,
	}
}

/*
CreatePXCClusterParams contains all the parameters to send to the API endpoint

	for the create PXC cluster operation.

	Typically these are written to a http.Request.
*/
type CreatePXCClusterParams struct {
	// Body.
	Body CreatePXCClusterBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create PXC cluster params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreatePXCClusterParams) WithDefaults() *CreatePXCClusterParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create PXC cluster params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreatePXCClusterParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create PXC cluster params
func (o *CreatePXCClusterParams) WithTimeout(timeout time.Duration) *CreatePXCClusterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create PXC cluster params
func (o *CreatePXCClusterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create PXC cluster params
func (o *CreatePXCClusterParams) WithContext(ctx context.Context) *CreatePXCClusterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create PXC cluster params
func (o *CreatePXCClusterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create PXC cluster params
func (o *CreatePXCClusterParams) WithHTTPClient(client *http.Client) *CreatePXCClusterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create PXC cluster params
func (o *CreatePXCClusterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create PXC cluster params
func (o *CreatePXCClusterParams) WithBody(body CreatePXCClusterBody) *CreatePXCClusterParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create PXC cluster params
func (o *CreatePXCClusterParams) SetBody(body CreatePXCClusterBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreatePXCClusterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {
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
