package sections

import (
	"fmt"
	"os/exec"

	"gioui.org/app"
	"gioui.org/io/key"
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/Harikrishnan-Ashok/samBaar/helpers"
	"github.com/Harikrishnan-Ashok/samBaar/state"
)

func RunCmdAsync(cmd string, store *state.UIState) {
	go func() {
		fmt.Println("Running:", cmd)
		command := exec.Command("sh", "-c", cmd)
		_, err := command.CombinedOutput()
		if err != nil {
			fmt.Println("Command failed:", err)
		}
	}()
	store.ShouldQuit = true
}

func ShortCutSection(gtx layout.Context, th *material.Theme, store *state.UIState, w *app.Window) layout.Dimensions {
	box := &store.TextBox
	box.SingleLine = true
	box.Submit = true

	if !store.EditorHasFocused {
		gtx.Execute(key.FocusCmd{Tag: box})
		store.EditorHasFocused = true
	}

	for {
		event, ok := box.Update(gtx)
		if !ok {
			break
		}
		switch e := event.(type) {
		case widget.SubmitEvent:
			cmd, err := helpers.SearchForCommandKeyword(e.Text)
			if err != nil {
				fmt.Println("Command lookup error:", err)
			} else {
				RunCmdAsync(cmd, store)
			}
			box.SetText("")
			gtx.Execute(key.FocusCmd{Tag: box})
		}
	}

	searchBox := material.Editor(th, box, "Type to search shortcuts")
	return searchBox.Layout(gtx)
}
