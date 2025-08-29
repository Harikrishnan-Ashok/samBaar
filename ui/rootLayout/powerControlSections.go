// Package rootlayout
package rootlayout

import (
	"fmt"
	"os/exec"

	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

var powerOffButton widget.Clickable
var rebootButton widget.Clickable
var suspendButton widget.Clickable
var hibernateButton widget.Clickable
var lockScreenButton widget.Clickable

func PowerControlSection(gtx layout.Context, th *material.Theme) layout.Dimensions {

	if powerOffButton.Clicked(gtx) {
		cmd := exec.Command("systemctl", "poweroff")
		if err := cmd.Run(); err != nil {
			fmt.Println("Command Failed")
		}
	}
	if rebootButton.Clicked(gtx) {
		cmd := exec.Command("systemctl", "reboot")
		if err := cmd.Run(); err != nil {
			fmt.Println("Command Failed")
		}
	}
	if suspendButton.Clicked(gtx) {
		cmd := exec.Command("systemctl", "suspend")
		if err := cmd.Run(); err != nil {
			fmt.Println("Command Failed")
		}
	}
	if hibernateButton.Clicked(gtx) {
		cmd := exec.Command("systemctl", "hibernate")
		if err := cmd.Run(); err != nil {
			fmt.Println("Command Failed")
		}
	}
	if lockScreenButton.Clicked(gtx) {
		cmd := exec.Command("swaylock")
		if err := cmd.Run(); err != nil {
			fmt.Println("Command Failed")
		}
	}
	return layout.Flex{
		Axis:    layout.Horizontal,
		Spacing: layout.SpaceAround,
	}.Layout(gtx,
		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			return layout.Inset{
				Right: unit.Dp(4),
				Left:  unit.Dp(4),
			}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				btn := material.Button(th, &powerOffButton, "⏻")
				return btn.Layout(gtx)
			})
		}),
		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			return layout.Inset{
				Right: unit.Dp(4),
				Left:  unit.Dp(4),
			}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				btn := material.Button(th, &rebootButton, "󰜉")
				return btn.Layout(gtx)
			})
		}),
		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			return layout.Inset{
				Right: unit.Dp(4),
				Left:  unit.Dp(4),
			}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				btn := material.Button(th, &suspendButton, "󰒲")
				return btn.Layout(gtx)
			})
		}),
		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			return layout.Inset{
				Right: unit.Dp(4),
				Left:  unit.Dp(4),
			}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				btn := material.Button(th, &hibernateButton, "")
				return btn.Layout(gtx)
			})
		}),
		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			return layout.Inset{
				Right: unit.Dp(4),
				Left:  unit.Dp(4),
			}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				btn := material.Button(th, &lockScreenButton, "󰌾")
				return btn.Layout(gtx)
			})
		}),
	)

}
