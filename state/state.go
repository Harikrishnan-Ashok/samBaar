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
	Value   string
	BgColor color.NRGBA
}

type BatteryStatus struct {
	BatteryValue  string
	AdapterStatus string
	RemainingTime string
	BgColor       color.NRGBA
}

type StatusControlState struct {
	TimeStatus string
	BatteryStatus
	DateStatus      string
	WifiStatus      Status
	BluetoothStatus Status
	EthernetStatus  Status
}

type UtilsControlState struct {
}

type UIState struct {
	PowerControlState
	StatusControlState
	UtilsControlState

	StatusBackground widget.Clickable
	StatusForeground widget.Clickable
}
