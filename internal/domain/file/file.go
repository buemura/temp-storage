package file

import "github.com/google/uuid"

type File struct {
	ID       string `json:"id"`
	FileName string `json:"fileName"`
	FileUrl  string `json:"fileUrl"`
}

func NewFile(filename, fileUrl string) *File {
	return &File{
		ID:       uuid.NewString(),
		FileName: filename,
		FileUrl:  fileUrl,
	}
}
