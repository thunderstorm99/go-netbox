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

package tenancy

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

// NewTenancyContactAssignmentsBulkPartialUpdateParams creates a new TenancyContactAssignmentsBulkPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTenancyContactAssignmentsBulkPartialUpdateParams() *TenancyContactAssignmentsBulkPartialUpdateParams {
	return &TenancyContactAssignmentsBulkPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTenancyContactAssignmentsBulkPartialUpdateParamsWithTimeout creates a new TenancyContactAssignmentsBulkPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewTenancyContactAssignmentsBulkPartialUpdateParamsWithTimeout(timeout time.Duration) *TenancyContactAssignmentsBulkPartialUpdateParams {
	return &TenancyContactAssignmentsBulkPartialUpdateParams{
		timeout: timeout,
	}
}

// NewTenancyContactAssignmentsBulkPartialUpdateParamsWithContext creates a new TenancyContactAssignmentsBulkPartialUpdateParams object
// with the ability to set a context for a request.
func NewTenancyContactAssignmentsBulkPartialUpdateParamsWithContext(ctx context.Context) *TenancyContactAssignmentsBulkPartialUpdateParams {
	return &TenancyContactAssignmentsBulkPartialUpdateParams{
		Context: ctx,
	}
}

// NewTenancyContactAssignmentsBulkPartialUpdateParamsWithHTTPClient creates a new TenancyContactAssignmentsBulkPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewTenancyContactAssignmentsBulkPartialUpdateParamsWithHTTPClient(client *http.Client) *TenancyContactAssignmentsBulkPartialUpdateParams {
	return &TenancyContactAssignmentsBulkPartialUpdateParams{
		HTTPClient: client,
	}
}

/*
TenancyContactAssignmentsBulkPartialUpdateParams contains all the parameters to send to the API endpoint

	for the tenancy contact assignments bulk partial update operation.

	Typically these are written to a http.Request.
*/
type TenancyContactAssignmentsBulkPartialUpdateParams struct {

	// Data.
	Data *models.WritableContactAssignment

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the tenancy contact assignments bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TenancyContactAssignmentsBulkPartialUpdateParams) WithDefaults() *TenancyContactAssignmentsBulkPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the tenancy contact assignments bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TenancyContactAssignmentsBulkPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the tenancy contact assignments bulk partial update params
func (o *TenancyContactAssignmentsBulkPartialUpdateParams) WithTimeout(timeout time.Duration) *TenancyContactAssignmentsBulkPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the tenancy contact assignments bulk partial update params
func (o *TenancyContactAssignmentsBulkPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the tenancy contact assignments bulk partial update params
func (o *TenancyContactAssignmentsBulkPartialUpdateParams) WithContext(ctx context.Context) *TenancyContactAssignmentsBulkPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the tenancy contact assignments bulk partial update params
func (o *TenancyContactAssignmentsBulkPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the tenancy contact assignments bulk partial update params
func (o *TenancyContactAssignmentsBulkPartialUpdateParams) WithHTTPClient(client *http.Client) *TenancyContactAssignmentsBulkPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the tenancy contact assignments bulk partial update params
func (o *TenancyContactAssignmentsBulkPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the tenancy contact assignments bulk partial update params
func (o *TenancyContactAssignmentsBulkPartialUpdateParams) WithData(data *models.WritableContactAssignment) *TenancyContactAssignmentsBulkPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the tenancy contact assignments bulk partial update params
func (o *TenancyContactAssignmentsBulkPartialUpdateParams) SetData(data *models.WritableContactAssignment) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *TenancyContactAssignmentsBulkPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
