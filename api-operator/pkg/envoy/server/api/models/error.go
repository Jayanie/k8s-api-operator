// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Error Error object returned with 4XX HTTP status
//
// swagger:model Error
type Error struct {

	// Error code
	// Required: true
	Code *int64 `json:"code"`

	// A detail description about the error message.
	//
	Description string `json:"description,omitempty"`

	// If there are more than one error list them out.
	// For example, list out validation errors by each field.
	//
	Error []*ErrorListItem `json:"error"`

	// Error message.
	// Required: true
	Message *string `json:"message"`

	// Preferably an url with more details about the error.
	//
	MoreInfo string `json:"moreInfo,omitempty"`
}

// Validate validates this error
func (m *Error) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateError(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Error) validateCode(formats strfmt.Registry) error {

	if err := validate.Required("code", "body", m.Code); err != nil {
		return err
	}

	return nil
}

func (m *Error) validateError(formats strfmt.Registry) error {

	if swag.IsZero(m.Error) { // not required
		return nil
	}

	for i := 0; i < len(m.Error); i++ {
		if swag.IsZero(m.Error[i]) { // not required
			continue
		}

		if m.Error[i] != nil {
			if err := m.Error[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("error" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Error) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("message", "body", m.Message); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Error) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Error) UnmarshalBinary(b []byte) error {
	var res Error
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
