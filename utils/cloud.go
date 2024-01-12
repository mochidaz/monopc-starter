package utils

import "mime/multipart"

type CloudStorage interface {
	UploadFile(file *multipart.FileHeader, folder string) (string, error)
	UploadSavedFile(filePath string, folder string) (string, error)
	DeleteFile(filePath string) error
}
