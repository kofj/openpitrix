// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixSyncRepoResponse openpitrix sync repo response
// swagger:model openpitrixSyncRepoResponse
type OpenpitrixSyncRepoResponse struct {

	// synchronized ok or not
	Failed bool `json:"failed,omitempty"`

	// result
	Result string `json:"result,omitempty"`
}

// Validate validates this openpitrix sync repo response
func (m *OpenpitrixSyncRepoResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixSyncRepoResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixSyncRepoResponse) UnmarshalBinary(b []byte) error {
	var res OpenpitrixSyncRepoResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
