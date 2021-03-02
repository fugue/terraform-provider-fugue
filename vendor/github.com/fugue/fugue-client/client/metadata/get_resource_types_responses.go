// Code generated by go-swagger; DO NOT EDIT.

package metadata

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/fugue/fugue-client/models"
)

// GetResourceTypesReader is a Reader for the GetResourceTypes structure.
type GetResourceTypesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetResourceTypesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetResourceTypesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetResourceTypesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetResourceTypesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetResourceTypesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetResourceTypesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetResourceTypesOK creates a GetResourceTypesOK with default headers values
func NewGetResourceTypesOK() *GetResourceTypesOK {
	return &GetResourceTypesOK{}
}

/*GetResourceTypesOK handles this case with default header values.

List of supported resource types.
*/
type GetResourceTypesOK struct {
	Payload *models.ResourceTypeMetadata
}

func (o *GetResourceTypesOK) Error() string {
	return fmt.Sprintf("[GET /metadata/{provider}/resource_types][%d] getResourceTypesOK  %+v", 200, o.Payload)
}

func (o *GetResourceTypesOK) GetPayload() *models.ResourceTypeMetadata {
	return o.Payload
}

func (o *GetResourceTypesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResourceTypeMetadata)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResourceTypesBadRequest creates a GetResourceTypesBadRequest with default headers values
func NewGetResourceTypesBadRequest() *GetResourceTypesBadRequest {
	return &GetResourceTypesBadRequest{}
}

/*GetResourceTypesBadRequest handles this case with default header values.

BadRequestError
*/
type GetResourceTypesBadRequest struct {
	Payload *models.BadRequestError
}

func (o *GetResourceTypesBadRequest) Error() string {
	return fmt.Sprintf("[GET /metadata/{provider}/resource_types][%d] getResourceTypesBadRequest  %+v", 400, o.Payload)
}

func (o *GetResourceTypesBadRequest) GetPayload() *models.BadRequestError {
	return o.Payload
}

func (o *GetResourceTypesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResourceTypesUnauthorized creates a GetResourceTypesUnauthorized with default headers values
func NewGetResourceTypesUnauthorized() *GetResourceTypesUnauthorized {
	return &GetResourceTypesUnauthorized{}
}

/*GetResourceTypesUnauthorized handles this case with default header values.

AuthenticationError
*/
type GetResourceTypesUnauthorized struct {
	Payload *models.AuthenticationError
}

func (o *GetResourceTypesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /metadata/{provider}/resource_types][%d] getResourceTypesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetResourceTypesUnauthorized) GetPayload() *models.AuthenticationError {
	return o.Payload
}

func (o *GetResourceTypesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AuthenticationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResourceTypesForbidden creates a GetResourceTypesForbidden with default headers values
func NewGetResourceTypesForbidden() *GetResourceTypesForbidden {
	return &GetResourceTypesForbidden{}
}

/*GetResourceTypesForbidden handles this case with default header values.

AuthorizationError
*/
type GetResourceTypesForbidden struct {
	Payload *models.AuthorizationError
}

func (o *GetResourceTypesForbidden) Error() string {
	return fmt.Sprintf("[GET /metadata/{provider}/resource_types][%d] getResourceTypesForbidden  %+v", 403, o.Payload)
}

func (o *GetResourceTypesForbidden) GetPayload() *models.AuthorizationError {
	return o.Payload
}

func (o *GetResourceTypesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AuthorizationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResourceTypesInternalServerError creates a GetResourceTypesInternalServerError with default headers values
func NewGetResourceTypesInternalServerError() *GetResourceTypesInternalServerError {
	return &GetResourceTypesInternalServerError{}
}

/*GetResourceTypesInternalServerError handles this case with default header values.

InternalServerError
*/
type GetResourceTypesInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetResourceTypesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /metadata/{provider}/resource_types][%d] getResourceTypesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetResourceTypesInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *GetResourceTypesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
