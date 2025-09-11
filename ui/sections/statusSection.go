package sections

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"github.com/Harikrishnan-Ashok/samBaar/state"
	"github.com/Harikrishnan-Ashok/samBaar/theme"
)

func StatusControlSection(gtx layout.Context, th *material.Theme, store *state.UIState) layout.Dimensions {

	switch {
	case store.WifiStatus.Button.Clicked(gtx):
		if store.WifiStatus.Value == "enabled" {
			store.WifiStatus.Value = "disabled"
			store.WifiStatus.BgColor = theme.DarkTheme.Colors.DisabledColor
			go runCmd("nmcli", "radio", "wifi", "off")
		} else {
			store.WifiStatus.Value = "enabled"
			store.WifiStatus.BgColor = theme.DarkTheme.Colors.EnabledColor
			go runCmd("nmcli", "radio", "wifi", "on")
		}

	case store.BluetoothStatus.Button.Clicked(gtx):
		if store.BluetoothStatus.Value == "enabled" {
			store.BluetoothStatus.Value = "disabled"
			store.BluetoothStatus.BgColor = theme.DarkTheme.Colors.DisabledColor
			go runCmd("bluetoothctl", "power", "off")
		} else {
			store.BluetoothStatus.Value = "enabled"
			store.BluetoothStatus.BgColor = theme.DarkTheme.Colors.EnabledColor
			go runCmd("bluetoothctl", "power", "on")
		}
	}

	return layout.Flex{
		Axis:    layout.Horizontal,
		Spacing: layout.SpaceBetween,
	}.Layout(gtx,

		// Left column (single button)
		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			return layout.Inset{
				Right: unit.Dp(10),
			}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return layout.Stack{}.Layout(gtx,
					layout.Stacked(func(gtx layout.Context) layout.Dimensions {
						gtx.Constraints.Min = gtx.Constraints.Max
						btn := material.Button(th, &store.StatusButton, "")
						btn.Background = store.BgColor
						return btn.Layout(gtx)
					}),
					layout.Stacked(func(gtx layout.Context) layout.Dimensions {
						return layout.Flex{
							Axis:    layout.Vertical,
							Spacing: layout.SpaceSides,
						}.Layout(gtx,
							layout.Flexed(1.8, func(gtx layout.Context) layout.Dimensions {
								btn := material.Button(th, &store.BatteryValue.Button, store.BatteryValue.Value)
								btn.TextSize = unit.Sp(35)
								btn.Background = theme.DarkTheme.Colors.TrasnparentBg
								btn.Inset = layout.UniformInset(unit.Dp(5))
								return btn.Layout(gtx)
							}),
							layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
								btn := material.Button(th, &store.AdapterStatus.Button, store.AdapterStatus.Value)
								btn.Background = theme.DarkTheme.Colors.TrasnparentBg
								btn.TextSize = unit.Sp(20)
								btn.Inset = layout.UniformInset(unit.Dp(5))
								return btn.Layout(gtx)
							}),
							layout.Flexed(1.2, func(gtx layout.Context) layout.Dimensions {
								btn := material.Button(th, &store.RemainingTime.Button, store.RemainingTime.Value)
								btn.Background = theme.DarkTheme.Colors.TrasnparentBg
								btn.Inset = layout.UniformInset(unit.Dp(5))
								return btn.Layout(gtx)
							}),
						)
					}),
				)
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
					btn := material.Button(th, &store.TimeStatus.Button, store.TimeStatus.Value)
					btn.TextSize = unit.Sp(40)
					btn.Inset = layout.UniformInset(unit.Dp(4))
					btn.Inset.Top = unit.Dp(15)
					return btn.Layout(gtx)
				}),

				// Bottom section with two buttons, wrapped in inset
				layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
					return layout.Inset{
						Top: unit.Dp(10),
					}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						return layout.Flex{
							Axis:    layout.Horizontal,
							Spacing: layout.SpaceAround,
						}.Layout(gtx,
							layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
								return layout.Inset{
									Right: unit.Dp(5),
								}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
									btn := material.Button(th, &store.WifiStatus.Button, "Wifi")
									btn.Background = store.WifiStatus.BgColor
									return btn.Layout(gtx)
								})
							}),
							layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
								return layout.Inset{
									Left: unit.Dp(5),
								}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
									btn := material.Button(th, &store.BluetoothStatus.Button, "Blue")
									btn.Background = store.BluetoothStatus.BgColor
									return btn.Layout(gtx)
								})
							}),
						)
					})
				}),
			)
		}),
	)
}
