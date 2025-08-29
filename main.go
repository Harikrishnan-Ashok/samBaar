package main

import (
	_ "embed"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/font"
	"gioui.org/font/opentype"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"github.com/Harikrishnan-Ashok/samBaar/ui"
)

//go:embed Assets/fonts/DaddyTimeMonoNerdFontMono-Regular.ttf
var nerdFontData []byte

func main() {
	go func() {
		// Load and parse the font
		nerdFontFace, err := opentype.Parse(nerdFontData)
		if err != nil {
			log.Fatalf("Failed to parse font: %v", err)
		}

		// Create a shaper with the font
		shaper := text.NewShaper(text.WithCollection([]text.FontFace{
			{
				Font: font.Font{Typeface: "NerdFont"},
				Face: nerdFontFace,
			},
		}))

		// Create custom theme with shaper
		th := material.NewTheme()
		th.Shaper = shaper

		// Start the app
		w := new(app.Window)
		w.Option(app.Title("samBaar"))
		w.Option(app.MaxSize(unit.Dp(380), unit.Dp(1080)))
		if err := start(w, th); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

func start(w *app.Window, th *material.Theme) error {
	var ops op.Ops
	for {
		e := w.Event()
		switch evt := e.(type) {
		case app.DestroyEvent:
			return evt.Err
		case app.FrameEvent:
			gtx := app.NewContext(&ops, evt)
			ui.RootLayout(gtx, th)
			evt.Frame(gtx.Ops)
		}
	}
}
