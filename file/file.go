package file

import (
	"PnoT/util"
	"crypto/sha512"
	"fmt"
	"time"
)

type File struct {
	Path        string    `json:"path"`
	Content     string    `json:"content"`
	Edit_date   time.Time `json:"edit_date"`
	Create_date time.Time `json:"create_date"`
	Public      bool      `json:"public"`
	Autor       string    `json:"autor"`
	Hash        string    `json:"hash"`
	History     []FileHistory
}
type FileHistory struct {
	Content   string    `json:"content"`
	Hash      string    `json:"hash"`
	Public    bool      `json:"publish"`
	Edit_date time.Time `json:"edit_date"`
}

func (f *File) ToHistory() FileHistory {
	return FileHistory{
		Content:   f.Content,
		Edit_date: f.Edit_date,
		Public:    f.Public,
		Hash:      f.Hash,
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

func (f *File) GetHash() string {
	a := ""
	if f.Public {
		a = "true"
	} else {
		a = "false"
	}
	return fmt.Sprintf("%x", sha512.Sum512([]byte(f.Content+f.Path+f.Edit_date.String()+a)))
}

func CreateFile(autor, path, content string, public bool) *File {
	now := time.Now()
	f := File{
		Path:        path,
		Content:     content,
		Edit_date:   now,
		Create_date: now,
		Public:      public,
		Autor:       autor,
		History:     []FileHistory{},
	}
	f.Hash = f.GetHash()
	return &f
}
