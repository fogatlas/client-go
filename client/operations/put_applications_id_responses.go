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

// PutApplicationsIDReader is a Reader for the PutApplicationsID structure.
type PutApplicationsIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutApplicationsIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPutApplicationsIDCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPutApplicationsIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPutApplicationsIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutApplicationsIDCreated creates a PutApplicationsIDCreated with default headers values
func NewPutApplicationsIDCreated() *PutApplicationsIDCreated {
	return &PutApplicationsIDCreated{}
}

/*PutApplicationsIDCreated handles this case with default header values.

Created
*/
type PutApplicationsIDCreated struct {
}

func (o *PutApplicationsIDCreated) Error() string {
	return fmt.Sprintf("[PUT /applications/{id}][%d] putApplicationsIdCreated ", 201)
}

func (o *PutApplicationsIDCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutApplicationsIDBadRequest creates a PutApplicationsIDBadRequest with default headers values
func NewPutApplicationsIDBadRequest() *PutApplicationsIDBadRequest {
	return &PutApplicationsIDBadRequest{}
}

/*PutApplicationsIDBadRequest handles this case with default header values.

Invalid request
*/
type PutApplicationsIDBadRequest struct {
	Payload *models.Message
}

func (o *PutApplicationsIDBadRequest) Error() string {
	return fmt.Sprintf("[PUT /applications/{id}][%d] putApplicationsIdBadRequest  %+v", 400, o.Payload)
}

func (o *PutApplicationsIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutApplicationsIDInternalServerError creates a PutApplicationsIDInternalServerError with default headers values
func NewPutApplicationsIDInternalServerError() *PutApplicationsIDInternalServerError {
	return &PutApplicationsIDInternalServerError{}
}

/*PutApplicationsIDInternalServerError handles this case with default header values.

Generic error
*/
type PutApplicationsIDInternalServerError struct {
	Payload *models.Message
}

func (o *PutApplicationsIDInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /applications/{id}][%d] putApplicationsIdInternalServerError  %+v", 500, o.Payload)
}

func (o *PutApplicationsIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
