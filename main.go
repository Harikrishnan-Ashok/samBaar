package main

import (
	_ "embed"
	"log"
	"os"
	"os/exec"
	"strings"

	"gioui.org/app"
	"gioui.org/font"
	"gioui.org/font/opentype"
	"gioui.org/io/event"
	"gioui.org/io/key"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"github.com/Harikrishnan-Ashok/samBaar/state"
	"github.com/Harikrishnan-Ashok/samBaar/ui"
)

//go:embed Assets/fonts/DaddyTimeMonoNerdFontMono-Regular.ttf
var nerdFontData []byte
var keyInputTag struct{}

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

		//decleartion stuff
		store := &state.UIState{}
		w := new(app.Window)
		w.Option(app.Title("samBaar"))
		w.Option(app.MaxSize(unit.Dp(380), unit.Dp(620)))
		w.Option(app.MinSize(unit.Dp(380), unit.Dp(620)))

		go func() {
			// Battery status
			out, err := exec.Command("acpi", "-b").Output()
			if err != nil {
				store.BatteryStatus = "Err querying battery info"
			} else {
				store.BatteryStatus = strings.TrimSpace(string(out))
			}

			// Time status
			out, err = exec.Command("date", "+%H:%M").Output()
			if err != nil {
				store.TimeStatus = "Err querying time info"
			} else {
				store.TimeStatus = strings.TrimSpace(string(out))
			}

			// WiFi status
			out, err = exec.Command("nmcli", "radio", "wifi").Output()
			if err != nil {
				store.WifiStatus = "Err querying WiFi"
			} else {
				wifiStatus := strings.TrimSpace(string(out))
				switch wifiStatus {
				case "enabled":
					store.WifiStatus = "WiFi Up"
				case "disabled":
					store.WifiStatus = " Down"
				default:
					store.WifiStatus = "Unknown: " + wifiStatus
				}
			}

			// Bluetooth status
			out, err = exec.Command("bluetoothctl", "show").Output()
			if err != nil {
				store.BluetoothStatus = "Err querying Bluetooth"
			} else {
				lines := strings.Split(string(out), "\n")
				status := "Unknown"
				for _, line := range lines {
					line = strings.TrimSpace(line)
					if strings.HasPrefix(line, "Powered:") {
						parts := strings.Fields(line)
						if len(parts) == 2 {
							if strings.ToLower(parts[1]) == "yes" {
								status = "Bl| Up "
							} else if strings.ToLower(parts[1]) == "no" {
								status = "Bl|Down"
							}
						}
						break
					}
				}
				store.BluetoothStatus = status
			}
			out, err = exec.Command("date", "+ %d - %b - %Y  %A").Output()
			if err != nil {
				store.DateStatus = "Err getting date"
			} else {
				store.DateStatus = string(out)
			}

			// Trigger redraw
			w.Invalidate()
		}()
		// Start the app
		if err := start(w, th, store); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

func start(w *app.Window, th *material.Theme, store *state.UIState) error {
	var ops op.Ops
	for {
		e := w.Event()
		switch evt := e.(type) {
		case app.DestroyEvent:
			return evt.Err
		case app.FrameEvent:
			gtx := app.NewContext(&ops, evt)

			// Register interest in key events
			event.Op(gtx.Ops, keyInputTag)

			// Handle key input
			for {
				ev, ok := gtx.Event(key.Filter{Name: key.NameEscape})
				if !ok {
					break
				}
				if ke, ok := ev.(key.Event); ok && ke.State == key.Press {
					os.Exit(0)
				}
			}

			// Draw UI
			ui.RootLayout(gtx, th, store, w)
			evt.Frame(gtx.Ops)
		}
	}
}
