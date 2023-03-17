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
	"github.com/go-openapi/swag"

	"github.com/thunderstorm99/go-netbox/v3/netbox/models"
)

// NewDcimCableTerminationsUpdateParams creates a new DcimCableTerminationsUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimCableTerminationsUpdateParams() *DcimCableTerminationsUpdateParams {
	return &DcimCableTerminationsUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimCableTerminationsUpdateParamsWithTimeout creates a new DcimCableTerminationsUpdateParams object
// with the ability to set a timeout on a request.
func NewDcimCableTerminationsUpdateParamsWithTimeout(timeout time.Duration) *DcimCableTerminationsUpdateParams {
	return &DcimCableTerminationsUpdateParams{
		timeout: timeout,
	}
}

// NewDcimCableTerminationsUpdateParamsWithContext creates a new DcimCableTerminationsUpdateParams object
// with the ability to set a context for a request.
func NewDcimCableTerminationsUpdateParamsWithContext(ctx context.Context) *DcimCableTerminationsUpdateParams {
	return &DcimCableTerminationsUpdateParams{
		Context: ctx,
	}
}

// NewDcimCableTerminationsUpdateParamsWithHTTPClient creates a new DcimCableTerminationsUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimCableTerminationsUpdateParamsWithHTTPClient(client *http.Client) *DcimCableTerminationsUpdateParams {
	return &DcimCableTerminationsUpdateParams{
		HTTPClient: client,
	}
}

/*
DcimCableTerminationsUpdateParams contains all the parameters to send to the API endpoint

	for the dcim cable terminations update operation.

	Typically these are written to a http.Request.
*/
type DcimCableTerminationsUpdateParams struct {

	// Data.
	Data *models.CableTermination

	/* ID.

	   A unique integer value identifying this cable termination.
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim cable terminations update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimCableTerminationsUpdateParams) WithDefaults() *DcimCableTerminationsUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim cable terminations update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimCableTerminationsUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim cable terminations update params
func (o *DcimCableTerminationsUpdateParams) WithTimeout(timeout time.Duration) *DcimCableTerminationsUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim cable terminations update params
func (o *DcimCableTerminationsUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim cable terminations update params
func (o *DcimCableTerminationsUpdateParams) WithContext(ctx context.Context) *DcimCableTerminationsUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim cable terminations update params
func (o *DcimCableTerminationsUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim cable terminations update params
func (o *DcimCableTerminationsUpdateParams) WithHTTPClient(client *http.Client) *DcimCableTerminationsUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim cable terminations update params
func (o *DcimCableTerminationsUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim cable terminations update params
func (o *DcimCableTerminationsUpdateParams) WithData(data *models.CableTermination) *DcimCableTerminationsUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim cable terminations update params
func (o *DcimCableTerminationsUpdateParams) SetData(data *models.CableTermination) {
	o.Data = data
}

// WithID adds the id to the dcim cable terminations update params
func (o *DcimCableTerminationsUpdateParams) WithID(id int64) *DcimCableTerminationsUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim cable terminations update params
func (o *DcimCableTerminationsUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimCableTerminationsUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
