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

// DeleteDynamicnodesIDReader is a Reader for the DeleteDynamicnodesID structure.
type DeleteDynamicnodesIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteDynamicnodesIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteDynamicnodesIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteDynamicnodesIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewDeleteDynamicnodesIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteDynamicnodesIDNoContent creates a DeleteDynamicnodesIDNoContent with default headers values
func NewDeleteDynamicnodesIDNoContent() *DeleteDynamicnodesIDNoContent {
	return &DeleteDynamicnodesIDNoContent{}
}

/*DeleteDynamicnodesIDNoContent handles this case with default header values.

Dynamic node deleted
*/
type DeleteDynamicnodesIDNoContent struct {
}

func (o *DeleteDynamicnodesIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dynamicnodes/{id}][%d] deleteDynamicnodesIdNoContent ", 204)
}

func (o *DeleteDynamicnodesIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteDynamicnodesIDBadRequest creates a DeleteDynamicnodesIDBadRequest with default headers values
func NewDeleteDynamicnodesIDBadRequest() *DeleteDynamicnodesIDBadRequest {
	return &DeleteDynamicnodesIDBadRequest{}
}

/*DeleteDynamicnodesIDBadRequest handles this case with default header values.

Invalid request
*/
type DeleteDynamicnodesIDBadRequest struct {
	Payload *models.Message
}

func (o *DeleteDynamicnodesIDBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /dynamicnodes/{id}][%d] deleteDynamicnodesIdBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteDynamicnodesIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDynamicnodesIDInternalServerError creates a DeleteDynamicnodesIDInternalServerError with default headers values
func NewDeleteDynamicnodesIDInternalServerError() *DeleteDynamicnodesIDInternalServerError {
	return &DeleteDynamicnodesIDInternalServerError{}
}

/*DeleteDynamicnodesIDInternalServerError handles this case with default header values.

Generic error
*/
type DeleteDynamicnodesIDInternalServerError struct {
	Payload *models.Message
}

func (o *DeleteDynamicnodesIDInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /dynamicnodes/{id}][%d] deleteDynamicnodesIdInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteDynamicnodesIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
