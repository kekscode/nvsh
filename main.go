package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/jroimartin/gocui"
	"github.com/kekscode/nvsh/internal/nv"
	"github.com/sahilm/fuzzy"
)

var err error
var g *gocui.Gui

func main() {

	nv := nv.New(getEditor())

	if nv.Editor == "" {
		log.Fatalf("Neither VISUAL nor EDITOR environment variables have been set or both are set but empty")
	}

	fmt.Printf("Editor found: %v\n", nv.Editor)

	g, err = gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.Cursor = true
	g.Mouse = false
	g.Highlight = true

	g.SetManagerFunc(layout)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	fmt.Print("")
	if v, err := g.SetView("query", 0, 0, maxX-1, 2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Wrap = true
		v.Editable = true
		v.Frame = true
		v.Title = "Search note"
		if _, err := g.SetCurrentView("query"); err != nil {
			return err
		}
		v.Editor = gocui.EditorFunc(finder)
	}
	if v, err := g.SetView("results", 0, 2, maxX-1, maxY+(maxY/2)); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Editable = false
		v.Wrap = true
		v.Frame = true
		v.Title = "Found notes"
	}

	if v, err := g.SetView("content", 0, maxY/2, maxX-1, maxY+(maxY/2)); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Editable = false
		v.Wrap = true
		v.Frame = true
		v.Title = "Found in content"
	}

	return nil
}

func finder(v *gocui.View, key gocui.Key, ch rune, mod gocui.Modifier) {
	// TODO: https://github.com/sahilm/fuzzy/blob/master/_example/main.go
	switch {
	case ch != 0 && mod == 0:
		v.EditWrite(ch)
		g.Update(func(gui *gocui.Gui) error {
			results, err := g.View("results")
			if err != nil {
				// handle error
			}
			results.Clear()
			t := time.Now()
			matches := fuzzy.Find(strings.TrimSpace(v.ViewBuffer()), []string{"abc", "def", "ghi"})
			elapsed := time.Since(t)
			fmt.Fprintf(results, "found %v matches in %v\n", len(matches), elapsed)
			for _, match := range matches {
				for i := 0; i < len(match.Str); i++ {
					if contains(i, match.MatchedIndexes) {
						fmt.Fprintf(results, fmt.Sprintf("\033[1m%s\033[0m", string(match.Str[i])))
					} else {
						fmt.Fprintf(results, string(match.Str[i]))
					}

				}
				fmt.Fprintln(results, "")
			}
			return nil
		})
	case key == gocui.KeySpace:
		v.EditWrite(' ')
	case key == gocui.KeyBackspace || key == gocui.KeyBackspace2:
		v.EditDelete(true)
		g.Update(func(gui *gocui.Gui) error {
			results, err := g.View("results")
			if err != nil {
				// handle error
			}
			results.Clear()
			t := time.Now()
			matches := fuzzy.Find(strings.TrimSpace(v.ViewBuffer()), []string{"abc", "def", "ghi"})
			elapsed := time.Since(t)
			fmt.Fprintf(results, "found %v matches in %v\n", len(matches), elapsed)
			for _, match := range matches {
				for i := 0; i < len(match.Str); i++ {
					if contains(i, match.MatchedIndexes) {
						fmt.Fprintf(results, fmt.Sprintf("\033[1m%s\033[0m", string(match.Str[i])))
					} else {
						fmt.Fprintf(results, string(match.Str[i]))
					}
				}
				fmt.Fprintln(results, "")
			}
			return nil
		})
	case key == gocui.KeyDelete:
		v.EditDelete(false)
		g.Update(func(gui *gocui.Gui) error {
			results, err := g.View("results")
			if err != nil {
				// handle error
			}
			results.Clear()
			t := time.Now()
			matches := fuzzy.Find(strings.TrimSpace(v.ViewBuffer()), []string{"abc", "def", "ghi"})
			elapsed := time.Since(t)
			fmt.Fprintf(results, "found %v matches in %v\n", len(matches), elapsed)
			for _, match := range matches {
				for i := 0; i < len(match.Str); i++ {
					if contains(i, match.MatchedIndexes) {
						fmt.Fprintf(results, fmt.Sprintf("\033[1m%s\033[0m", string(match.Str[i])))
					} else {
						fmt.Fprintf(results, string(match.Str[i]))
					}
				}
				fmt.Fprintln(results, "")
			}
			return nil
		})
	case key == gocui.KeyInsert:
		v.Overwrite = !v.Overwrite
	}
}

func contains(needle int, haystack []int) bool {
	for _, i := range haystack {
		if needle == i {
			return true
		}
	}
	return false
}

func getEditor() string {
	editor := os.Getenv("VISUAL")
	if editor == "" {
		editor = os.Getenv("EDITOR")
	}

	return editor
}
