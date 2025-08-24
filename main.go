package main

import (
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/op"
	"gioui.org/unit"
	"github.com/Harikrishnan-Ashok/samBaar/ui/rootLayot"
)

// Start func
func start(w *app.Window) error {

	var ops op.Ops
	for {
		evt := w.Event()
		switch typ := evt.(type) {
		case app.DestroyEvent:
			return typ.Err
		case app.FrameEvent:
			gtx := app.NewContext(&ops, typ)
			rootlayot.RootLayout(gtx)
			typ.Frame(gtx.Ops)
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
