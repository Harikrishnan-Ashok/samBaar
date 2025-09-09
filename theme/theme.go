// Package theme for light and dark theme
package theme

import "image/color"

type Colors struct {
	EnabledColor  color.NRGBA
	DisabledColor color.NRGBA
	WarningColor  color.NRGBA
	TrasnparentBg color.NRGBA
}

type DefaultTheme struct {
	Colors Colors
}

// DarkTheme ...
var DarkTheme = DefaultTheme{
	Colors: Colors{
		EnabledColor:  color.NRGBA{R: 0, G: 200, B: 0, A: 255},
		DisabledColor: color.NRGBA{R: 200, G: 0, B: 0, A: 255},
		WarningColor:  color.NRGBA{R: 255, G: 150, B: 0, A: 200},
		TrasnparentBg: color.NRGBA{R: 0, G: 0, B: 0, A: 0},
	},
}
var LightTheme = DefaultTheme{
	Colors: Colors{
		EnabledColor:  color.NRGBA{R: 0, G: 200, B: 0, A: 255},
		DisabledColor: color.NRGBA{R: 200, G: 0, B: 0, A: 255},
		WarningColor:  color.NRGBA{R: 255, G: 150, B: 0, A: 200},
		TrasnparentBg: color.NRGBA{R: 0, G: 0, B: 0, A: 0},
	},
}
