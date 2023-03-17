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

// NewIpamVlanGroupsAvailableVlansCreateParams creates a new IpamVlanGroupsAvailableVlansCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIpamVlanGroupsAvailableVlansCreateParams() *IpamVlanGroupsAvailableVlansCreateParams {
	return &IpamVlanGroupsAvailableVlansCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIpamVlanGroupsAvailableVlansCreateParamsWithTimeout creates a new IpamVlanGroupsAvailableVlansCreateParams object
// with the ability to set a timeout on a request.
func NewIpamVlanGroupsAvailableVlansCreateParamsWithTimeout(timeout time.Duration) *IpamVlanGroupsAvailableVlansCreateParams {
	return &IpamVlanGroupsAvailableVlansCreateParams{
		timeout: timeout,
	}
}

// NewIpamVlanGroupsAvailableVlansCreateParamsWithContext creates a new IpamVlanGroupsAvailableVlansCreateParams object
// with the ability to set a context for a request.
func NewIpamVlanGroupsAvailableVlansCreateParamsWithContext(ctx context.Context) *IpamVlanGroupsAvailableVlansCreateParams {
	return &IpamVlanGroupsAvailableVlansCreateParams{
		Context: ctx,
	}
}

// NewIpamVlanGroupsAvailableVlansCreateParamsWithHTTPClient creates a new IpamVlanGroupsAvailableVlansCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewIpamVlanGroupsAvailableVlansCreateParamsWithHTTPClient(client *http.Client) *IpamVlanGroupsAvailableVlansCreateParams {
	return &IpamVlanGroupsAvailableVlansCreateParams{
		HTTPClient: client,
	}
}

/*
IpamVlanGroupsAvailableVlansCreateParams contains all the parameters to send to the API endpoint

	for the ipam vlan groups available vlans create operation.

	Typically these are written to a http.Request.
*/
type IpamVlanGroupsAvailableVlansCreateParams struct {

	// Data.
	Data *models.WritableCreateAvailableVLAN

	/* ID.

	   A unique integer value identifying this VLAN.
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ipam vlan groups available vlans create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamVlanGroupsAvailableVlansCreateParams) WithDefaults() *IpamVlanGroupsAvailableVlansCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ipam vlan groups available vlans create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamVlanGroupsAvailableVlansCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ipam vlan groups available vlans create params
func (o *IpamVlanGroupsAvailableVlansCreateParams) WithTimeout(timeout time.Duration) *IpamVlanGroupsAvailableVlansCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam vlan groups available vlans create params
func (o *IpamVlanGroupsAvailableVlansCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam vlan groups available vlans create params
func (o *IpamVlanGroupsAvailableVlansCreateParams) WithContext(ctx context.Context) *IpamVlanGroupsAvailableVlansCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam vlan groups available vlans create params
func (o *IpamVlanGroupsAvailableVlansCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam vlan groups available vlans create params
func (o *IpamVlanGroupsAvailableVlansCreateParams) WithHTTPClient(client *http.Client) *IpamVlanGroupsAvailableVlansCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam vlan groups available vlans create params
func (o *IpamVlanGroupsAvailableVlansCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the ipam vlan groups available vlans create params
func (o *IpamVlanGroupsAvailableVlansCreateParams) WithData(data *models.WritableCreateAvailableVLAN) *IpamVlanGroupsAvailableVlansCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the ipam vlan groups available vlans create params
func (o *IpamVlanGroupsAvailableVlansCreateParams) SetData(data *models.WritableCreateAvailableVLAN) {
	o.Data = data
}

// WithID adds the id to the ipam vlan groups available vlans create params
func (o *IpamVlanGroupsAvailableVlansCreateParams) WithID(id int64) *IpamVlanGroupsAvailableVlansCreateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ipam vlan groups available vlans create params
func (o *IpamVlanGroupsAvailableVlansCreateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *IpamVlanGroupsAvailableVlansCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
