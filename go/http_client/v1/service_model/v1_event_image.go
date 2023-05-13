// Code generated by go-swagger; DO NOT EDIT.

package service_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1EventImage Image spec definition
//
// swagger:model v1EventImage
type V1EventImage struct {

	// Valid colorspace values are
	//   1 - grayscale
	//   2 - grayscale + alpha
	//   3 - RGB
	//   4 - RGBA
	//   5 - DIGITAL_YUV
	//   6 - BGRA
	Colorspace int32 `json:"colorspace,omitempty"`

	// Height of the image.
	Height int32 `json:"height,omitempty"`

	// Filepath
	Path string `json:"path,omitempty"`

	// Width of the image.
	Width int32 `json:"width,omitempty"`
}

// Validate validates this v1 event image
func (m *V1EventImage) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 event image based on context it is used
func (m *V1EventImage) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1EventImage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1EventImage) UnmarshalBinary(b []byte) error {
	var res V1EventImage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
