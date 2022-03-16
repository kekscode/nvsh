package nv

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name   string
		editor string
		want   *NV
	}{
		{"test create new NV object", "nvim", &NV{"nvim"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.editor); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNV_FuzzyFilterNotes(t *testing.T) {
	type args struct {
		q string
		n []string
	}
	tests := []struct {
		name    string
		nv      *NV
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "test exact match",
			nv:   &NV{"nvim"},
			args: args{
				q: "abc",
				n: []string{"cde", "abc", "efg"},
			},
			want:    []string{"abc"},
			wantErr: false,
		},
		{
			name: "test near match",
			nv:   &NV{"nvim"},
			args: args{
				q: "efa",
				n: []string{"cde", "abc", "efg"},
			},
			want:    []string{"efg"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.nv.FuzzyFilterNotes(tt.args.q, tt.args.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("NV.FuzzyFilterNotes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NV.FuzzyFilterNotes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNV_CreateNote(t *testing.T) {
	type args struct {
		q string
	}
	tests := []struct {
		name    string
		nv      *NV
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.nv.CreateNote(tt.args.q); (err != nil) != tt.wantErr {
				t.Errorf("NV.CreateNote() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNV_FuzzyFindNoteContent(t *testing.T) {
	type args struct {
		q string
		n []string
	}
	tests := []struct {
		name    string
		nv      *NV
		args    args
		want    []string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.nv.FuzzyFindNoteContent(tt.args.q, tt.args.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("NV.FuzzyFindNoteContent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NV.FuzzyFindNoteContent() = %v, want %v", got, tt.want)
			}
		})
	}
}
