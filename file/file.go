package file

import (
	"PnoT/util"
	"time"
)

type File struct {
	Path        string    `json:"path"`
	Content     string    `json:"content"`
	Edit_date   time.Time `json:"edit_date"`
	Create_date time.Time `json:"create_date"`
	Public      bool      `json:"public"`
	Autor       string    `json:"autor"`
	History     []FileHistory
}
type FileHistory struct {
	Content   string    `json:"content"`
	Public    bool      `json:"publish"`
	Edit_date time.Time `json:"edit_date"`
}

func (f *File) ToHistory() FileHistory {
	return FileHistory{
		Content:   f.Content,
		Edit_date: f.Edit_date,
		Public:    f.Public,
	}
}
func (f *File) SaveHistory() {
	f.History = append(f.History, f.ToHistory())
}

func (f *File) LoadFile(b []byte) error {
	return util.BytetoStruct(b, f)
}

func (f *File) ToByte() ([]byte, error) {
	return util.StructtoByte(f)
}

func CreateFile(autor, path, content string, public bool) *File {
	return &File{
		Path:        path,
		Content:     content,
		Edit_date:   time.Now(),
		Create_date: time.Now(),
		Public:      public,
		Autor:       autor,
		History:     []FileHistory{},
	}
}
