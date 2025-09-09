package sections

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"github.com/Harikrishnan-Ashok/samBaar/state"
)

func SystemControlSection(gtx layout.Context, th *material.Theme, store *state.UIState) layout.Dimensions {
	return layout.Flex{
		Axis:    layout.Vertical,
		Spacing: layout.SpaceBetween,
	}.Layout(gtx,
		layout.Flexed(0.5, func(gtx layout.Context) layout.Dimensions {
			return layout.Inset{
				Bottom: unit.Dp(10),
			}.
				Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					return layout.Flex{
						Axis:    layout.Horizontal,
						Spacing: layout.SpaceBetween,
					}.Layout(gtx,
						layout.Flexed(3, func(gtx layout.Context) layout.Dimensions {
							return layout.Inset{
								Right: unit.Dp(5),
							}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
								return material.Button(th, &store.StatusBackground, store.DateStatus).Layout(gtx)
							})
						}),
						layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
							return layout.Inset{
								Left: unit.Dp(5),
							}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
								btn := material.Button(th, &store.StatusBackground, "Eth")
								btn.Background = store.EthernetStatus.BgColor
								return btn.Layout(gtx)
							})
						}),
					)
				})
		}),
		layout.Flexed(0.6, func(gtx layout.Context) layout.Dimensions {
			return layout.Inset{
				Bottom: unit.Dp(10),
			}.
				Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					return layout.Flex{
						Axis:    layout.Horizontal,
						Spacing: layout.SpaceBetween,
					}.Layout(gtx,
						layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
							return layout.Inset{
								Right: unit.Dp(8),
							}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
								return material.Button(th, &store.StatusBackground, "Vo +/-").Layout(gtx)
							})
						}),
						layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
							return layout.Inset{
								Right: unit.Dp(5),
								Left:  unit.Dp(2),
							}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
								return material.Button(th, &store.StatusBackground, "Br +/-").Layout(gtx)
							})
						}),

						layout.Flexed(2, func(gtx layout.Context) layout.Dimensions {
							return layout.Inset{
								Left: unit.Dp(5),
							}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
								return material.Button(th, &store.StatusBackground, "Change Sound Mode").Layout(gtx)
							})
						}),
					)
				})
		}),
	)
}
