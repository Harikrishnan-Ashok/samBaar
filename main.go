package main

import (
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/unit"
)

// Start func
func start(w *app.Window) error {
	for {
		evt := w.Event()
		switch typ := evt.(type) {
		case app.DestroyEvent:
			return typ.Err
		}
	}
}

func main() {
	go func() {
		w := new(app.Window)
		w.Option(app.Title("samBaar"))
		w.Option(app.MaxSize(unit.Dp(380), unit.Dp(1080)))
		if err := start(w); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}
