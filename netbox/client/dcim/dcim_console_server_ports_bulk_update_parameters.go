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

// NewDcimConsoleServerPortsBulkUpdateParams creates a new DcimConsoleServerPortsBulkUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimConsoleServerPortsBulkUpdateParams() *DcimConsoleServerPortsBulkUpdateParams {
	return &DcimConsoleServerPortsBulkUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimConsoleServerPortsBulkUpdateParamsWithTimeout creates a new DcimConsoleServerPortsBulkUpdateParams object
// with the ability to set a timeout on a request.
func NewDcimConsoleServerPortsBulkUpdateParamsWithTimeout(timeout time.Duration) *DcimConsoleServerPortsBulkUpdateParams {
	return &DcimConsoleServerPortsBulkUpdateParams{
		timeout: timeout,
	}
}

// NewDcimConsoleServerPortsBulkUpdateParamsWithContext creates a new DcimConsoleServerPortsBulkUpdateParams object
// with the ability to set a context for a request.
func NewDcimConsoleServerPortsBulkUpdateParamsWithContext(ctx context.Context) *DcimConsoleServerPortsBulkUpdateParams {
	return &DcimConsoleServerPortsBulkUpdateParams{
		Context: ctx,
	}
}

// NewDcimConsoleServerPortsBulkUpdateParamsWithHTTPClient creates a new DcimConsoleServerPortsBulkUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimConsoleServerPortsBulkUpdateParamsWithHTTPClient(client *http.Client) *DcimConsoleServerPortsBulkUpdateParams {
	return &DcimConsoleServerPortsBulkUpdateParams{
		HTTPClient: client,
	}
}

/*
DcimConsoleServerPortsBulkUpdateParams contains all the parameters to send to the API endpoint

	for the dcim console server ports bulk update operation.

	Typically these are written to a http.Request.
*/
type DcimConsoleServerPortsBulkUpdateParams struct {

	// Data.
	Data *models.WritableConsoleServerPort

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim console server ports bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimConsoleServerPortsBulkUpdateParams) WithDefaults() *DcimConsoleServerPortsBulkUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim console server ports bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimConsoleServerPortsBulkUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim console server ports bulk update params
func (o *DcimConsoleServerPortsBulkUpdateParams) WithTimeout(timeout time.Duration) *DcimConsoleServerPortsBulkUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim console server ports bulk update params
func (o *DcimConsoleServerPortsBulkUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim console server ports bulk update params
func (o *DcimConsoleServerPortsBulkUpdateParams) WithContext(ctx context.Context) *DcimConsoleServerPortsBulkUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim console server ports bulk update params
func (o *DcimConsoleServerPortsBulkUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim console server ports bulk update params
func (o *DcimConsoleServerPortsBulkUpdateParams) WithHTTPClient(client *http.Client) *DcimConsoleServerPortsBulkUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim console server ports bulk update params
func (o *DcimConsoleServerPortsBulkUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim console server ports bulk update params
func (o *DcimConsoleServerPortsBulkUpdateParams) WithData(data *models.WritableConsoleServerPort) *DcimConsoleServerPortsBulkUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim console server ports bulk update params
func (o *DcimConsoleServerPortsBulkUpdateParams) SetData(data *models.WritableConsoleServerPort) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *DcimConsoleServerPortsBulkUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
