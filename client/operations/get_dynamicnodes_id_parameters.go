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

// NewGetDynamicnodesIDParams creates a new GetDynamicnodesIDParams object
// with the default values initialized.
func NewGetDynamicnodesIDParams() *GetDynamicnodesIDParams {
	var ()
	return &GetDynamicnodesIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDynamicnodesIDParamsWithTimeout creates a new GetDynamicnodesIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDynamicnodesIDParamsWithTimeout(timeout time.Duration) *GetDynamicnodesIDParams {
	var ()
	return &GetDynamicnodesIDParams{

		timeout: timeout,
	}
}

// NewGetDynamicnodesIDParamsWithContext creates a new GetDynamicnodesIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDynamicnodesIDParamsWithContext(ctx context.Context) *GetDynamicnodesIDParams {
	var ()
	return &GetDynamicnodesIDParams{

		Context: ctx,
	}
}

// NewGetDynamicnodesIDParamsWithHTTPClient creates a new GetDynamicnodesIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDynamicnodesIDParamsWithHTTPClient(client *http.Client) *GetDynamicnodesIDParams {
	var ()
	return &GetDynamicnodesIDParams{
		HTTPClient: client,
	}
}

/*GetDynamicnodesIDParams contains all the parameters to send to the API endpoint
for the get dynamicnodes ID operation typically these are written to a http.Request
*/
type GetDynamicnodesIDParams struct {

	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get dynamicnodes ID params
func (o *GetDynamicnodesIDParams) WithTimeout(timeout time.Duration) *GetDynamicnodesIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get dynamicnodes ID params
func (o *GetDynamicnodesIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get dynamicnodes ID params
func (o *GetDynamicnodesIDParams) WithContext(ctx context.Context) *GetDynamicnodesIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get dynamicnodes ID params
func (o *GetDynamicnodesIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get dynamicnodes ID params
func (o *GetDynamicnodesIDParams) WithHTTPClient(client *http.Client) *GetDynamicnodesIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get dynamicnodes ID params
func (o *GetDynamicnodesIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get dynamicnodes ID params
func (o *GetDynamicnodesIDParams) WithID(id string) *GetDynamicnodesIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get dynamicnodes ID params
func (o *GetDynamicnodesIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetDynamicnodesIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
