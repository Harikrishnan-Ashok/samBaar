package sections

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"github.com/Harikrishnan-Ashok/samBaar/state"
)

func UtilsSection(gtx layout.Context, th *material.Theme, store *state.UIState) layout.Dimensions {
	return layout.Flex{
		Axis:    layout.Horizontal,
		Spacing: layout.SpaceBetween,
	}.Layout(gtx,
		layout.Flexed(0.6, func(gtx layout.Context) layout.Dimensions {
			return layout.Inset{
				Right: unit.Dp(5),
			}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return layout.Stack{}.Layout(gtx,
					layout.Stacked(func(gtx layout.Context) layout.Dimensions {
						btn := material.Button(th, &store.WallpaperStatus, "Wallpaper temporarly disabled , preloading is needed")
						btn.TextSize = unit.Sp(10)
						gtx.Constraints.Min = gtx.Constraints.Max
						return btn.Layout(gtx)
					}),
					layout.Stacked(func(gtx layout.Context) layout.Dimensions {
						return layout.Inset{Top: unit.Dp(45)}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
							return layout.Flex{
								Axis:    layout.Horizontal,
								Spacing: layout.SpaceBetween,
							}.Layout(gtx,
								layout.Rigid(func(gtx layout.Context) layout.Dimensions {
									btn := material.Button(th, &store.PrevWallpaper, "<")
									transparentBG := btn.Background
									transparentBG.A = 200
									btn.Background = transparentBG
									return btn.Layout(gtx)
								}),
								layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
									return layout.Spacer{}.Layout(gtx)
								}),
								layout.Rigid(func(gtx layout.Context) layout.Dimensions {
									btn := material.Button(th, &store.NextWallpaper, ">")
									transparentBG := btn.Background
									transparentBG.A = 200
									btn.Background = transparentBG
									return btn.Layout(gtx)
								}),
							)
						})
					}),
				)
			})
		}),
		layout.Flexed(0.4, func(gtx layout.Context) layout.Dimensions {
			return layout.Inset{
				Left: unit.Dp(5),
			}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return layout.Flex{
					Axis:    layout.Vertical,
					Spacing: layout.SpaceBetween,
				}.Layout(gtx,
					layout.Flexed(0.5, func(gtx layout.Context) layout.Dimensions {
						return layout.Inset{
							Bottom: unit.Dp(10),
						}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
							btn := material.Button(th, &store.TimerStatus.Button, "00:00")
							btn.TextSize = unit.Sp(25)
							return btn.Layout(gtx)
						})
					}),
					layout.Flexed(0.3, func(gtx layout.Context) layout.Dimensions {
						return material.Button(th, &store.TimerStartButton.Button, "start Timer").Layout(gtx)
					}),
				)
			})
		}),
	)
}
