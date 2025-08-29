// Package sections
package sections

import (
	"fmt"
	"os"
	"os/exec"

	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/Harikrishnan-Ashok/samBaar/state"
)

// helper ro run command
func runCmd(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	if err := cmd.Run(); err != nil {
		fmt.Println("Command Failed")
		return err
	}
	return nil
}

// helper to replace button
func actionButton(gtx layout.Context, th *material.Theme, action *widget.Clickable, icon string) layout.Dimensions {
	return layout.Inset{
		Right: unit.Dp(4),
		Left:  unit.Dp(4),
	}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		btn := material.Button(th, action, icon)
		return btn.Layout(gtx)
	})
}

func PowerControlSection(gtx layout.Context, th *material.Theme, store *state.UIState) layout.Dimensions {

	if store.PowerOffButton.Clicked(gtx) {
		runCmd("systemctl", "poweroff")
		os.Exit(0)
	}
	if store.RebootButton.Clicked(gtx) {
		runCmd("systemctl", "reboot")
		os.Exit(0)
	}
	if store.SuspendButton.Clicked(gtx) {
		runCmd("systemctl", "suspend")
		os.Exit(0)
	}
	if store.HibernateButton.Clicked(gtx) {
		runCmd("systemctl", "hibernate")
		os.Exit(0)
	}
	if store.LockScreenButton.Clicked(gtx) {
		runCmd("swaylock")
		os.Exit(0)
	}
	return layout.Flex{
		Axis:    layout.Horizontal,
		Spacing: layout.SpaceAround,
	}.Layout(gtx,
		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			return actionButton(gtx, th, &store.PowerOffButton, "⏻")
		}),
		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			return actionButton(gtx, th, &store.RebootButton, "󰜉")
		}),
		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			return actionButton(gtx, th, &store.SuspendButton, "󰒲")
		}),
		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			return actionButton(gtx, th, &store.HibernateButton, "")
		}),
		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			return actionButton(gtx, th, &store.LockScreenButton, "󰌾")
		}),
	)

}
