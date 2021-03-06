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

	"github.com/fogatlas/client-go/models"
)

// NewPutExternalendpointsIDParams creates a new PutExternalendpointsIDParams object
// with the default values initialized.
func NewPutExternalendpointsIDParams() *PutExternalendpointsIDParams {
	var ()
	return &PutExternalendpointsIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutExternalendpointsIDParamsWithTimeout creates a new PutExternalendpointsIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutExternalendpointsIDParamsWithTimeout(timeout time.Duration) *PutExternalendpointsIDParams {
	var ()
	return &PutExternalendpointsIDParams{

		timeout: timeout,
	}
}

// NewPutExternalendpointsIDParamsWithContext creates a new PutExternalendpointsIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutExternalendpointsIDParamsWithContext(ctx context.Context) *PutExternalendpointsIDParams {
	var ()
	return &PutExternalendpointsIDParams{

		Context: ctx,
	}
}

// NewPutExternalendpointsIDParamsWithHTTPClient creates a new PutExternalendpointsIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutExternalendpointsIDParamsWithHTTPClient(client *http.Client) *PutExternalendpointsIDParams {
	var ()
	return &PutExternalendpointsIDParams{
		HTTPClient: client,
	}
}

/*PutExternalendpointsIDParams contains all the parameters to send to the API endpoint
for the put externalendpoints ID operation typically these are written to a http.Request
*/
type PutExternalendpointsIDParams struct {

	/*Externalendpoint*/
	Externalendpoint *models.ExternalEndpoint
	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put externalendpoints ID params
func (o *PutExternalendpointsIDParams) WithTimeout(timeout time.Duration) *PutExternalendpointsIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put externalendpoints ID params
func (o *PutExternalendpointsIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put externalendpoints ID params
func (o *PutExternalendpointsIDParams) WithContext(ctx context.Context) *PutExternalendpointsIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put externalendpoints ID params
func (o *PutExternalendpointsIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put externalendpoints ID params
func (o *PutExternalendpointsIDParams) WithHTTPClient(client *http.Client) *PutExternalendpointsIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put externalendpoints ID params
func (o *PutExternalendpointsIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExternalendpoint adds the externalendpoint to the put externalendpoints ID params
func (o *PutExternalendpointsIDParams) WithExternalendpoint(externalendpoint *models.ExternalEndpoint) *PutExternalendpointsIDParams {
	o.SetExternalendpoint(externalendpoint)
	return o
}

// SetExternalendpoint adds the externalendpoint to the put externalendpoints ID params
func (o *PutExternalendpointsIDParams) SetExternalendpoint(externalendpoint *models.ExternalEndpoint) {
	o.Externalendpoint = externalendpoint
}

// WithID adds the id to the put externalendpoints ID params
func (o *PutExternalendpointsIDParams) WithID(id string) *PutExternalendpointsIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the put externalendpoints ID params
func (o *PutExternalendpointsIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PutExternalendpointsIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Externalendpoint != nil {
		if err := r.SetBodyParam(o.Externalendpoint); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
