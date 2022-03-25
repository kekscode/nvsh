package nv

import (
	"errors"
	"fmt"
	"io/fs"
	"path/filepath"

	"github.com/sahilm/fuzzy"
)

type NV struct {
	Editor string
}

func New(editor string) *NV {
	return &NV{Editor: editor}
}

// FuzzyFilterNotes takes a query string and returns a sub-slice of notes with fuzzy matched results
func (nv *NV) FuzzyFilterNotes(q string, n []string) ([]string, error) {
	matches := fuzzy.Find(q, n)
	s := []string{}
	for _, m := range matches {
		s = append(s, m.Str)
	}
	if len(s) == 0 {
		return s, errors.New("No matches found")
	}
	return s, nil
}

// FuzzyFindNoteContent takes a query string and returns a sub-slice of fuzzy matches in content
func (nv *NV) FuzzyFindNoteContent(q string, n []string) ([]string, error) {
	return []string{}, nil
}

// CreateNote creates a new note for the given query string q
func (nv *NV) CreateNote(q string) error {
	return nil
}

func (nv *NV) GetFiles(dir string) ([]string, error) {

	var files []string

	filepath.WalkDir(dir, func(path string, info fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		fmt.Println(path, info.Name())
		files = append(files, path)
		return nil
	})

	return files, nil
}
