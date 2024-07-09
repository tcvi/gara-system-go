package s3

import (
	"context"
	"garasystem/internal/core/domain"
	"garasystem/internal/core/myerror"
	"garasystem/pkg/config"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gabriel-vasile/mimetype"
	uuid "github.com/satori/go.uuid"
	"mime/multipart"
)

func (s *Service) UploadFile(file *multipart.FileHeader) (*domain.FileUpload, error) {
	reader, err := file.Open()
	if err != nil {
		return nil, myerror.ErrFileOpenFile(err)
	}
	defer reader.Close()

	mime, err := mimetype.DetectReader(reader)
	if err != nil {
		return nil, myerror.ErrFileDetectMimeFile(err)
	}

	// Reset point reader
	_, err = reader.Seek(0, 0)
	if err != nil {
		return nil, myerror.ErrFileSeek(err)
	}

	var (
		key         = uuid.NewV4().String() + mime.Extension()
		contentType = mime.String()
		cfg         = config.GetConfig()
	)

	_, err = s.client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket:        aws.String(cfg.AWS.BucketName),
		Key:           aws.String(key),
		ContentType:   aws.String(contentType),
		Body:          reader,
		ContentLength: aws.Int64(file.Size),
	})
	if err != nil {
		return nil, myerror.ErrFileUploadS3(err)
	}

	return &domain.FileUpload{
		Key:         key,
		ContentType: contentType,
	}, nil
}
