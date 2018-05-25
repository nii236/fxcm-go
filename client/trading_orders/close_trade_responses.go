// Code generated by go-swagger; DO NOT EDIT.

package trading_orders

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/nii236/fxcm-go/models"
)

// CloseTradeReader is a Reader for the CloseTrade structure.
type CloseTradeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CloseTradeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCloseTradeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewCloseTradeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewCloseTradeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewCloseTradeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCloseTradeOK creates a CloseTradeOK with default headers values
func NewCloseTradeOK() *CloseTradeOK {
	return &CloseTradeOK{}
}

/*CloseTradeOK handles this case with default header values.

OK
*/
type CloseTradeOK struct {
	Payload *models.TraderOrdersResponseDefinition
}

func (o *CloseTradeOK) Error() string {
	return fmt.Sprintf("[POST /trading/close_trade][%d] closeTradeOK  %+v", 200, o.Payload)
}

func (o *CloseTradeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TraderOrdersResponseDefinition)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCloseTradeUnauthorized creates a CloseTradeUnauthorized with default headers values
func NewCloseTradeUnauthorized() *CloseTradeUnauthorized {
	return &CloseTradeUnauthorized{}
}

/*CloseTradeUnauthorized handles this case with default header values.

Unauthorized
*/
type CloseTradeUnauthorized struct {
	Payload *models.Error401Unauthorized
}

func (o *CloseTradeUnauthorized) Error() string {
	return fmt.Sprintf("[POST /trading/close_trade][%d] closeTradeUnauthorized  %+v", 401, o.Payload)
}

func (o *CloseTradeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error401Unauthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCloseTradeForbidden creates a CloseTradeForbidden with default headers values
func NewCloseTradeForbidden() *CloseTradeForbidden {
	return &CloseTradeForbidden{}
}

/*CloseTradeForbidden handles this case with default header values.

Forbidden
*/
type CloseTradeForbidden struct {
	Payload *models.Error403Forbidden
}

func (o *CloseTradeForbidden) Error() string {
	return fmt.Sprintf("[POST /trading/close_trade][%d] closeTradeForbidden  %+v", 403, o.Payload)
}

func (o *CloseTradeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error403Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCloseTradeNotFound creates a CloseTradeNotFound with default headers values
func NewCloseTradeNotFound() *CloseTradeNotFound {
	return &CloseTradeNotFound{}
}

/*CloseTradeNotFound handles this case with default header values.

Not found
*/
type CloseTradeNotFound struct {
	Payload *models.Error404NotFound
}

func (o *CloseTradeNotFound) Error() string {
	return fmt.Sprintf("[POST /trading/close_trade][%d] closeTradeNotFound  %+v", 404, o.Payload)
}

func (o *CloseTradeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error404NotFound)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}