// Package state
package state

import "gioui.org/widget"

type PowerControlState struct {
	PowerOffButton   widget.Clickable
	RebootButton     widget.Clickable
	SuspendButton    widget.Clickable
	HibernateButton  widget.Clickable
	LockScreenButton widget.Clickable

	ToolTipLabel string
}

type StatusControlState struct {
	StatusBackground widget.Clickable

	BatteryStatus string
}

type UIState struct {
	PowerControlState
	StatusControlState
}
