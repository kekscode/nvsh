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

func TestNV_FilterNotes(t *testing.T) {
	type args struct {
		q string
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
			got, err := tt.nv.FilterNotes(tt.args.q)
			if (err != nil) != tt.wantErr {
				t.Errorf("NV.FilterNotes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NV.FilterNotes() = %v, want %v", got, tt.want)
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
