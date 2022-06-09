package cli

import (
	"log"

	ui "github.com/gizak/termui/v3"
	"github.com/lipefebn/spotify-crud/src/adapter/query"
	"github.com/lipefebn/spotify-crud/src/infra/config"
)

func render(header, tab ui.Drawable) {
	ui.Render(tab, header)
	if resultsQuery[0] != nil {
		ui.Render(resultsQuery[0])
	}
	if resultsQuery[1] != nil {
		ui.Render(resultsQuery[1])
	}
	if resultsQuery[2] != nil {
		ui.Render(resultsQuery[2])
	}
}

var (
	connectionDb = config.Connect("..."/* path sqlite */)
	queryDb = query.DB{ Db: connectionDb }
	resultsQuery = [3]ui.Drawable{} 
	uiEvents = ui.PollEvents()
)

func renderChannel() (chan<- ui.Drawable){
	recepetor := make(chan ui.Drawable, 1)
	go func() {
		for i := 0; i < 3; i++ {
			resultsQuery[i] = <-recepetor
			ui.Render(resultsQuery[i])
		}
		close(recepetor)
	}()
	return recepetor
}


func Init() {
	chans := renderChannel()
	go Querys(queryDb, chans)

	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	header := Header()
	tab := Tab{ InitTab() }
	render(header, tab.obj)

	for {
		e := <-uiEvents
		switch e.ID {
			case "q", "<C-c>":
				return
			case "<Left>":
				tab.Left()
				ui.Clear()
				render(header, tab.obj)
			case "<Right>":
				tab.Right()
				ui.Clear()
				render(header, tab.obj)
			case "<Enter>":
				switch tab.obj.ActiveTabIndex {
				case 0:
					InitAlbumCli()
				case 1:
					InitArtistCli()
				case 2:
					InitGenreCli()
				case 3:
					InitTrackCli()
				}
				ui.Clear()
				render(header, tab.obj)
			case "<Resize>":
				ui.Clear()
				render(header, tab.obj)
		}
	}
}