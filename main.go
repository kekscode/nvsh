package main

import (
	"fmt"
	"log"

	"github.com/jroimartin/gocui"
)

var err error
var g *gocui.Gui

func main() {
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
	case key == gocui.KeySpace:
		v.EditWrite(' ')
		g.Update(func(gui *gocui.Gui) error {
			_, err := g.View("query")
			if err != nil {
				// handle error
			}
			return nil
		})
	case key == gocui.KeyBackspace || key == gocui.KeyBackspace2:
		v.EditDelete(true)
	}

}
