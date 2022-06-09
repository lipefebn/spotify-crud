package cli

import (
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func Header() ui.Drawable{
	header := widgets.NewParagraph()
	header.Text = "CLICK ARROWS ( ← and → ) TO SWITCH TABS AND <ENTER> TO SELECT TABLE"
	header.SetRect(1, 1, 74, 4)
	header.TextStyle = ui.NewStyle(ui.ColorRed)
	return header 
}