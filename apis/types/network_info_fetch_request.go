// Code generated by go-swagger; DO NOT EDIT.

package types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// NetworkInfoFetchRequest The request is to fetch p2p network info from supernode.
// swagger:model NetworkInfoFetchRequest
type NetworkInfoFetchRequest struct {

	// The urls is to filter the peer node, the url should be match with taskURL in TaskInfo.
	//
	Urls []string `json:"urls"`
}

// Validate validates this network info fetch request
func (m *NetworkInfoFetchRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NetworkInfoFetchRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetworkInfoFetchRequest) UnmarshalBinary(b []byte) error {
	var res NetworkInfoFetchRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}