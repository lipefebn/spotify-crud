package cli

import (
	"strconv"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/lipefebn/spotify-crud/src/adapter/query"
	"github.com/lipefebn/spotify-crud/src/adapter/relational"
	"github.com/lipefebn/spotify-crud/src/core/repository/spotify"
)

func topTracksCli() {
	text := widgets.NewParagraph()
	text.Text = "DIGITE [9] (TOP TRACKS OF ANY ALBUM)"
	text.SetRect(0, 1, 45, 4)
	ui.Render(text)
}

func InitAlbumCli() {
	PrintPageCrud("Album")
	topTracksCli()
	for {
		switch (<-uiEvents).ID {
			case "1":
				createAlbum()
			case "2":
				readAlbum()
			case "3":
				updateAlbum()
			case "4":
				deleteAlbum()
			case "9":
				topTracksAlbums()
			case "q":
				return
		}
		PrintPageCrud("Album")
		topTracksCli()
	}
}

func topTracksAlbums() {
	forms := createForms(80, 5, "Name")
	forms.Draw()
	resultInput := relational.NewTable[repository.Album](connectionDb).Select("popularity DESC", "name LIKE ?", forms.ChooseUser())
	
	//---- len(resultInput) == 0 -> return

	query := query.DB{ Db: connectionDb }
	result := query.TracksToAlbum(resultInput[0])
	table := widgets.NewTable()
	table.RowStyles[0] = ui.NewStyle(ui.ColorRed)
	table.SetRect(0, 0, 200, 100)
	table.Rows = [][]string { { "Name", "Popularity"} }
	table.Title = resultInput[0].Name
	for _, elem := range(result) {
		popularity := strconv.Itoa(elem.Popularity)
		table.Rows = append(table.Rows, []string{ elem.Name,  popularity})
	}

	ui.Clear()
	ui.Render(table)

	for { if (<-uiEvents).ID == "q" { return } }		
}

func readAlbum() {
	forms := createForms(80, 5, "Name")
	forms.Draw()
	read := InitModel(*relational.NewTable[repository.Album](connectionDb), "album")
	read.Read("popularity DESC", forms.ChooseUser()[0], "name")
}

func updateAlbum() {
	album := relational.NewTable[repository.Album](connectionDb)
	forms := createForms(80, 5, "ID", "Name", "Album_group", "Album_type", "Release_date", "Popularity")
	forms.Draw()
	result := forms.ChooseUser()
	releaseDate, _ := strconv.Atoi(result[4])
	popularity, _ := strconv.Atoi(result[5])
	album.Update(result[0], repository.Album { 
		ID: result[0],
		Name: result[1],
		Album_group: result[2],
		Album_type: result[3],
		Release_date: int64(releaseDate),
		Popularity: popularity,
	})
}

func createAlbum(){
	album := relational.NewTable[repository.Album](connectionDb)
	forms := createForms(80, 5, "ID", "Name", "Album_group", "Album_type", "Release_date", "Popularity")
	forms.Draw()
	result := forms.ChooseUser()
	releaseDate, _ := strconv.Atoi(result[4])
	popularity, _ := strconv.Atoi(result[5])
	album.Create(repository.Album { 
		ID: result[0],
		Name: result[1],
		Album_group: result[2],
		Album_type: result[3],
		Release_date: int64(releaseDate),
		Popularity: popularity,
	})
}

func deleteAlbum() {
	album := relational.NewTable[repository.Album](connectionDb)
	forms := createForms(80, 5, "id")
	forms.Draw()
	id := forms.ChooseUser()[0]
	album.Delete(id, repository.Album{})

}