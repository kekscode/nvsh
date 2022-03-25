package main

import (
	"os"
	"testing"
)

func Test_getEditor(t *testing.T) {
	tests := []struct {
		name  string
		setup func()
		want  string
		env   map[string]string
	}{
		{"test for unset variables", func() {
			os.Unsetenv("VISUAL")
			os.Unsetenv("EDITOR")
		}, "", map[string]string{}},
		{"test for empty variables", func() {
			os.Setenv("VISUAL", "")
			os.Setenv("EDITOR", "")
		}, "", map[string]string{"VISUAL": "", "EDITOR": ""}},
		{"test for VISUAL variable (prefered)", func() {
			os.Setenv("VISUAL", "nvim")
		}, "nvim", map[string]string{"VISUAL": "nvim"}},
		{"test for EDITOR variable (backup)", func() {
			os.Setenv("EDITOR", "nvim")
		}, "nvim", map[string]string{"EDITOR": "nvim"}},
		{"test for EDITOR and VISUAL variable", func() {
			os.Setenv("VISUAL", "nvim")
			os.Setenv("EDITOR", "nvim")
		}, "nvim", map[string]string{"EDITOR": "nvim"}},
	}

	os.Unsetenv("VISUAL")
	os.Unsetenv("EDITOR")

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()
			if got := getEditor(); got != tt.want {
				t.Errorf("getEditor() = %v, want %v", got, tt.want)
			}
		})
	}
}
