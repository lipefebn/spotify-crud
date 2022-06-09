package cli

import (
	"strconv"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/lipefebn/spotify-crud/src/adapter/query"
)

func Querys(db query.DB, recepetor chan<- ui.Drawable){
	go artistToTracks(db, recepetor)
	go artistToAlbum(db, recepetor)
	go topTracks(db, recepetor)
}

func artistToAlbum(db query.DB, recepetor chan<- ui.Drawable) {
	querys := db.ArtistToAlbums()
	result := widgets.NewTable()
	result.Rows = [][]string{ {"Name", "Amount Albums"} }
	for _, line := range(querys) {
		result.Rows = append(result.Rows, []string{line.Name, strconv.Itoa(line.Amount)})
	}
	result.Title = "TOP 10 ARTISTS WITH MORE ALBUMS"
	result.TextStyle = ui.NewStyle(ui.ColorWhite)
	result.RowSeparator = true
	result.BorderStyle = ui.NewStyle(ui.ColorGreen)
	result.SetRect(160, 28, 200, 51)
	result.FillRow = true
	result.RowStyles[0] = ui.NewStyle(ui.ColorWhite, ui.ColorBlack, ui.ModifierBold)
	recepetor <- result	
}

func artistToTracks(db query.DB, recepetor chan<- ui.Drawable){
	querys := db.ArtistToTracks()
	result := widgets.NewTable()
	result.Rows = [][]string{ {"Name", "Amount Tracks"} }
	for _, line := range(querys) {
		result.Rows = append(result.Rows, []string{line.Name, strconv.Itoa(line.Cont)})
	}
	result.Title = "TOP 10 ARTISTS WITH MORE TRACKS"
	result.TextStyle = ui.NewStyle(ui.ColorWhite)
	result.RowSeparator = true
	result.BorderStyle = ui.NewStyle(ui.ColorGreen)
	result.SetRect(160, 3, 200, 26)
	result.FillRow = true
	result.RowStyles[0] = ui.NewStyle(ui.ColorWhite, ui.ColorBlack, ui.ModifierBold)
	recepetor <- result
}

func topTracks(db query.DB, recepetor chan<- ui.Drawable) {
	querys := db.TopTracks()
	result := widgets.NewTable()
	result.Rows = [][]string{ {"Name", "Popularity"} }
	for _, line := range(querys) {
		result.Rows = append(result.Rows, []string{line.Name, strconv.Itoa(line.Popularity)})
	} 
	result.Title = "TOP 10 Tracks"
	result.TextStyle = ui.NewStyle(ui.ColorWhite)
	result.RowSeparator = true
	result.BorderStyle = ui.NewStyle(ui.ColorGreen)
	result.SetRect(110, 28, 155, 51)
	result.FillRow = true
	result.RowStyles[0] = ui.NewStyle(ui.ColorWhite, ui.ColorBlack, ui.ModifierBold)
	recepetor <- result
}