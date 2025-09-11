package sections

import (
	"fmt"
	"gioui.org/app"
	"gioui.org/io/key"
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/Harikrishnan-Ashok/samBaar/state"
)

func ShortCutSection(gtx layout.Context, th *material.Theme, store *state.UIState, w *app.Window) layout.Dimensions {
	box := &store.TextBox
	box.SingleLine = true
	box.Submit = true

	// Auto-focus on first frame only
	if !store.EditorHasFocused {
		gtx.Execute(key.FocusCmd{Tag: box})
		store.EditorHasFocused = true
	}

	// Use the modern Editor.Update() method
	for {
		event, ok := box.Update(gtx)
		if !ok {
			break
		}

		switch e := event.(type) {
		case widget.ChangeEvent:
			// Optional: handle text changes
			fmt.Printf("Text: %s\n", box.Text())

		case widget.SubmitEvent:
			fmt.Printf("Editor submitted: %s\n", e.Text)
			// Keep focus after submit
			gtx.Execute(key.FocusCmd{Tag: box})
		}
	}

	SearchBox := material.Editor(th, box, "Type to search shortcts")
	return SearchBox.Layout(gtx)
}
