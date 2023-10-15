// Code generated by go-swagger; DO NOT EDIT.

package participation

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"eda-in-golang/stores/storesclient/models"
)

// GetParticipatingStoresReader is a Reader for the GetParticipatingStores structure.
type GetParticipatingStoresReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetParticipatingStoresReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetParticipatingStoresOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetParticipatingStoresDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetParticipatingStoresOK creates a GetParticipatingStoresOK with default headers values
func NewGetParticipatingStoresOK() *GetParticipatingStoresOK {
	return &GetParticipatingStoresOK{}
}

/* GetParticipatingStoresOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetParticipatingStoresOK struct {
	Payload *models.StorespbGetParticipatingStoresResponse
}

func (o *GetParticipatingStoresOK) Error() string {
	return fmt.Sprintf("[GET /api/stores/participating][%d] getParticipatingStoresOK  %+v", 200, o.Payload)
}
func (o *GetParticipatingStoresOK) GetPayload() *models.StorespbGetParticipatingStoresResponse {
	return o.Payload
}

func (o *GetParticipatingStoresOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StorespbGetParticipatingStoresResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetParticipatingStoresDefault creates a GetParticipatingStoresDefault with default headers values
func NewGetParticipatingStoresDefault(code int) *GetParticipatingStoresDefault {
	return &GetParticipatingStoresDefault{
		_statusCode: code,
	}
}

/* GetParticipatingStoresDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type GetParticipatingStoresDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// Code gets the status code for the get participating stores default response
func (o *GetParticipatingStoresDefault) Code() int {
	return o._statusCode
}

func (o *GetParticipatingStoresDefault) Error() string {
	return fmt.Sprintf("[GET /api/stores/participating][%d] getParticipatingStores default  %+v", o._statusCode, o.Payload)
}
func (o *GetParticipatingStoresDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *GetParticipatingStoresDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
