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

package circuits

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

// NewCircuitsProviderNetworksCreateParams creates a new CircuitsProviderNetworksCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCircuitsProviderNetworksCreateParams() *CircuitsProviderNetworksCreateParams {
	return &CircuitsProviderNetworksCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCircuitsProviderNetworksCreateParamsWithTimeout creates a new CircuitsProviderNetworksCreateParams object
// with the ability to set a timeout on a request.
func NewCircuitsProviderNetworksCreateParamsWithTimeout(timeout time.Duration) *CircuitsProviderNetworksCreateParams {
	return &CircuitsProviderNetworksCreateParams{
		timeout: timeout,
	}
}

// NewCircuitsProviderNetworksCreateParamsWithContext creates a new CircuitsProviderNetworksCreateParams object
// with the ability to set a context for a request.
func NewCircuitsProviderNetworksCreateParamsWithContext(ctx context.Context) *CircuitsProviderNetworksCreateParams {
	return &CircuitsProviderNetworksCreateParams{
		Context: ctx,
	}
}

// NewCircuitsProviderNetworksCreateParamsWithHTTPClient creates a new CircuitsProviderNetworksCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewCircuitsProviderNetworksCreateParamsWithHTTPClient(client *http.Client) *CircuitsProviderNetworksCreateParams {
	return &CircuitsProviderNetworksCreateParams{
		HTTPClient: client,
	}
}

/*
CircuitsProviderNetworksCreateParams contains all the parameters to send to the API endpoint

	for the circuits provider networks create operation.

	Typically these are written to a http.Request.
*/
type CircuitsProviderNetworksCreateParams struct {

	// Data.
	Data *models.WritableProviderNetwork

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the circuits provider networks create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CircuitsProviderNetworksCreateParams) WithDefaults() *CircuitsProviderNetworksCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the circuits provider networks create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CircuitsProviderNetworksCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the circuits provider networks create params
func (o *CircuitsProviderNetworksCreateParams) WithTimeout(timeout time.Duration) *CircuitsProviderNetworksCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the circuits provider networks create params
func (o *CircuitsProviderNetworksCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the circuits provider networks create params
func (o *CircuitsProviderNetworksCreateParams) WithContext(ctx context.Context) *CircuitsProviderNetworksCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the circuits provider networks create params
func (o *CircuitsProviderNetworksCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the circuits provider networks create params
func (o *CircuitsProviderNetworksCreateParams) WithHTTPClient(client *http.Client) *CircuitsProviderNetworksCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the circuits provider networks create params
func (o *CircuitsProviderNetworksCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the circuits provider networks create params
func (o *CircuitsProviderNetworksCreateParams) WithData(data *models.WritableProviderNetwork) *CircuitsProviderNetworksCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the circuits provider networks create params
func (o *CircuitsProviderNetworksCreateParams) SetData(data *models.WritableProviderNetwork) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *CircuitsProviderNetworksCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
