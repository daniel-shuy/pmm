// Code generated by go-swagger; DO NOT EDIT.

package channels

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

// NewChangeChannelParams creates a new ChangeChannelParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewChangeChannelParams() *ChangeChannelParams {
	return &ChangeChannelParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewChangeChannelParamsWithTimeout creates a new ChangeChannelParams object
// with the ability to set a timeout on a request.
func NewChangeChannelParamsWithTimeout(timeout time.Duration) *ChangeChannelParams {
	return &ChangeChannelParams{
		timeout: timeout,
	}
}

// NewChangeChannelParamsWithContext creates a new ChangeChannelParams object
// with the ability to set a context for a request.
func NewChangeChannelParamsWithContext(ctx context.Context) *ChangeChannelParams {
	return &ChangeChannelParams{
		Context: ctx,
	}
}

// NewChangeChannelParamsWithHTTPClient creates a new ChangeChannelParams object
// with the ability to set a custom HTTPClient for a request.
func NewChangeChannelParamsWithHTTPClient(client *http.Client) *ChangeChannelParams {
	return &ChangeChannelParams{
		HTTPClient: client,
	}
}

/*
ChangeChannelParams contains all the parameters to send to the API endpoint

	for the change channel operation.

	Typically these are written to a http.Request.
*/
type ChangeChannelParams struct {
	// Body.
	Body ChangeChannelBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the change channel params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ChangeChannelParams) WithDefaults() *ChangeChannelParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the change channel params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ChangeChannelParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the change channel params
func (o *ChangeChannelParams) WithTimeout(timeout time.Duration) *ChangeChannelParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the change channel params
func (o *ChangeChannelParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the change channel params
func (o *ChangeChannelParams) WithContext(ctx context.Context) *ChangeChannelParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the change channel params
func (o *ChangeChannelParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the change channel params
func (o *ChangeChannelParams) WithHTTPClient(client *http.Client) *ChangeChannelParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the change channel params
func (o *ChangeChannelParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the change channel params
func (o *ChangeChannelParams) WithBody(body ChangeChannelBody) *ChangeChannelParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the change channel params
func (o *ChangeChannelParams) SetBody(body ChangeChannelBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ChangeChannelParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {
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
