// Code generated by go-swagger; DO NOT EDIT.

package alerting

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

// NewDeleteTemplateParams creates a new DeleteTemplateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteTemplateParams() *DeleteTemplateParams {
	return &DeleteTemplateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteTemplateParamsWithTimeout creates a new DeleteTemplateParams object
// with the ability to set a timeout on a request.
func NewDeleteTemplateParamsWithTimeout(timeout time.Duration) *DeleteTemplateParams {
	return &DeleteTemplateParams{
		timeout: timeout,
	}
}

// NewDeleteTemplateParamsWithContext creates a new DeleteTemplateParams object
// with the ability to set a context for a request.
func NewDeleteTemplateParamsWithContext(ctx context.Context) *DeleteTemplateParams {
	return &DeleteTemplateParams{
		Context: ctx,
	}
}

// NewDeleteTemplateParamsWithHTTPClient creates a new DeleteTemplateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteTemplateParamsWithHTTPClient(client *http.Client) *DeleteTemplateParams {
	return &DeleteTemplateParams{
		HTTPClient: client,
	}
}

/*
DeleteTemplateParams contains all the parameters to send to the API endpoint

	for the delete template operation.

	Typically these are written to a http.Request.
*/
type DeleteTemplateParams struct {
	// Body.
	Body DeleteTemplateBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete template params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteTemplateParams) WithDefaults() *DeleteTemplateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete template params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteTemplateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete template params
func (o *DeleteTemplateParams) WithTimeout(timeout time.Duration) *DeleteTemplateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete template params
func (o *DeleteTemplateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete template params
func (o *DeleteTemplateParams) WithContext(ctx context.Context) *DeleteTemplateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete template params
func (o *DeleteTemplateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete template params
func (o *DeleteTemplateParams) WithHTTPClient(client *http.Client) *DeleteTemplateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete template params
func (o *DeleteTemplateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the delete template params
func (o *DeleteTemplateParams) WithBody(body DeleteTemplateBody) *DeleteTemplateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the delete template params
func (o *DeleteTemplateParams) SetBody(body DeleteTemplateBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteTemplateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {
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
