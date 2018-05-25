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

// ChangeTradeStopLimitReader is a Reader for the ChangeTradeStopLimit structure.
type ChangeTradeStopLimitReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ChangeTradeStopLimitReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewChangeTradeStopLimitOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewChangeTradeStopLimitUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewChangeTradeStopLimitForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewChangeTradeStopLimitNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewChangeTradeStopLimitOK creates a ChangeTradeStopLimitOK with default headers values
func NewChangeTradeStopLimitOK() *ChangeTradeStopLimitOK {
	return &ChangeTradeStopLimitOK{}
}

/*ChangeTradeStopLimitOK handles this case with default header values.

OK
*/
type ChangeTradeStopLimitOK struct {
	Payload *models.TraderOrdersResponseDefinition
}

func (o *ChangeTradeStopLimitOK) Error() string {
	return fmt.Sprintf("[POST /trading/change_trade_stop_limit][%d] changeTradeStopLimitOK  %+v", 200, o.Payload)
}

func (o *ChangeTradeStopLimitOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TraderOrdersResponseDefinition)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewChangeTradeStopLimitUnauthorized creates a ChangeTradeStopLimitUnauthorized with default headers values
func NewChangeTradeStopLimitUnauthorized() *ChangeTradeStopLimitUnauthorized {
	return &ChangeTradeStopLimitUnauthorized{}
}

/*ChangeTradeStopLimitUnauthorized handles this case with default header values.

Unauthorized
*/
type ChangeTradeStopLimitUnauthorized struct {
	Payload *models.Error401Unauthorized
}

func (o *ChangeTradeStopLimitUnauthorized) Error() string {
	return fmt.Sprintf("[POST /trading/change_trade_stop_limit][%d] changeTradeStopLimitUnauthorized  %+v", 401, o.Payload)
}

func (o *ChangeTradeStopLimitUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error401Unauthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewChangeTradeStopLimitForbidden creates a ChangeTradeStopLimitForbidden with default headers values
func NewChangeTradeStopLimitForbidden() *ChangeTradeStopLimitForbidden {
	return &ChangeTradeStopLimitForbidden{}
}

/*ChangeTradeStopLimitForbidden handles this case with default header values.

Forbidden
*/
type ChangeTradeStopLimitForbidden struct {
	Payload *models.Error403Forbidden
}

func (o *ChangeTradeStopLimitForbidden) Error() string {
	return fmt.Sprintf("[POST /trading/change_trade_stop_limit][%d] changeTradeStopLimitForbidden  %+v", 403, o.Payload)
}

func (o *ChangeTradeStopLimitForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error403Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewChangeTradeStopLimitNotFound creates a ChangeTradeStopLimitNotFound with default headers values
func NewChangeTradeStopLimitNotFound() *ChangeTradeStopLimitNotFound {
	return &ChangeTradeStopLimitNotFound{}
}

/*ChangeTradeStopLimitNotFound handles this case with default header values.

Not found
*/
type ChangeTradeStopLimitNotFound struct {
	Payload *models.Error404NotFound
}

func (o *ChangeTradeStopLimitNotFound) Error() string {
	return fmt.Sprintf("[POST /trading/change_trade_stop_limit][%d] changeTradeStopLimitNotFound  %+v", 404, o.Payload)
}

func (o *ChangeTradeStopLimitNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error404NotFound)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
