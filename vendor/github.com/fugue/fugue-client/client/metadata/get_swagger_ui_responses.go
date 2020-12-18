// Code generated by go-swagger; DO NOT EDIT.

package metadata

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetSwaggerUIReader is a Reader for the GetSwaggerUI structure.
type GetSwaggerUIReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSwaggerUIReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSwaggerUIOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetSwaggerUIOK creates a GetSwaggerUIOK with default headers values
func NewGetSwaggerUIOK() *GetSwaggerUIOK {
	return &GetSwaggerUIOK{}
}

/*GetSwaggerUIOK handles this case with default header values.

The Swagger UI
*/
type GetSwaggerUIOK struct {
	AccessControlAllowHeaders string

	AccessControlAllowMethods string

	AccessControlAllowOrigin string

	ContentType string
}

func (o *GetSwaggerUIOK) Error() string {
	return fmt.Sprintf("[GET /swagger/ui][%d] getSwaggerUiOK ", 200)
}

func (o *GetSwaggerUIOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Access-Control-Allow-Headers
	o.AccessControlAllowHeaders = response.GetHeader("Access-Control-Allow-Headers")

	// response header Access-Control-Allow-Methods
	o.AccessControlAllowMethods = response.GetHeader("Access-Control-Allow-Methods")

	// response header Access-Control-Allow-Origin
	o.AccessControlAllowOrigin = response.GetHeader("Access-Control-Allow-Origin")

	// response header Content-Type
	o.ContentType = response.GetHeader("Content-Type")

	return nil
}