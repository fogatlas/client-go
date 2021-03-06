// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/fogatlas/client-go/models"
)

// GetRegionsReader is a Reader for the GetRegions structure.
type GetRegionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRegionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetRegionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetRegionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetRegionsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetRegionsOK creates a GetRegionsOK with default headers values
func NewGetRegionsOK() *GetRegionsOK {
	return &GetRegionsOK{}
}

/*GetRegionsOK handles this case with default header values.

Successfully returned the list of regions
*/
type GetRegionsOK struct {
	Payload *models.GetRegionsOKBody
}

func (o *GetRegionsOK) Error() string {
	return fmt.Sprintf("[GET /regions][%d] getRegionsOK  %+v", 200, o.Payload)
}

func (o *GetRegionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetRegionsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRegionsBadRequest creates a GetRegionsBadRequest with default headers values
func NewGetRegionsBadRequest() *GetRegionsBadRequest {
	return &GetRegionsBadRequest{}
}

/*GetRegionsBadRequest handles this case with default header values.

Invalid request
*/
type GetRegionsBadRequest struct {
	Payload *models.Message
}

func (o *GetRegionsBadRequest) Error() string {
	return fmt.Sprintf("[GET /regions][%d] getRegionsBadRequest  %+v", 400, o.Payload)
}

func (o *GetRegionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRegionsInternalServerError creates a GetRegionsInternalServerError with default headers values
func NewGetRegionsInternalServerError() *GetRegionsInternalServerError {
	return &GetRegionsInternalServerError{}
}

/*GetRegionsInternalServerError handles this case with default header values.

Generic error
*/
type GetRegionsInternalServerError struct {
	Payload *models.Message
}

func (o *GetRegionsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /regions][%d] getRegionsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRegionsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
