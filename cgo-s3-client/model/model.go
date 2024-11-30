package model

import (
	"time"
)

type AuthRequestBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthResponseBody struct {
	Token string `json:"token"`
}

type ListFilesResponseBody struct {
	Files []File `json:"files"`
}

type FileAccess string

const (
	FileAccessPrivate FileAccess = "private"
	FileAccessPublic  FileAccess = "public"
)

type File struct {
	CreatedTS time.Time
	Filename  string
	MIME      string
	Access    FileAccess
	ID        string
	BucketID  int64
	SizeBytes int64
}
