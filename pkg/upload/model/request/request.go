package requestmodel

import (
	"errors"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/thoas/go-funk"
	"github.com/trungluongwww/goupload/internal/constant"
	"github.com/trungluongwww/goupload/internal/response"
)

type ClientPayload struct {
	BucketName string `json:"bucketName"`
	Resize     string `json:"resize"`
}

func (m ClientPayload) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Resize, validation.In(constant.ListResizeNameInterfaces...).Error(response.CommonInvalidPayload)),
		validation.Field(&m.BucketName, validation.Required.Error(response.CommonInvalidPayload),
			validation.In(constant.BucketShortInterfaces...).Error(response.CommonInvalidPayload)),
	)
}

type FileInfoPayload struct {
	OriginName string `json:"originName"`
	Path       string `json:"path"`
	Name       string `json:"name"`
	Size       int64  `json:"size"`
	Width      int    `json:"width"`
	Height     int    `json:"height"`
	Ext        string `json:"ext"`
	Type       string `json:"type"`
}

func (f FileInfoPayload) ValidateExtensionPhoto() error {
	if !funk.ContainsString(constant.ListExtensionPhotoValid, f.Ext) {
		return errors.New(response.CommonInvalidExtension)
	}
	return nil
}

func (f FileInfoPayload) ValidateSize(maxSize int64) error {
	if f.Size > maxSize {
		return errors.New(response.CommonInvalidSize)
	}
	return nil
}
