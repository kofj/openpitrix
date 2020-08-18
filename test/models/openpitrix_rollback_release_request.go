// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixRollbackReleaseRequest openpitrix rollback release request
// swagger:model openpitrixRollbackReleaseRequest
type OpenpitrixRollbackReleaseRequest struct {

	// release name
	ReleaseName string `json:"release_name,omitempty"`

	// runtime id
	RuntimeID string `json:"runtime_id,omitempty"`

	// version
	Version int32 `json:"version,omitempty"`
}

// Validate validates this openpitrix rollback release request
func (m *OpenpitrixRollbackReleaseRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixRollbackReleaseRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixRollbackReleaseRequest) UnmarshalBinary(b []byte) error {
	var res OpenpitrixRollbackReleaseRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}