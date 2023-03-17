// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2020 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package dcim

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/thunderstorm99/go-netbox/v3/netbox/models"
)

// NewDcimRearPortTemplatesBulkUpdateParams creates a new DcimRearPortTemplatesBulkUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimRearPortTemplatesBulkUpdateParams() *DcimRearPortTemplatesBulkUpdateParams {
	return &DcimRearPortTemplatesBulkUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimRearPortTemplatesBulkUpdateParamsWithTimeout creates a new DcimRearPortTemplatesBulkUpdateParams object
// with the ability to set a timeout on a request.
func NewDcimRearPortTemplatesBulkUpdateParamsWithTimeout(timeout time.Duration) *DcimRearPortTemplatesBulkUpdateParams {
	return &DcimRearPortTemplatesBulkUpdateParams{
		timeout: timeout,
	}
}

// NewDcimRearPortTemplatesBulkUpdateParamsWithContext creates a new DcimRearPortTemplatesBulkUpdateParams object
// with the ability to set a context for a request.
func NewDcimRearPortTemplatesBulkUpdateParamsWithContext(ctx context.Context) *DcimRearPortTemplatesBulkUpdateParams {
	return &DcimRearPortTemplatesBulkUpdateParams{
		Context: ctx,
	}
}

// NewDcimRearPortTemplatesBulkUpdateParamsWithHTTPClient creates a new DcimRearPortTemplatesBulkUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimRearPortTemplatesBulkUpdateParamsWithHTTPClient(client *http.Client) *DcimRearPortTemplatesBulkUpdateParams {
	return &DcimRearPortTemplatesBulkUpdateParams{
		HTTPClient: client,
	}
}

/*
DcimRearPortTemplatesBulkUpdateParams contains all the parameters to send to the API endpoint

	for the dcim rear port templates bulk update operation.

	Typically these are written to a http.Request.
*/
type DcimRearPortTemplatesBulkUpdateParams struct {

	// Data.
	Data *models.WritableRearPortTemplate

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim rear port templates bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimRearPortTemplatesBulkUpdateParams) WithDefaults() *DcimRearPortTemplatesBulkUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim rear port templates bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimRearPortTemplatesBulkUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim rear port templates bulk update params
func (o *DcimRearPortTemplatesBulkUpdateParams) WithTimeout(timeout time.Duration) *DcimRearPortTemplatesBulkUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim rear port templates bulk update params
func (o *DcimRearPortTemplatesBulkUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim rear port templates bulk update params
func (o *DcimRearPortTemplatesBulkUpdateParams) WithContext(ctx context.Context) *DcimRearPortTemplatesBulkUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim rear port templates bulk update params
func (o *DcimRearPortTemplatesBulkUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim rear port templates bulk update params
func (o *DcimRearPortTemplatesBulkUpdateParams) WithHTTPClient(client *http.Client) *DcimRearPortTemplatesBulkUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim rear port templates bulk update params
func (o *DcimRearPortTemplatesBulkUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim rear port templates bulk update params
func (o *DcimRearPortTemplatesBulkUpdateParams) WithData(data *models.WritableRearPortTemplate) *DcimRearPortTemplatesBulkUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim rear port templates bulk update params
func (o *DcimRearPortTemplatesBulkUpdateParams) SetData(data *models.WritableRearPortTemplate) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *DcimRearPortTemplatesBulkUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
