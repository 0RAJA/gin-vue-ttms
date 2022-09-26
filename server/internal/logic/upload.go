package logic

import (
	"errors"
	"mime/multipart"

	"ttms/internal/global"
	"ttms/internal/pkg/app/errcode"
	"ttms/internal/upload"
)

type up struct {
}

func (up) Upload(file *multipart.FileHeader) (string, errcode.Err) {
	url, _, err := upload.NewOSS().UploadFile(file)
	if err != nil {
		if errors.Is(err, upload.ErrFileOpen) {
			return "", errcode.ErrFileOpen
		}
		global.Logger.Error(err.Error())
		return "", errcode.ErrServer
	}
	return url, nil
}
