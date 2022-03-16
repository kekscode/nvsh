package main

import (
	"os"
	"testing"
)

func Test_getEditor(t *testing.T) {
	tests := []struct {
		name string
		want string
		env  map[string]string
	}{
		{"test for unset variables", "nvim", map[string]string{}},
		{"test for empty variables", "nvim", map[string]string{"VISUAL": "", "EDITOR": ""}},
		{"test for VISUAL variable (prefered)", "nvim", map[string]string{"VISUAL": "nvim"}},
		{"test for EDITOR variable (backup)", "nvim", map[string]string{"EDITOR": "nvim"}},
		{"test for EDITOR and VISUAL variable", "nvim", map[string]string{"EDITOR": "nvim"}},
	}

	os.Unsetenv("EDITOR")
	os.Unsetenv("VISUAL")

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getEditor(); got != tt.want {
				t.Errorf("getEditor() = %v, want %v", got, tt.want)
			}
		})
	}
}
