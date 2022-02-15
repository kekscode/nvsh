package nv

import (
	"reflect"
	"testing"
)

func TestNewNV(t *testing.T) {
	tests := []struct {
		name string
		want *NV
	}{
		{"New NV object", NewNV()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewNV(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNV() = %v, want %v", got, tt.want)
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
			name: "Get exact match",
			nv:   &NV{},
			args: args{
				q: "abc",
				n: []string{"cde", "abc", "efg"},
			},
			want:    []string{"abc"},
			wantErr: false,
		},
		{
			name: "No match",
			nv:   &NV{},
			args: args{
				q: "xxx",
				n: []string{"cde", "abc", "efg"},
			},
			want:    []string{""},
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
