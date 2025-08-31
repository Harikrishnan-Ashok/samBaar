package sections

import (
	"gioui.org/layout"
	"gioui.org/widget/material"
	"github.com/Harikrishnan-Ashok/samBaar/state"
)

func StatusControlSection(gtx layout.Context, th *material.Theme, store *state.UIState) layout.Dimensions {
	return layout.Flex{
		Axis:    layout.Horizontal,
		Spacing: layout.SpaceBetween,
	}.Layout(gtx,
		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			return layout.Inset{
				Right: 10,
			}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				btn := material.Button(th, &store.StatusBackground, store.BatteryStatus)
				return btn.Layout(gtx)
			})
		}),
		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{
				Axis:    layout.Horizontal,
				Spacing: layout.SpaceAround,
			}.Layout(gtx,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return layout.Flex{
						Axis:    layout.Vertical,
						Spacing: layout.SpaceAround,
					}.Layout(gtx,
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							btn := material.Button(th, &store.StatusBackground, "test")
							return btn.Layout(gtx)
						}),
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							btn := material.Button(th, &store.StatusBackground, "test")
							return btn.Layout(gtx)
						}),
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							btn := material.Button(th, &store.StatusBackground, "test")
							return btn.Layout(gtx)
						}))
				}),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return layout.Flex{
						Axis:    layout.Vertical,
						Spacing: layout.SpaceAround,
					}.Layout(gtx,
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							btn := material.Button(th, &store.StatusBackground, "test")
							return btn.Layout(gtx)
						}),
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							btn := material.Button(th, &store.StatusBackground, "test")
							return btn.Layout(gtx)
						}),
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							btn := material.Button(th, &store.StatusBackground, "test")
							return btn.Layout(gtx)
						}),
					)
				}),
			)
		}),
	)
}
