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

// NewPutThingsIDParams creates a new PutThingsIDParams object
// with the default values initialized.
func NewPutThingsIDParams() *PutThingsIDParams {
	var ()
	return &PutThingsIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutThingsIDParamsWithTimeout creates a new PutThingsIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutThingsIDParamsWithTimeout(timeout time.Duration) *PutThingsIDParams {
	var ()
	return &PutThingsIDParams{

		timeout: timeout,
	}
}

// NewPutThingsIDParamsWithContext creates a new PutThingsIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutThingsIDParamsWithContext(ctx context.Context) *PutThingsIDParams {
	var ()
	return &PutThingsIDParams{

		Context: ctx,
	}
}

// NewPutThingsIDParamsWithHTTPClient creates a new PutThingsIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutThingsIDParamsWithHTTPClient(client *http.Client) *PutThingsIDParams {
	var ()
	return &PutThingsIDParams{
		HTTPClient: client,
	}
}

/*PutThingsIDParams contains all the parameters to send to the API endpoint
for the put things ID operation typically these are written to a http.Request
*/
type PutThingsIDParams struct {

	/*ID*/
	ID string
	/*Thing*/
	Thing *models.Thing

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put things ID params
func (o *PutThingsIDParams) WithTimeout(timeout time.Duration) *PutThingsIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put things ID params
func (o *PutThingsIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put things ID params
func (o *PutThingsIDParams) WithContext(ctx context.Context) *PutThingsIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put things ID params
func (o *PutThingsIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put things ID params
func (o *PutThingsIDParams) WithHTTPClient(client *http.Client) *PutThingsIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put things ID params
func (o *PutThingsIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the put things ID params
func (o *PutThingsIDParams) WithID(id string) *PutThingsIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the put things ID params
func (o *PutThingsIDParams) SetID(id string) {
	o.ID = id
}

// WithThing adds the thing to the put things ID params
func (o *PutThingsIDParams) WithThing(thing *models.Thing) *PutThingsIDParams {
	o.SetThing(thing)
	return o
}

// SetThing adds the thing to the put things ID params
func (o *PutThingsIDParams) SetThing(thing *models.Thing) {
	o.Thing = thing
}

// WriteToRequest writes these params to a swagger request
func (o *PutThingsIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if o.Thing != nil {
		if err := r.SetBodyParam(o.Thing); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
