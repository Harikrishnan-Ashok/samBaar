package sections

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"github.com/Harikrishnan-Ashok/samBaar/state"
)

func StatusControlSection(gtx layout.Context, th *material.Theme, store *state.UIState) layout.Dimensions {
	return layout.Flex{
		Axis:    layout.Horizontal,
		Spacing: layout.SpaceBetween,
	}.Layout(gtx,

		// Left column (single button)
		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			return layout.Inset{
				Right: unit.Dp(10),
			}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				btn := material.Button(th, &store.StatusBackground, store.BatteryStatus)
				btn.Inset = layout.UniformInset(unit.Dp(4))
				return btn.Layout(gtx)
			})
		}),

		// Right column (vertical stack of buttons)
		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{
				Axis:    layout.Vertical,
				Spacing: layout.SpaceBetween,
			}.Layout(gtx,

				// Top button
				layout.Flexed(2, func(gtx layout.Context) layout.Dimensions {
					btn := material.Button(th, &store.StatusBackground, store.TimeStatus)
					btn.TextSize = unit.Sp(40)
					btn.Inset = layout.UniformInset(unit.Dp(4))
					btn.Inset.Top = unit.Dp(25)
					return btn.Layout(gtx)
				}),

				// Bottom section with two buttons, wrapped in inset
				layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
					return layout.Inset{
						Top: unit.Dp(10),
					}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						return layout.Flex{
							Axis:    layout.Horizontal,
							Spacing: layout.SpaceBetween,
						}.Layout(gtx,
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								btn := material.Button(th, &store.StatusBackground, store.WifiStatus)
								return btn.Layout(gtx)
							}),
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								btn := material.Button(th, &store.StatusBackground, store.BluetoothStatus)
								return btn.Layout(gtx)
							}),
						)
					})
				}),
			)
		}),
	)
}
