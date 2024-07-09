package file

import (
	"garasystem/internal/core/domain"
	"garasystem/internal/core/myerror"
	"mime/multipart"
	"time"
)

func (s *Service) Upload(files []*multipart.FileHeader) ([]domain.FileUpload, error) {
	filesUpload := make([]domain.FileUpload, len(files))

	for i, file := range files {
		fileRes, err := s.s3.UploadFile(file)
		if err != nil {
			return nil, err.(myerror.MyError)
		}

		url, err := s.s3.PresignObject(fileRes.Key, 15*time.Minute)
		if err != nil {
			return nil, err.(myerror.MyError)
		}

		fileRes.Url = url
		filesUpload[i] = *fileRes
	}

	return filesUpload, nil
}
