// Code generated by go-swagger; DO NOT EDIT.

package trading_orders

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/nii236/fxcm-go/models"
)

// NewEditOCOParams creates a new EditOCOParams object
// with the default values initialized.
func NewEditOCOParams() *EditOCOParams {
	var ()
	return &EditOCOParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewEditOCOParamsWithTimeout creates a new EditOCOParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewEditOCOParamsWithTimeout(timeout time.Duration) *EditOCOParams {
	var ()
	return &EditOCOParams{

		timeout: timeout,
	}
}

// NewEditOCOParamsWithContext creates a new EditOCOParams object
// with the default values initialized, and the ability to set a context for a request
func NewEditOCOParamsWithContext(ctx context.Context) *EditOCOParams {
	var ()
	return &EditOCOParams{

		Context: ctx,
	}
}

// NewEditOCOParamsWithHTTPClient creates a new EditOCOParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewEditOCOParamsWithHTTPClient(client *http.Client) *EditOCOParams {
	var ()
	return &EditOCOParams{
		HTTPClient: client,
	}
}

/*EditOCOParams contains all the parameters to send to the API endpoint
for the edit o c o operation typically these are written to a http.Request
*/
type EditOCOParams struct {

	/*AddOrderID
	  The list orders identifiers to add to the oco order.

	*/
	AddOrderID float64
	/*Body*/
	Body models.EditOCORequestDefinition
	/*OcoBulkID
	  The oco bulk identifier (if equals zero then new oco order will be created).

	*/
	OcoBulkID float64
	/*RemoveOrderID
	  The list orders identifiers to remove from the oco order.

	*/
	RemoveOrderID float64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the edit o c o params
func (o *EditOCOParams) WithTimeout(timeout time.Duration) *EditOCOParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the edit o c o params
func (o *EditOCOParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the edit o c o params
func (o *EditOCOParams) WithContext(ctx context.Context) *EditOCOParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the edit o c o params
func (o *EditOCOParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the edit o c o params
func (o *EditOCOParams) WithHTTPClient(client *http.Client) *EditOCOParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the edit o c o params
func (o *EditOCOParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAddOrderID adds the addOrderID to the edit o c o params
func (o *EditOCOParams) WithAddOrderID(addOrderID float64) *EditOCOParams {
	o.SetAddOrderID(addOrderID)
	return o
}

// SetAddOrderID adds the addOrderId to the edit o c o params
func (o *EditOCOParams) SetAddOrderID(addOrderID float64) {
	o.AddOrderID = addOrderID
}

// WithBody adds the body to the edit o c o params
func (o *EditOCOParams) WithBody(body models.EditOCORequestDefinition) *EditOCOParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the edit o c o params
func (o *EditOCOParams) SetBody(body models.EditOCORequestDefinition) {
	o.Body = body
}

// WithOcoBulkID adds the ocoBulkID to the edit o c o params
func (o *EditOCOParams) WithOcoBulkID(ocoBulkID float64) *EditOCOParams {
	o.SetOcoBulkID(ocoBulkID)
	return o
}

// SetOcoBulkID adds the ocoBulkId to the edit o c o params
func (o *EditOCOParams) SetOcoBulkID(ocoBulkID float64) {
	o.OcoBulkID = ocoBulkID
}

// WithRemoveOrderID adds the removeOrderID to the edit o c o params
func (o *EditOCOParams) WithRemoveOrderID(removeOrderID float64) *EditOCOParams {
	o.SetRemoveOrderID(removeOrderID)
	return o
}

// SetRemoveOrderID adds the removeOrderId to the edit o c o params
func (o *EditOCOParams) SetRemoveOrderID(removeOrderID float64) {
	o.RemoveOrderID = removeOrderID
}

// WriteToRequest writes these params to a swagger request
func (o *EditOCOParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param addOrderId
	qrAddOrderID := o.AddOrderID
	qAddOrderID := swag.FormatFloat64(qrAddOrderID)
	if qAddOrderID != "" {
		if err := r.SetQueryParam("addOrderId", qAddOrderID); err != nil {
			return err
		}
	}

	// query param ocoBulkId
	qrOcoBulkID := o.OcoBulkID
	qOcoBulkID := swag.FormatFloat64(qrOcoBulkID)
	if qOcoBulkID != "" {
		if err := r.SetQueryParam("ocoBulkId", qOcoBulkID); err != nil {
			return err
		}
	}

	// query param removeOrderId
	qrRemoveOrderID := o.RemoveOrderID
	qRemoveOrderID := swag.FormatFloat64(qrRemoveOrderID)
	if qRemoveOrderID != "" {
		if err := r.SetQueryParam("removeOrderId", qRemoveOrderID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
