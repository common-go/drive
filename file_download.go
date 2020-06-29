package drive

type FileDownload struct {
	FileId        string `json:"id"`
	FilePath      string `json:"path"`
	LocalFilePath string
}
