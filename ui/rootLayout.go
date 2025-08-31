// Package ui contains the diffrent sections in the basic layout of the app
package ui

import (
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"github.com/Harikrishnan-Ashok/samBaar/state"
	"github.com/Harikrishnan-Ashok/samBaar/ui/sections"
)

func RootLayout(gtx layout.Context, th *material.Theme, store *state.UIState, w *app.Window) layout.Dimensions {
	margin := layout.Inset{
		Top:    unit.Dp(25),
		Left:   unit.Dp(25),
		Right:  unit.Dp(25),
		Bottom: unit.Dp(25),
	}

	return layout.Flex{
		Axis:    layout.Vertical,
		Spacing: layout.SpaceAround,
	}.Layout(gtx,
		layout.Flexed(3, func(gtx layout.Context) layout.Dimensions {
			return margin.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return sections.StatusControlSection(gtx, th, store)
			})
		}), layout.Flexed(4.5, func(gtx layout.Context) layout.Dimensions {
			return layout.Spacer{}.Layout(gtx)
		}),
		layout.Flexed(2, func(gtx layout.Context) layout.Dimensions {
			return margin.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return sections.PowerControlSection(gtx, th, store, w)
			})
		}),
	)
}
