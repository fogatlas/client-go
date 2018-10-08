// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetThingsIDParams creates a new GetThingsIDParams object
// with the default values initialized.
func NewGetThingsIDParams() *GetThingsIDParams {
	var ()
	return &GetThingsIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetThingsIDParamsWithTimeout creates a new GetThingsIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetThingsIDParamsWithTimeout(timeout time.Duration) *GetThingsIDParams {
	var ()
	return &GetThingsIDParams{

		timeout: timeout,
	}
}

// NewGetThingsIDParamsWithContext creates a new GetThingsIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetThingsIDParamsWithContext(ctx context.Context) *GetThingsIDParams {
	var ()
	return &GetThingsIDParams{

		Context: ctx,
	}
}

// NewGetThingsIDParamsWithHTTPClient creates a new GetThingsIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetThingsIDParamsWithHTTPClient(client *http.Client) *GetThingsIDParams {
	var ()
	return &GetThingsIDParams{
		HTTPClient: client,
	}
}

/*GetThingsIDParams contains all the parameters to send to the API endpoint
for the get things ID operation typically these are written to a http.Request
*/
type GetThingsIDParams struct {

	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get things ID params
func (o *GetThingsIDParams) WithTimeout(timeout time.Duration) *GetThingsIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get things ID params
func (o *GetThingsIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get things ID params
func (o *GetThingsIDParams) WithContext(ctx context.Context) *GetThingsIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get things ID params
func (o *GetThingsIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get things ID params
func (o *GetThingsIDParams) WithHTTPClient(client *http.Client) *GetThingsIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get things ID params
func (o *GetThingsIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get things ID params
func (o *GetThingsIDParams) WithID(id string) *GetThingsIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get things ID params
func (o *GetThingsIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetThingsIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}