package cli

import (
	"strconv"

	"github.com/lipefebn/spotify-crud/src/adapter/relational"
	"github.com/lipefebn/spotify-crud/src/core/repository"
)

func InitArtistCli() {
	PrintPageCrud("Album")
	switch (<-uiEvents).ID {
		case "1":
			createArtist()
		case "2":
			readArtist()
		case "3":
			updateArtist()
		case "4":
			deleteArtist()
		case "q":
			return	
	}
}

func readArtist() {
	forms := createForms(80, 5, "Name")
	forms.Draw()
	read := InitModel(*relational.NewTable[repository.Artist](connectionDb), "artist")
	read.Read("popularity DESC", forms.ChooseUser()[0], "name")
}

func updateArtist() {
	artist := relational.NewTable[repository.Artist](connectionDb)
	forms := createForms(80, 5, "ID", "Name", "Popularity", "Followers")
	forms.Draw()
	result := forms.ChooseUser()
	popularity, _ := strconv.Atoi(result[2])
	followers, _ := strconv.Atoi(result[3])
	artist.Update(result[0], repository.Artist { 
		ID: result[0],
		Name: result[1],		
		Popularity: popularity,
		Followers: followers,
	})
}

func createArtist(){
	artist := relational.NewTable[repository.Artist](connectionDb)
	forms := createForms(80, 5, "ID", "Name", "Popularity", "Followers")
	forms.Draw()
	result := forms.ChooseUser()
	popularity, _ := strconv.Atoi(result[2])
	followers, _ := strconv.Atoi(result[3])
	artist.Create(repository.Artist { 
		ID: result[0],
		Name: result[1],
		Popularity: popularity,
		Followers: followers,
	})
}

func deleteArtist() {
	artist := relational.NewTable[repository.Artist](connectionDb)
	forms := createForms(80, 5, "id")
	forms.Draw()
	id := forms.ChooseUser()[0]
	artist.Delete(id, repository.Artist{})

}