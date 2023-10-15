// Code generated by go-swagger; DO NOT EDIT.

package customer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"eda-in-golang/customers/customersclient/models"
)

// RegisterCustomerReader is a Reader for the RegisterCustomer structure.
type RegisterCustomerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RegisterCustomerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRegisterCustomerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewRegisterCustomerDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRegisterCustomerOK creates a RegisterCustomerOK with default headers values
func NewRegisterCustomerOK() *RegisterCustomerOK {
	return &RegisterCustomerOK{}
}

/* RegisterCustomerOK describes a response with status code 200, with default header values.

A successful response.
*/
type RegisterCustomerOK struct {
	Payload *models.CustomerspbRegisterCustomerResponse
}

// IsSuccess returns true when this register customer o k response has a 2xx status code
func (o *RegisterCustomerOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this register customer o k response has a 3xx status code
func (o *RegisterCustomerOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this register customer o k response has a 4xx status code
func (o *RegisterCustomerOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this register customer o k response has a 5xx status code
func (o *RegisterCustomerOK) IsServerError() bool {
	return false
}

// IsCode returns true when this register customer o k response a status code equal to that given
func (o *RegisterCustomerOK) IsCode(code int) bool {
	return code == 200
}

func (o *RegisterCustomerOK) Error() string {
	return fmt.Sprintf("[POST /api/customers][%d] registerCustomerOK  %+v", 200, o.Payload)
}

func (o *RegisterCustomerOK) String() string {
	return fmt.Sprintf("[POST /api/customers][%d] registerCustomerOK  %+v", 200, o.Payload)
}

func (o *RegisterCustomerOK) GetPayload() *models.CustomerspbRegisterCustomerResponse {
	return o.Payload
}

func (o *RegisterCustomerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CustomerspbRegisterCustomerResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRegisterCustomerDefault creates a RegisterCustomerDefault with default headers values
func NewRegisterCustomerDefault(code int) *RegisterCustomerDefault {
	return &RegisterCustomerDefault{
		_statusCode: code,
	}
}

/* RegisterCustomerDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type RegisterCustomerDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// Code gets the status code for the register customer default response
func (o *RegisterCustomerDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this register customer default response has a 2xx status code
func (o *RegisterCustomerDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this register customer default response has a 3xx status code
func (o *RegisterCustomerDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this register customer default response has a 4xx status code
func (o *RegisterCustomerDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this register customer default response has a 5xx status code
func (o *RegisterCustomerDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this register customer default response a status code equal to that given
func (o *RegisterCustomerDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *RegisterCustomerDefault) Error() string {
	return fmt.Sprintf("[POST /api/customers][%d] registerCustomer default  %+v", o._statusCode, o.Payload)
}

func (o *RegisterCustomerDefault) String() string {
	return fmt.Sprintf("[POST /api/customers][%d] registerCustomer default  %+v", o._statusCode, o.Payload)
}

func (o *RegisterCustomerDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *RegisterCustomerDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
