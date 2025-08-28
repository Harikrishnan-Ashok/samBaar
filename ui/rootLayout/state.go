package rootlayout

import "gioui.org/widget"

type powerControlState struct {
	poweroffButton widget.Clickable
	rebootButton   widget.Clickable
}

type UiState struct {
	powerControlState
}
