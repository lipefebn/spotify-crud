package cli

import (
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

var crudTables = []string{
	"ALBUM",
	"ARTIST",
	"GENRE",
	"TRACK",
}

func InitTab() (result *widgets.TabPane) {
	result = widgets.NewTabPane(crudTables...)
	result.SetRect(5, 5, 40, 9)
	result.Title = "TABLES"
	result.TitleStyle = ui.NewStyle(ui.ColorGreen)
	result.BorderStyle = ui.NewStyle(ui.ColorBlack)
	result.BorderStyle.Fg = ui.ColorWhite
	result.ActiveTabStyle.Fg = ui.ColorBlue
	result.Border = true
	return
}


type Tab struct {
	obj *widgets.TabPane
}

func (t Tab) Left() { t.obj.FocusLeft() }
func (t Tab) Right() { t.obj.FocusRight() }