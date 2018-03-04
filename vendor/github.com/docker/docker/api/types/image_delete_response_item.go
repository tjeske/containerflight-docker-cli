package types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

// ImageDeleteResponseItem image delete response item
// swagger:model ImageDeleteResponseItem
type ImageDeleteResponseItem struct {

	// The image ID of an image that was deleted
	Deleted string `json:"Deleted,omitempty"`

	// The image ID of an image that was untagged
	Untagged string `json:"Untagged,omitempty"`
}
