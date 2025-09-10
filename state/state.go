// Package state
package state

import (
	"image/color"

	"gioui.org/widget"
)

type PowerControlState struct {
	PowerOffButton   widget.Clickable
	RebootButton     widget.Clickable
	SuspendButton    widget.Clickable
	HibernateButton  widget.Clickable
	LockScreenButton widget.Clickable

	ToolTipLabel string
}

// Status ...  for status and background color
type Status struct {
	Button  widget.Clickable
	Value   string
	BgColor color.NRGBA
}

type BatteryStatus struct {
	BatteryValue  Status
	AdapterStatus Status
	RemainingTime Status
	BgColor       color.NRGBA
	StatusButton  widget.Clickable
}

type StatusControlState struct {
	TimeStatus Status
	BatteryStatus
	DateStatus      Status
	WifiStatus      Status
	BluetoothStatus Status
	EthernetStatus  Status
	VolStatus       Status
	Brightness      Status
	SoundMode       Status
}

type UtilsControlState struct {
	TimerStartButton Status
	TimerStatus      Status
	NextWallpaper    widget.Clickable
	PrevWallpaper    widget.Clickable
	WallpaperStatus  widget.Clickable
}

type UIState struct {
	PowerControlState
	StatusControlState
	UtilsControlState
}
