package service

import (
	"context"
	requestmodel "github.com/trungluongwww/goupload/pkg/upload/model/request"
	responsemodel "github.com/trungluongwww/goupload/pkg/upload/model/response"
)

type PhotoInterface interface {
	Upload(c context.Context, payloads []requestmodel.FileInfoPayload) ([]responsemodel.PhotoResponse, error)
}

type photoImpl struct {
}

func Photo() PhotoInterface {
	return photoImpl{}
}

func (photoImpl) Upload(c context.Context, payloads []requestmodel.FileInfoPayload) (result []responsemodel.PhotoResponse, err error) {

	return
}
