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

package ipam

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
	"github.com/go-openapi/swag"

	"github.com/thunderstorm99/go-netbox/v3/netbox/models"
)

// NewIpamServicesPartialUpdateParams creates a new IpamServicesPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIpamServicesPartialUpdateParams() *IpamServicesPartialUpdateParams {
	return &IpamServicesPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIpamServicesPartialUpdateParamsWithTimeout creates a new IpamServicesPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewIpamServicesPartialUpdateParamsWithTimeout(timeout time.Duration) *IpamServicesPartialUpdateParams {
	return &IpamServicesPartialUpdateParams{
		timeout: timeout,
	}
}

// NewIpamServicesPartialUpdateParamsWithContext creates a new IpamServicesPartialUpdateParams object
// with the ability to set a context for a request.
func NewIpamServicesPartialUpdateParamsWithContext(ctx context.Context) *IpamServicesPartialUpdateParams {
	return &IpamServicesPartialUpdateParams{
		Context: ctx,
	}
}

// NewIpamServicesPartialUpdateParamsWithHTTPClient creates a new IpamServicesPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewIpamServicesPartialUpdateParamsWithHTTPClient(client *http.Client) *IpamServicesPartialUpdateParams {
	return &IpamServicesPartialUpdateParams{
		HTTPClient: client,
	}
}

/*
IpamServicesPartialUpdateParams contains all the parameters to send to the API endpoint

	for the ipam services partial update operation.

	Typically these are written to a http.Request.
*/
type IpamServicesPartialUpdateParams struct {

	// Data.
	Data *models.WritableService

	/* ID.

	   A unique integer value identifying this service.
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ipam services partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamServicesPartialUpdateParams) WithDefaults() *IpamServicesPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ipam services partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamServicesPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ipam services partial update params
func (o *IpamServicesPartialUpdateParams) WithTimeout(timeout time.Duration) *IpamServicesPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam services partial update params
func (o *IpamServicesPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam services partial update params
func (o *IpamServicesPartialUpdateParams) WithContext(ctx context.Context) *IpamServicesPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam services partial update params
func (o *IpamServicesPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam services partial update params
func (o *IpamServicesPartialUpdateParams) WithHTTPClient(client *http.Client) *IpamServicesPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam services partial update params
func (o *IpamServicesPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the ipam services partial update params
func (o *IpamServicesPartialUpdateParams) WithData(data *models.WritableService) *IpamServicesPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the ipam services partial update params
func (o *IpamServicesPartialUpdateParams) SetData(data *models.WritableService) {
	o.Data = data
}

// WithID adds the id to the ipam services partial update params
func (o *IpamServicesPartialUpdateParams) WithID(id int64) *IpamServicesPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ipam services partial update params
func (o *IpamServicesPartialUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *IpamServicesPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
