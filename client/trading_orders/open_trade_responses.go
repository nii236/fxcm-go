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

// OpenTradeReader is a Reader for the OpenTrade structure.
type OpenTradeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OpenTradeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewOpenTradeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewOpenTradeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewOpenTradeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewOpenTradeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewOpenTradeOK creates a OpenTradeOK with default headers values
func NewOpenTradeOK() *OpenTradeOK {
	return &OpenTradeOK{}
}

/*OpenTradeOK handles this case with default header values.

OK
*/
type OpenTradeOK struct {
	Payload *models.TraderOrdersResponseDefinition
}

func (o *OpenTradeOK) Error() string {
	return fmt.Sprintf("[POST /trading/open_trade][%d] openTradeOK  %+v", 200, o.Payload)
}

func (o *OpenTradeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TraderOrdersResponseDefinition)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOpenTradeUnauthorized creates a OpenTradeUnauthorized with default headers values
func NewOpenTradeUnauthorized() *OpenTradeUnauthorized {
	return &OpenTradeUnauthorized{}
}

/*OpenTradeUnauthorized handles this case with default header values.

Unauthorized
*/
type OpenTradeUnauthorized struct {
	Payload *models.Error401Unauthorized
}

func (o *OpenTradeUnauthorized) Error() string {
	return fmt.Sprintf("[POST /trading/open_trade][%d] openTradeUnauthorized  %+v", 401, o.Payload)
}

func (o *OpenTradeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error401Unauthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOpenTradeForbidden creates a OpenTradeForbidden with default headers values
func NewOpenTradeForbidden() *OpenTradeForbidden {
	return &OpenTradeForbidden{}
}

/*OpenTradeForbidden handles this case with default header values.

Forbidden
*/
type OpenTradeForbidden struct {
	Payload *models.Error403Forbidden
}

func (o *OpenTradeForbidden) Error() string {
	return fmt.Sprintf("[POST /trading/open_trade][%d] openTradeForbidden  %+v", 403, o.Payload)
}

func (o *OpenTradeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error403Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOpenTradeNotFound creates a OpenTradeNotFound with default headers values
func NewOpenTradeNotFound() *OpenTradeNotFound {
	return &OpenTradeNotFound{}
}

/*OpenTradeNotFound handles this case with default header values.

Not found
*/
type OpenTradeNotFound struct {
	Payload *models.Error404NotFound
}

func (o *OpenTradeNotFound) Error() string {
	return fmt.Sprintf("[POST /trading/open_trade][%d] openTradeNotFound  %+v", 404, o.Payload)
}

func (o *OpenTradeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error404NotFound)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
