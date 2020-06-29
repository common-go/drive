package drive

import "io"

type FileUpload struct {
	Name     string    `json:"name"`
	MimeType string    `json:"mimeType"`
	Content  io.Reader `json:"content"`
	Size     int64     `json:"size"`
}
