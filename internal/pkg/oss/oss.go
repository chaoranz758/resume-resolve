package oss

import "mime/multipart"

type OSS interface {
	Init() error
	Upload(multipart.File, string) error
}
