package nv

import (
	"github.com/sahilm/fuzzy"
)

type NV struct {
}

func NewNV() *NV {
	return &NV{}
}

// FuzzyFilterNotes takes a query string and returns a sub-slice of notes with fuzzy matched results
func (nv *NV) FuzzyFilterNotes(q string, n []string) ([]string, error) {
	matches := fuzzy.Find(q, n)
	s := []string{}
	for _, m := range matches {
		s = append(s, m.Str)
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
