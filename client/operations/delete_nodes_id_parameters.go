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

// NewDeleteNodesIDParams creates a new DeleteNodesIDParams object
// with the default values initialized.
func NewDeleteNodesIDParams() *DeleteNodesIDParams {
	var ()
	return &DeleteNodesIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteNodesIDParamsWithTimeout creates a new DeleteNodesIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteNodesIDParamsWithTimeout(timeout time.Duration) *DeleteNodesIDParams {
	var ()
	return &DeleteNodesIDParams{

		timeout: timeout,
	}
}

// NewDeleteNodesIDParamsWithContext creates a new DeleteNodesIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteNodesIDParamsWithContext(ctx context.Context) *DeleteNodesIDParams {
	var ()
	return &DeleteNodesIDParams{

		Context: ctx,
	}
}

// NewDeleteNodesIDParamsWithHTTPClient creates a new DeleteNodesIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteNodesIDParamsWithHTTPClient(client *http.Client) *DeleteNodesIDParams {
	var ()
	return &DeleteNodesIDParams{
		HTTPClient: client,
	}
}

/*DeleteNodesIDParams contains all the parameters to send to the API endpoint
for the delete nodes ID operation typically these are written to a http.Request
*/
type DeleteNodesIDParams struct {

	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete nodes ID params
func (o *DeleteNodesIDParams) WithTimeout(timeout time.Duration) *DeleteNodesIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete nodes ID params
func (o *DeleteNodesIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete nodes ID params
func (o *DeleteNodesIDParams) WithContext(ctx context.Context) *DeleteNodesIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete nodes ID params
func (o *DeleteNodesIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete nodes ID params
func (o *DeleteNodesIDParams) WithHTTPClient(client *http.Client) *DeleteNodesIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete nodes ID params
func (o *DeleteNodesIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete nodes ID params
func (o *DeleteNodesIDParams) WithID(id string) *DeleteNodesIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete nodes ID params
func (o *DeleteNodesIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteNodesIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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