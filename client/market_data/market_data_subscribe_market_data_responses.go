// Code generated by go-swagger; DO NOT EDIT.

package market_data

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/nii236/fxcm-go/models"
)

// MarketDataSubscribeMarketDataReader is a Reader for the MarketDataSubscribeMarketData structure.
type MarketDataSubscribeMarketDataReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MarketDataSubscribeMarketDataReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewMarketDataSubscribeMarketDataOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewMarketDataSubscribeMarketDataUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewMarketDataSubscribeMarketDataForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewMarketDataSubscribeMarketDataNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewMarketDataSubscribeMarketDataOK creates a MarketDataSubscribeMarketDataOK with default headers values
func NewMarketDataSubscribeMarketDataOK() *MarketDataSubscribeMarketDataOK {
	return &MarketDataSubscribeMarketDataOK{}
}

/*MarketDataSubscribeMarketDataOK handles this case with default header values.

OK
*/
type MarketDataSubscribeMarketDataOK struct {
	Payload *models.SubscribeMarketDataResponseDefinition
}

func (o *MarketDataSubscribeMarketDataOK) Error() string {
	return fmt.Sprintf("[POST /subscribe][%d] marketDataSubscribeMarketDataOK  %+v", 200, o.Payload)
}

func (o *MarketDataSubscribeMarketDataOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SubscribeMarketDataResponseDefinition)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMarketDataSubscribeMarketDataUnauthorized creates a MarketDataSubscribeMarketDataUnauthorized with default headers values
func NewMarketDataSubscribeMarketDataUnauthorized() *MarketDataSubscribeMarketDataUnauthorized {
	return &MarketDataSubscribeMarketDataUnauthorized{}
}

/*MarketDataSubscribeMarketDataUnauthorized handles this case with default header values.

Unauthorized
*/
type MarketDataSubscribeMarketDataUnauthorized struct {
	Payload *models.Error401Unauthorized
}

func (o *MarketDataSubscribeMarketDataUnauthorized) Error() string {
	return fmt.Sprintf("[POST /subscribe][%d] marketDataSubscribeMarketDataUnauthorized  %+v", 401, o.Payload)
}

func (o *MarketDataSubscribeMarketDataUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error401Unauthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMarketDataSubscribeMarketDataForbidden creates a MarketDataSubscribeMarketDataForbidden with default headers values
func NewMarketDataSubscribeMarketDataForbidden() *MarketDataSubscribeMarketDataForbidden {
	return &MarketDataSubscribeMarketDataForbidden{}
}

/*MarketDataSubscribeMarketDataForbidden handles this case with default header values.

Forbidden
*/
type MarketDataSubscribeMarketDataForbidden struct {
	Payload *models.Error403Forbidden
}

func (o *MarketDataSubscribeMarketDataForbidden) Error() string {
	return fmt.Sprintf("[POST /subscribe][%d] marketDataSubscribeMarketDataForbidden  %+v", 403, o.Payload)
}

func (o *MarketDataSubscribeMarketDataForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error403Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMarketDataSubscribeMarketDataNotFound creates a MarketDataSubscribeMarketDataNotFound with default headers values
func NewMarketDataSubscribeMarketDataNotFound() *MarketDataSubscribeMarketDataNotFound {
	return &MarketDataSubscribeMarketDataNotFound{}
}

/*MarketDataSubscribeMarketDataNotFound handles this case with default header values.

Not found
*/
type MarketDataSubscribeMarketDataNotFound struct {
	Payload *models.Error404NotFound
}

func (o *MarketDataSubscribeMarketDataNotFound) Error() string {
	return fmt.Sprintf("[POST /subscribe][%d] marketDataSubscribeMarketDataNotFound  %+v", 404, o.Payload)
}

func (o *MarketDataSubscribeMarketDataNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error404NotFound)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}