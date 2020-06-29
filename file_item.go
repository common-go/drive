package drive

import "time"

// FileItem for OneDrive item object
type FileItem struct {
	Id           string    `json:"id"`
	Name         string    `json:"name"`
	Size         int       `json:"size"`
	ModifiedTime time.Time `json:"lastModifiedDateTime"`
	Owners       CreatedBy `json:"createdBy"`
}

type CreatedBy struct {
	User DriveUser `json:"user"`
	//Application DriveUser `json:"application"`
}

type DriveUser struct {
	DisplayName string `json:"displayName"`
	ID          string `json:"id"`
}
