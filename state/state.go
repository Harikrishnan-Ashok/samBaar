// Package state
package state

import (
	"image/color"

	"gioui.org/widget"
)

type Shortcut struct {
	Token   string `json:"token"`
	Command string `json:"command"`
}

type ConfigData struct {
	Shortcuts []Shortcut `json:"shortcuts"`
}

type AppConfig struct {
	Config ConfigData `json:"config"`
}

var DefaultConfig = AppConfig{
	Config: ConfigData{
		Shortcuts: []Shortcut{},
	},
}

type PowerControlState struct {
	PowerOffButton   widget.Clickable
	RebootButton     widget.Clickable
	SuspendButton    widget.Clickable
	HibernateButton  widget.Clickable
	LockScreenButton widget.Clickable

	ToolTipLabel string
}

type SeachEditor struct {
	TextBox          widget.Editor
	LastText         string
	EditorHasFocused bool
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
	SeachEditor
	ShouldQuit bool
}
