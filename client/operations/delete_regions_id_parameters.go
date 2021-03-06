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

// NewDeleteRegionsIDParams creates a new DeleteRegionsIDParams object
// with the default values initialized.
func NewDeleteRegionsIDParams() *DeleteRegionsIDParams {
	var ()
	return &DeleteRegionsIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteRegionsIDParamsWithTimeout creates a new DeleteRegionsIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteRegionsIDParamsWithTimeout(timeout time.Duration) *DeleteRegionsIDParams {
	var ()
	return &DeleteRegionsIDParams{

		timeout: timeout,
	}
}

// NewDeleteRegionsIDParamsWithContext creates a new DeleteRegionsIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteRegionsIDParamsWithContext(ctx context.Context) *DeleteRegionsIDParams {
	var ()
	return &DeleteRegionsIDParams{

		Context: ctx,
	}
}

// NewDeleteRegionsIDParamsWithHTTPClient creates a new DeleteRegionsIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteRegionsIDParamsWithHTTPClient(client *http.Client) *DeleteRegionsIDParams {
	var ()
	return &DeleteRegionsIDParams{
		HTTPClient: client,
	}
}

/*DeleteRegionsIDParams contains all the parameters to send to the API endpoint
for the delete regions ID operation typically these are written to a http.Request
*/
type DeleteRegionsIDParams struct {

	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete regions ID params
func (o *DeleteRegionsIDParams) WithTimeout(timeout time.Duration) *DeleteRegionsIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete regions ID params
func (o *DeleteRegionsIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete regions ID params
func (o *DeleteRegionsIDParams) WithContext(ctx context.Context) *DeleteRegionsIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete regions ID params
func (o *DeleteRegionsIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete regions ID params
func (o *DeleteRegionsIDParams) WithHTTPClient(client *http.Client) *DeleteRegionsIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete regions ID params
func (o *DeleteRegionsIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete regions ID params
func (o *DeleteRegionsIDParams) WithID(id string) *DeleteRegionsIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete regions ID params
func (o *DeleteRegionsIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteRegionsIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
