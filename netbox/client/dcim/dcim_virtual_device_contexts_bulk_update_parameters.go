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

// NewDcimVirtualDeviceContextsBulkUpdateParams creates a new DcimVirtualDeviceContextsBulkUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimVirtualDeviceContextsBulkUpdateParams() *DcimVirtualDeviceContextsBulkUpdateParams {
	return &DcimVirtualDeviceContextsBulkUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimVirtualDeviceContextsBulkUpdateParamsWithTimeout creates a new DcimVirtualDeviceContextsBulkUpdateParams object
// with the ability to set a timeout on a request.
func NewDcimVirtualDeviceContextsBulkUpdateParamsWithTimeout(timeout time.Duration) *DcimVirtualDeviceContextsBulkUpdateParams {
	return &DcimVirtualDeviceContextsBulkUpdateParams{
		timeout: timeout,
	}
}

// NewDcimVirtualDeviceContextsBulkUpdateParamsWithContext creates a new DcimVirtualDeviceContextsBulkUpdateParams object
// with the ability to set a context for a request.
func NewDcimVirtualDeviceContextsBulkUpdateParamsWithContext(ctx context.Context) *DcimVirtualDeviceContextsBulkUpdateParams {
	return &DcimVirtualDeviceContextsBulkUpdateParams{
		Context: ctx,
	}
}

// NewDcimVirtualDeviceContextsBulkUpdateParamsWithHTTPClient creates a new DcimVirtualDeviceContextsBulkUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimVirtualDeviceContextsBulkUpdateParamsWithHTTPClient(client *http.Client) *DcimVirtualDeviceContextsBulkUpdateParams {
	return &DcimVirtualDeviceContextsBulkUpdateParams{
		HTTPClient: client,
	}
}

/*
DcimVirtualDeviceContextsBulkUpdateParams contains all the parameters to send to the API endpoint

	for the dcim virtual device contexts bulk update operation.

	Typically these are written to a http.Request.
*/
type DcimVirtualDeviceContextsBulkUpdateParams struct {

	// Data.
	Data *models.WritableVirtualDeviceContext

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim virtual device contexts bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimVirtualDeviceContextsBulkUpdateParams) WithDefaults() *DcimVirtualDeviceContextsBulkUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim virtual device contexts bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimVirtualDeviceContextsBulkUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim virtual device contexts bulk update params
func (o *DcimVirtualDeviceContextsBulkUpdateParams) WithTimeout(timeout time.Duration) *DcimVirtualDeviceContextsBulkUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim virtual device contexts bulk update params
func (o *DcimVirtualDeviceContextsBulkUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim virtual device contexts bulk update params
func (o *DcimVirtualDeviceContextsBulkUpdateParams) WithContext(ctx context.Context) *DcimVirtualDeviceContextsBulkUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim virtual device contexts bulk update params
func (o *DcimVirtualDeviceContextsBulkUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim virtual device contexts bulk update params
func (o *DcimVirtualDeviceContextsBulkUpdateParams) WithHTTPClient(client *http.Client) *DcimVirtualDeviceContextsBulkUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim virtual device contexts bulk update params
func (o *DcimVirtualDeviceContextsBulkUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim virtual device contexts bulk update params
func (o *DcimVirtualDeviceContextsBulkUpdateParams) WithData(data *models.WritableVirtualDeviceContext) *DcimVirtualDeviceContextsBulkUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim virtual device contexts bulk update params
func (o *DcimVirtualDeviceContextsBulkUpdateParams) SetData(data *models.WritableVirtualDeviceContext) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *DcimVirtualDeviceContextsBulkUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
