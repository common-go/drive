package drive

import (
	"context"
	. "github.com/common-go/search"
)

type FileService interface {
	Search(ctx context.Context, searchModel SearchModel) (*SearchResult, error)

	GetId(ctx context.Context, fullName string) (string, bool, int64, error)
	GetFullName(ctx context.Context, id string) (string, bool, int64, error)

	CreateDirectory(ctx context.Context, parentId string, name string) (string, error)
	Rename(ctx context.Context, id string, parentId string, newName string) error
	Delete(ctx context.Context, id string) error
	DeleteAll(ctx context.Context, id string) error

	Upload(ctx context.Context, parentId string, file *FileUpload) (string, error)
	Download(ctx context.Context, fileDownload *FileDownload) error

	CreateDirectoryWithParentPath(ctx context.Context, parentPath string, name string) (string, error)
	RenameToFullName(ctx context.Context, id string, newFullName string) error
	DeleteByFullName(ctx context.Context, fullName string) error
}
