package main

import (
	"errors"
	"log"

	"github.com/awesome-gocui/gocui"
)

func main() {
	g, err := gocui.NewGui(gocui.OutputNormal, true)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.Highlight = true

	if err := keybindings(g, ""); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && !errors.Is(err, gocui.ErrQuit) {
		log.Panicln(err)
	}
}

func keybindings(g *gocui.Gui, view string) error {
	type keybindFunction func(*gocui.Gui, *gocui.View) error
	keybindingMap := map[gocui.Key]keybindFunction{
		gocui.KeyCtrlC: func(g *gocui.Gui, v *gocui.View) error {
			return gocui.ErrQuit
		},
	}

	for key, fn := range keybindingMap {
		if err := g.SetKeybinding(view, key, gocui.ModNone, fn); err != nil {
			return err
		}
	}

	return nil
}
