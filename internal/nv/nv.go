package nv

type NV struct {
}

func NewNV() *NV {
	return &NV{}
}

// FilterNotes takes a query string and returns a slice of notes (a note is a file)
func (nv *NV) FilterNotes(q string) ([]string, error) {
	return []string{}, nil
}

// CreateNote creates a new note for the given query string q
func (nv *NV) CreateNote(q string) error {
	return nil
}
