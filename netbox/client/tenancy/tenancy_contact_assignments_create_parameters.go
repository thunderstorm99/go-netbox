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

// NewTenancyContactAssignmentsCreateParams creates a new TenancyContactAssignmentsCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTenancyContactAssignmentsCreateParams() *TenancyContactAssignmentsCreateParams {
	return &TenancyContactAssignmentsCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTenancyContactAssignmentsCreateParamsWithTimeout creates a new TenancyContactAssignmentsCreateParams object
// with the ability to set a timeout on a request.
func NewTenancyContactAssignmentsCreateParamsWithTimeout(timeout time.Duration) *TenancyContactAssignmentsCreateParams {
	return &TenancyContactAssignmentsCreateParams{
		timeout: timeout,
	}
}

// NewTenancyContactAssignmentsCreateParamsWithContext creates a new TenancyContactAssignmentsCreateParams object
// with the ability to set a context for a request.
func NewTenancyContactAssignmentsCreateParamsWithContext(ctx context.Context) *TenancyContactAssignmentsCreateParams {
	return &TenancyContactAssignmentsCreateParams{
		Context: ctx,
	}
}

// NewTenancyContactAssignmentsCreateParamsWithHTTPClient creates a new TenancyContactAssignmentsCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewTenancyContactAssignmentsCreateParamsWithHTTPClient(client *http.Client) *TenancyContactAssignmentsCreateParams {
	return &TenancyContactAssignmentsCreateParams{
		HTTPClient: client,
	}
}

/*
TenancyContactAssignmentsCreateParams contains all the parameters to send to the API endpoint

	for the tenancy contact assignments create operation.

	Typically these are written to a http.Request.
*/
type TenancyContactAssignmentsCreateParams struct {

	// Data.
	Data *models.WritableContactAssignment

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the tenancy contact assignments create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TenancyContactAssignmentsCreateParams) WithDefaults() *TenancyContactAssignmentsCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the tenancy contact assignments create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TenancyContactAssignmentsCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the tenancy contact assignments create params
func (o *TenancyContactAssignmentsCreateParams) WithTimeout(timeout time.Duration) *TenancyContactAssignmentsCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the tenancy contact assignments create params
func (o *TenancyContactAssignmentsCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the tenancy contact assignments create params
func (o *TenancyContactAssignmentsCreateParams) WithContext(ctx context.Context) *TenancyContactAssignmentsCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the tenancy contact assignments create params
func (o *TenancyContactAssignmentsCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the tenancy contact assignments create params
func (o *TenancyContactAssignmentsCreateParams) WithHTTPClient(client *http.Client) *TenancyContactAssignmentsCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the tenancy contact assignments create params
func (o *TenancyContactAssignmentsCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the tenancy contact assignments create params
func (o *TenancyContactAssignmentsCreateParams) WithData(data *models.WritableContactAssignment) *TenancyContactAssignmentsCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the tenancy contact assignments create params
func (o *TenancyContactAssignmentsCreateParams) SetData(data *models.WritableContactAssignment) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *TenancyContactAssignmentsCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
