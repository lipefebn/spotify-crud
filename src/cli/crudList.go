package cli

import (
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func PrintPageCrud(table string) {
	title := widgets.NewParagraph()
	title.Text = "[TABLE " + table + "](bg:red)"
	title.SetRect(90, 0, 110, 4)
	title.Border = false

	list := widgets.NewList()
	list.Rows = []string {
		"[1] Create",
		"[2] Read",
		"[3] Update",
		"[4] Delete",
	}

	list.Title = "SELECT THE OPERATION"
	list.PaddingLeft = 5
	list.PaddingTop = 1
	list.TitleStyle = ui.NewStyle(ui.ColorRed)
	list.BorderStyle = ui.NewStyle(ui.ColorBlue)
	list.SetRect(2, 5, 25, 13)
	ui.Clear()
	ui.Render(list, title)

}