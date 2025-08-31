// Package sections
package sections

import (
	"fmt"
	"os"
	"os/exec"

	"gioui.org/app"
	"gioui.org/io/event"
	"gioui.org/io/key"
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/Harikrishnan-Ashok/samBaar/state"
)

var powerControlTag struct{}

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

func PowerControlSection(gtx layout.Context, th *material.Theme, store *state.UIState, w *app.Window) layout.Dimensions {
	event.Op(gtx.Ops, powerControlTag)
	gtx.Execute(key.FocusCmd{Tag: powerControlTag})

	for {
		ev, ok := gtx.Event(
			key.FocusFilter{Target: powerControlTag},
			key.Filter{Focus: powerControlTag, Name: key.NameLeftArrow},
			key.Filter{Focus: powerControlTag, Name: key.NameRightArrow},
		)
		if !ok {
			break
		}
		ke, ok := ev.(key.Event)
		if !ok {
			continue
		}
		if ke.State == key.Press {
			switch ke.Name {
			case key.NameLeftArrow:
				store.ToolTipLabel = "Left"
				fmt.Println("LeftClicked")
				w.Invalidate()
			case key.NameRightArrow:
				store.ToolTipLabel = "Right"
				fmt.Println("RightClicked")
				w.Invalidate()
			}
		}
	}

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

	//for hovers
	switch {
	case store.PowerOffButton.Hovered():
		store.ToolTipLabel = "Power Off"
	case store.RebootButton.Hovered():
		store.ToolTipLabel = "Reboot"
	case store.SuspendButton.Hovered():
		store.ToolTipLabel = "Suspend"
	case store.HibernateButton.Hovered():
		store.ToolTipLabel = "Hibernate"
	case store.LockScreenButton.Hovered():
		store.ToolTipLabel = "Lock Screen"
	default:
		store.ToolTipLabel = ""
	}

	return layout.Flex{
		Axis:    layout.Vertical,
		Spacing: layout.SpaceBetween,
	}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Inset{
				Bottom: unit.Dp(5),
				Left:   unit.Dp(4),
			}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return material.Label(th, unit.Sp(15), store.ToolTipLabel).Layout(gtx)
			})
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			// Horizontal row of buttons
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
		}),
	)
}
