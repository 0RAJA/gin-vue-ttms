package request

import (
	"mime/multipart"

	"ttms/internal/global"
	"ttms/internal/pkg/app/errcode"
)

type Upload struct {
	File *multipart.FileHeader `form:"file" binding:"required"`
}

func (r *Upload) Judge() errcode.Err {
	switch {
	case r.File.Size > global.Settings.Rule.MaxFileSize:
		return errcode.ErrFileOutSize
	default:
		return nil
	}
}
