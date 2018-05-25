// Code generated by go-swagger; DO NOT EDIT.

package trading_tables

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/nii236/fxcm-go/models"
)

// TradingTableSummaryTableReader is a Reader for the TradingTableSummaryTable structure.
type TradingTableSummaryTableReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TradingTableSummaryTableReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewTradingTableSummaryTableOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewTradingTableSummaryTableUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewTradingTableSummaryTableForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewTradingTableSummaryTableNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewTradingTableSummaryTableOK creates a TradingTableSummaryTableOK with default headers values
func NewTradingTableSummaryTableOK() *TradingTableSummaryTableOK {
	return &TradingTableSummaryTableOK{}
}

/*TradingTableSummaryTableOK handles this case with default header values.

OK
*/
type TradingTableSummaryTableOK struct {
	Payload *models.SummaryTableResponseDefinition
}

func (o *TradingTableSummaryTableOK) Error() string {
	return fmt.Sprintf("[GET /trading/get_model&model=Summary][%d] tradingTableSummaryTableOK  %+v", 200, o.Payload)
}

func (o *TradingTableSummaryTableOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SummaryTableResponseDefinition)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTradingTableSummaryTableUnauthorized creates a TradingTableSummaryTableUnauthorized with default headers values
func NewTradingTableSummaryTableUnauthorized() *TradingTableSummaryTableUnauthorized {
	return &TradingTableSummaryTableUnauthorized{}
}

/*TradingTableSummaryTableUnauthorized handles this case with default header values.

Unauthorized
*/
type TradingTableSummaryTableUnauthorized struct {
	Payload *models.Error401Unauthorized
}

func (o *TradingTableSummaryTableUnauthorized) Error() string {
	return fmt.Sprintf("[GET /trading/get_model&model=Summary][%d] tradingTableSummaryTableUnauthorized  %+v", 401, o.Payload)
}

func (o *TradingTableSummaryTableUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error401Unauthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTradingTableSummaryTableForbidden creates a TradingTableSummaryTableForbidden with default headers values
func NewTradingTableSummaryTableForbidden() *TradingTableSummaryTableForbidden {
	return &TradingTableSummaryTableForbidden{}
}

/*TradingTableSummaryTableForbidden handles this case with default header values.

Forbidden
*/
type TradingTableSummaryTableForbidden struct {
	Payload *models.Error403Forbidden
}

func (o *TradingTableSummaryTableForbidden) Error() string {
	return fmt.Sprintf("[GET /trading/get_model&model=Summary][%d] tradingTableSummaryTableForbidden  %+v", 403, o.Payload)
}

func (o *TradingTableSummaryTableForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error403Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTradingTableSummaryTableNotFound creates a TradingTableSummaryTableNotFound with default headers values
func NewTradingTableSummaryTableNotFound() *TradingTableSummaryTableNotFound {
	return &TradingTableSummaryTableNotFound{}
}

/*TradingTableSummaryTableNotFound handles this case with default header values.

Not found
*/
type TradingTableSummaryTableNotFound struct {
	Payload *models.Error404NotFound
}

func (o *TradingTableSummaryTableNotFound) Error() string {
	return fmt.Sprintf("[GET /trading/get_model&model=Summary][%d] tradingTableSummaryTableNotFound  %+v", 404, o.Payload)
}

func (o *TradingTableSummaryTableNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error404NotFound)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}