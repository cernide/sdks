// Code generated by go-swagger; DO NOT EDIT.

package service_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1EventAudio Audio spec definition
//
// swagger:model v1EventAudio
type V1EventAudio struct {

	// Content type of the audio
	ContentType string `json:"content_type,omitempty"`

	// Length of the audio in frames (samples per channel).
	LengthFrames int32 `json:"length_frames,omitempty"`

	// Number of channels of audio.
	NumChannels int32 `json:"num_channels,omitempty"`

	// / Filepath
	Path string `json:"path,omitempty"`

	// Sample rate of the audio in Hz.
	SampleRate float32 `json:"sample_rate,omitempty"`
}

// Validate validates this v1 event audio
func (m *V1EventAudio) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 event audio based on context it is used
func (m *V1EventAudio) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1EventAudio) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1EventAudio) UnmarshalBinary(b []byte) error {
	var res V1EventAudio
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
