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

package wireless

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/thunderstorm99/go-netbox/v3/netbox/models"
)

// WirelessWirelessLansListReader is a Reader for the WirelessWirelessLansList structure.
type WirelessWirelessLansListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WirelessWirelessLansListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWirelessWirelessLansListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWirelessWirelessLansListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWirelessWirelessLansListOK creates a WirelessWirelessLansListOK with default headers values
func NewWirelessWirelessLansListOK() *WirelessWirelessLansListOK {
	return &WirelessWirelessLansListOK{}
}

/*
WirelessWirelessLansListOK describes a response with status code 200, with default header values.

WirelessWirelessLansListOK wireless wireless lans list o k
*/
type WirelessWirelessLansListOK struct {
	Payload *WirelessWirelessLansListOKBody
}

// IsSuccess returns true when this wireless wireless lans list o k response has a 2xx status code
func (o *WirelessWirelessLansListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this wireless wireless lans list o k response has a 3xx status code
func (o *WirelessWirelessLansListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this wireless wireless lans list o k response has a 4xx status code
func (o *WirelessWirelessLansListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this wireless wireless lans list o k response has a 5xx status code
func (o *WirelessWirelessLansListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this wireless wireless lans list o k response a status code equal to that given
func (o *WirelessWirelessLansListOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the wireless wireless lans list o k response
func (o *WirelessWirelessLansListOK) Code() int {
	return 200
}

func (o *WirelessWirelessLansListOK) Error() string {
	return fmt.Sprintf("[GET /wireless/wireless-lans/][%d] wirelessWirelessLansListOK  %+v", 200, o.Payload)
}

func (o *WirelessWirelessLansListOK) String() string {
	return fmt.Sprintf("[GET /wireless/wireless-lans/][%d] wirelessWirelessLansListOK  %+v", 200, o.Payload)
}

func (o *WirelessWirelessLansListOK) GetPayload() *WirelessWirelessLansListOKBody {
	return o.Payload
}

func (o *WirelessWirelessLansListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(WirelessWirelessLansListOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWirelessWirelessLansListDefault creates a WirelessWirelessLansListDefault with default headers values
func NewWirelessWirelessLansListDefault(code int) *WirelessWirelessLansListDefault {
	return &WirelessWirelessLansListDefault{
		_statusCode: code,
	}
}

/*
WirelessWirelessLansListDefault describes a response with status code -1, with default header values.

WirelessWirelessLansListDefault wireless wireless lans list default
*/
type WirelessWirelessLansListDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this wireless wireless lans list default response has a 2xx status code
func (o *WirelessWirelessLansListDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this wireless wireless lans list default response has a 3xx status code
func (o *WirelessWirelessLansListDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this wireless wireless lans list default response has a 4xx status code
func (o *WirelessWirelessLansListDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this wireless wireless lans list default response has a 5xx status code
func (o *WirelessWirelessLansListDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this wireless wireless lans list default response a status code equal to that given
func (o *WirelessWirelessLansListDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the wireless wireless lans list default response
func (o *WirelessWirelessLansListDefault) Code() int {
	return o._statusCode
}

func (o *WirelessWirelessLansListDefault) Error() string {
	return fmt.Sprintf("[GET /wireless/wireless-lans/][%d] wireless_wireless-lans_list default  %+v", o._statusCode, o.Payload)
}

func (o *WirelessWirelessLansListDefault) String() string {
	return fmt.Sprintf("[GET /wireless/wireless-lans/][%d] wireless_wireless-lans_list default  %+v", o._statusCode, o.Payload)
}

func (o *WirelessWirelessLansListDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *WirelessWirelessLansListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
WirelessWirelessLansListOKBody wireless wireless lans list o k body
swagger:model WirelessWirelessLansListOKBody
*/
type WirelessWirelessLansListOKBody struct {

	// count
	// Required: true
	Count *int64 `json:"count"`

	// next
	// Format: uri
	Next *strfmt.URI `json:"next,omitempty"`

	// previous
	// Format: uri
	Previous *strfmt.URI `json:"previous,omitempty"`

	// results
	// Required: true
	Results []*models.WirelessLAN `json:"results"`
}

// Validate validates this wireless wireless lans list o k body
func (o *WirelessWirelessLansListOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCount(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateNext(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePrevious(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateResults(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *WirelessWirelessLansListOKBody) validateCount(formats strfmt.Registry) error {

	if err := validate.Required("wirelessWirelessLansListOK"+"."+"count", "body", o.Count); err != nil {
		return err
	}

	return nil
}

func (o *WirelessWirelessLansListOKBody) validateNext(formats strfmt.Registry) error {
	if swag.IsZero(o.Next) { // not required
		return nil
	}

	if err := validate.FormatOf("wirelessWirelessLansListOK"+"."+"next", "body", "uri", o.Next.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *WirelessWirelessLansListOKBody) validatePrevious(formats strfmt.Registry) error {
	if swag.IsZero(o.Previous) { // not required
		return nil
	}

	if err := validate.FormatOf("wirelessWirelessLansListOK"+"."+"previous", "body", "uri", o.Previous.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *WirelessWirelessLansListOKBody) validateResults(formats strfmt.Registry) error {

	if err := validate.Required("wirelessWirelessLansListOK"+"."+"results", "body", o.Results); err != nil {
		return err
	}

	for i := 0; i < len(o.Results); i++ {
		if swag.IsZero(o.Results[i]) { // not required
			continue
		}

		if o.Results[i] != nil {
			if err := o.Results[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("wirelessWirelessLansListOK" + "." + "results" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("wirelessWirelessLansListOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this wireless wireless lans list o k body based on the context it is used
func (o *WirelessWirelessLansListOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateResults(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *WirelessWirelessLansListOKBody) contextValidateResults(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Results); i++ {

		if o.Results[i] != nil {
			if err := o.Results[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("wirelessWirelessLansListOK" + "." + "results" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("wirelessWirelessLansListOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *WirelessWirelessLansListOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *WirelessWirelessLansListOKBody) UnmarshalBinary(b []byte) error {
	var res WirelessWirelessLansListOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
