package rootlayot

import (
	"gioui.org/layout"
	"gioui.org/widget/material"
	"github.com/Harikrishnan-Ashok/samBaar/ui/customTypes"
)

func PowerControlSection(gtx layout.Context, th *material.Theme) layout.Dimensions {
	var powerOffButton, rebootButton customtypes.TypeButton
	return layout.Flex{
		Axis:    layout.Horizontal,
		Spacing: layout.SpaceAround,
	}.Layout(gtx, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
		btn := material.Button(th, &powerOffButton, "Poweroff")
		return btn.Layout(gtx)
	}), layout.Rigid(func(gtx layout.Context) layout.Dimensions {
		btn := material.Button(th, &rebootButton, "Reboot")
		return btn.Layout(gtx)
	}),
	)
}
