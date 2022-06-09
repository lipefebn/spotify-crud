package cli

import (
	"strconv"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/lipefebn/spotify-crud/src/adapter/query"
	"github.com/lipefebn/spotify-crud/src/adapter/relational"
	"github.com/lipefebn/spotify-crud/src/core/repository"
)

func topArtistsCli() {
	text := widgets.NewParagraph()
	text.Text = "DIGITE [9] (TOP X ARTISTS OF ANY GENRE)"
	text.SetRect(0, 1, 45, 4)
	ui.Render(text)
}

func renderGenre() {
	PrintPageCrud("Genre")
	topArtistsCli()
}

func InitGenreCli() {
	renderGenre()
	for {
		switch (<-uiEvents).ID {
			case "1":
				createGenre()
			case "2":
				readGenre()
			case "3":
				updateGenre()
			case "4":
				deleteGenre()
			case "8":
				topArtists()
			case "9":
				topArtists()
			case "q":
				return	
		}
		renderGenre()
	}
}

func topArtists() {
	query := query.DB{ Db: connectionDb }
	forms := createForms(80, 5, "Id", "Number of artists")
	forms.Draw()
	inputs := forms.ChooseUser()
	limit, _ := strconv.Atoi(inputs[1])
	result := query.TopArtistsGenre(inputs[0], limit)
	table := widgets.NewTable()
	table.RowStyles[0] = ui.NewStyle(ui.ColorRed)
	table.SetRect(0, 0, 200, 100)
	table.Rows = [][]string { { "ID",  "Name", "Popularity"} }
	for _, elem := range(result) {
		popularity := strconv.Itoa(elem.Popularity)
		table.Rows = append(table.Rows, []string{ elem.ID, elem.Name,  popularity })
	}

	ui.Clear()
	ui.Render(table)

	for { if (<-uiEvents).ID == "q" { return } }		
}

func createGenre(){
	genre := relational.NewTable[repository.Genre] (connectionDb)
	forms := createForms(80, 5, "Id")
	forms.Draw()
	result := forms.ChooseUser()
	genre.Create(repository.Genre { 
		ID: result[0],
	})
}
func readGenre() {
	forms := createForms(80, 5, "Id")
	forms.Draw()
	read := InitModel(*relational.NewTable[repository.Genre](connectionDb), "genre")
	read.Read("id", forms.ChooseUser()[0], "id")
}

func updateGenre() {
	genre := relational.NewTable[repository.Genre](connectionDb)
	forms := createForms(80, 5, "Id")
	forms.Draw()
	result := forms.ChooseUser()
	genre.Update(result[0], repository.Genre { 
		ID: result[0],
	})
}


func deleteGenre() {
	genre := relational.NewTable[repository.Genre](connectionDb)
	forms := createForms(80, 5, "Id")
	forms.Draw()
	id := forms.ChooseUser()[0]
	genre.Delete(id, repository.Genre{})

}