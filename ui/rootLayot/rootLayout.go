// Package rootlayot contains the diffrent sections in the basic layout of the app
package rootlayot

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget/material"
)

func RootLayout(gtx layout.Context) layout.Dimensions {
	th := material.NewTheme()
	return layout.Flex{
		Axis:    layout.Vertical,
		Spacing: layout.SpaceStart,
	}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			margin := layout.Inset{
				Top:    unit.Dp(25),
				Left:   unit.Dp(25),
				Right:  unit.Dp(25),
				Bottom: unit.Dp(25),
			}
			return margin.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return PowerControlSection(gtx, th)
			})
		}),
	)
}
