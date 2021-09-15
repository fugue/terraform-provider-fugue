// Code generated by go-swagger; DO NOT EDIT.

package families

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/fugue/fugue-client/models"
)

// ListFamiliesReader is a Reader for the ListFamilies structure.
type ListFamiliesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListFamiliesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListFamiliesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListFamiliesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewListFamiliesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListFamiliesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListFamiliesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListFamiliesOK creates a ListFamiliesOK with default headers values
func NewListFamiliesOK() *ListFamiliesOK {
	return &ListFamiliesOK{}
}

/*ListFamiliesOK handles this case with default header values.

List of compliance families.
*/
type ListFamiliesOK struct {
	Payload *models.Families
}

func (o *ListFamiliesOK) Error() string {
	return fmt.Sprintf("[GET /families][%d] listFamiliesOK  %+v", 200, o.Payload)
}

func (o *ListFamiliesOK) GetPayload() *models.Families {
	return o.Payload
}

func (o *ListFamiliesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Families)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListFamiliesBadRequest creates a ListFamiliesBadRequest with default headers values
func NewListFamiliesBadRequest() *ListFamiliesBadRequest {
	return &ListFamiliesBadRequest{}
}

/*ListFamiliesBadRequest handles this case with default header values.

BadRequestError
*/
type ListFamiliesBadRequest struct {
	Payload *models.BadRequestError
}

func (o *ListFamiliesBadRequest) Error() string {
	return fmt.Sprintf("[GET /families][%d] listFamiliesBadRequest  %+v", 400, o.Payload)
}

func (o *ListFamiliesBadRequest) GetPayload() *models.BadRequestError {
	return o.Payload
}

func (o *ListFamiliesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListFamiliesUnauthorized creates a ListFamiliesUnauthorized with default headers values
func NewListFamiliesUnauthorized() *ListFamiliesUnauthorized {
	return &ListFamiliesUnauthorized{}
}

/*ListFamiliesUnauthorized handles this case with default header values.

AuthenticationError
*/
type ListFamiliesUnauthorized struct {
	Payload *models.AuthenticationError
}

func (o *ListFamiliesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /families][%d] listFamiliesUnauthorized  %+v", 401, o.Payload)
}

func (o *ListFamiliesUnauthorized) GetPayload() *models.AuthenticationError {
	return o.Payload
}

func (o *ListFamiliesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AuthenticationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListFamiliesForbidden creates a ListFamiliesForbidden with default headers values
func NewListFamiliesForbidden() *ListFamiliesForbidden {
	return &ListFamiliesForbidden{}
}

/*ListFamiliesForbidden handles this case with default header values.

AuthorizationError
*/
type ListFamiliesForbidden struct {
	Payload *models.AuthorizationError
}

func (o *ListFamiliesForbidden) Error() string {
	return fmt.Sprintf("[GET /families][%d] listFamiliesForbidden  %+v", 403, o.Payload)
}

func (o *ListFamiliesForbidden) GetPayload() *models.AuthorizationError {
	return o.Payload
}

func (o *ListFamiliesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AuthorizationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListFamiliesInternalServerError creates a ListFamiliesInternalServerError with default headers values
func NewListFamiliesInternalServerError() *ListFamiliesInternalServerError {
	return &ListFamiliesInternalServerError{}
}

/*ListFamiliesInternalServerError handles this case with default header values.

InternalServerError
*/
type ListFamiliesInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *ListFamiliesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /families][%d] listFamiliesInternalServerError  %+v", 500, o.Payload)
}

func (o *ListFamiliesInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *ListFamiliesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
