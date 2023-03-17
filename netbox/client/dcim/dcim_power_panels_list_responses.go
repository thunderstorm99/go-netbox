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

// DcimPowerPanelsListReader is a Reader for the DcimPowerPanelsList structure.
type DcimPowerPanelsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerPanelsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimPowerPanelsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimPowerPanelsListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimPowerPanelsListOK creates a DcimPowerPanelsListOK with default headers values
func NewDcimPowerPanelsListOK() *DcimPowerPanelsListOK {
	return &DcimPowerPanelsListOK{}
}

/*
DcimPowerPanelsListOK describes a response with status code 200, with default header values.

DcimPowerPanelsListOK dcim power panels list o k
*/
type DcimPowerPanelsListOK struct {
	Payload *DcimPowerPanelsListOKBody
}

// IsSuccess returns true when this dcim power panels list o k response has a 2xx status code
func (o *DcimPowerPanelsListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim power panels list o k response has a 3xx status code
func (o *DcimPowerPanelsListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim power panels list o k response has a 4xx status code
func (o *DcimPowerPanelsListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim power panels list o k response has a 5xx status code
func (o *DcimPowerPanelsListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim power panels list o k response a status code equal to that given
func (o *DcimPowerPanelsListOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the dcim power panels list o k response
func (o *DcimPowerPanelsListOK) Code() int {
	return 200
}

func (o *DcimPowerPanelsListOK) Error() string {
	return fmt.Sprintf("[GET /dcim/power-panels/][%d] dcimPowerPanelsListOK  %+v", 200, o.Payload)
}

func (o *DcimPowerPanelsListOK) String() string {
	return fmt.Sprintf("[GET /dcim/power-panels/][%d] dcimPowerPanelsListOK  %+v", 200, o.Payload)
}

func (o *DcimPowerPanelsListOK) GetPayload() *DcimPowerPanelsListOKBody {
	return o.Payload
}

func (o *DcimPowerPanelsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(DcimPowerPanelsListOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimPowerPanelsListDefault creates a DcimPowerPanelsListDefault with default headers values
func NewDcimPowerPanelsListDefault(code int) *DcimPowerPanelsListDefault {
	return &DcimPowerPanelsListDefault{
		_statusCode: code,
	}
}

/*
DcimPowerPanelsListDefault describes a response with status code -1, with default header values.

DcimPowerPanelsListDefault dcim power panels list default
*/
type DcimPowerPanelsListDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this dcim power panels list default response has a 2xx status code
func (o *DcimPowerPanelsListDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim power panels list default response has a 3xx status code
func (o *DcimPowerPanelsListDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim power panels list default response has a 4xx status code
func (o *DcimPowerPanelsListDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim power panels list default response has a 5xx status code
func (o *DcimPowerPanelsListDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim power panels list default response a status code equal to that given
func (o *DcimPowerPanelsListDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the dcim power panels list default response
func (o *DcimPowerPanelsListDefault) Code() int {
	return o._statusCode
}

func (o *DcimPowerPanelsListDefault) Error() string {
	return fmt.Sprintf("[GET /dcim/power-panels/][%d] dcim_power-panels_list default  %+v", o._statusCode, o.Payload)
}

func (o *DcimPowerPanelsListDefault) String() string {
	return fmt.Sprintf("[GET /dcim/power-panels/][%d] dcim_power-panels_list default  %+v", o._statusCode, o.Payload)
}

func (o *DcimPowerPanelsListDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimPowerPanelsListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
DcimPowerPanelsListOKBody dcim power panels list o k body
swagger:model DcimPowerPanelsListOKBody
*/
type DcimPowerPanelsListOKBody struct {

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
	Results []*models.PowerPanel `json:"results"`
}

// Validate validates this dcim power panels list o k body
func (o *DcimPowerPanelsListOKBody) Validate(formats strfmt.Registry) error {
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

func (o *DcimPowerPanelsListOKBody) validateCount(formats strfmt.Registry) error {

	if err := validate.Required("dcimPowerPanelsListOK"+"."+"count", "body", o.Count); err != nil {
		return err
	}

	return nil
}

func (o *DcimPowerPanelsListOKBody) validateNext(formats strfmt.Registry) error {
	if swag.IsZero(o.Next) { // not required
		return nil
	}

	if err := validate.FormatOf("dcimPowerPanelsListOK"+"."+"next", "body", "uri", o.Next.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *DcimPowerPanelsListOKBody) validatePrevious(formats strfmt.Registry) error {
	if swag.IsZero(o.Previous) { // not required
		return nil
	}

	if err := validate.FormatOf("dcimPowerPanelsListOK"+"."+"previous", "body", "uri", o.Previous.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *DcimPowerPanelsListOKBody) validateResults(formats strfmt.Registry) error {

	if err := validate.Required("dcimPowerPanelsListOK"+"."+"results", "body", o.Results); err != nil {
		return err
	}

	for i := 0; i < len(o.Results); i++ {
		if swag.IsZero(o.Results[i]) { // not required
			continue
		}

		if o.Results[i] != nil {
			if err := o.Results[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dcimPowerPanelsListOK" + "." + "results" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dcimPowerPanelsListOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this dcim power panels list o k body based on the context it is used
func (o *DcimPowerPanelsListOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateResults(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DcimPowerPanelsListOKBody) contextValidateResults(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Results); i++ {

		if o.Results[i] != nil {
			if err := o.Results[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dcimPowerPanelsListOK" + "." + "results" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dcimPowerPanelsListOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *DcimPowerPanelsListOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DcimPowerPanelsListOKBody) UnmarshalBinary(b []byte) error {
	var res DcimPowerPanelsListOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
