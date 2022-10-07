package cli

import (
	"strconv"

	"github.com/lipefebn/spotify-crud/src/adapter/relational"
	"github.com/lipefebn/spotify-crud/src/core/repository/spotify"
)

func InitTrackCli() {
	PrintPageCrud("Album")
	switch (<-uiEvents).ID {
		case "1":
			createTrack()
		case "2":
			readTrack()
		case "3":
			updateTrack()
		case "4":
			deleteTrack()
		case "q":
			return	
	}
}

func readTrack() {
	forms := createForms(80, 5, "Name")
	forms.Draw()
	read := InitModel(*relational.NewTable[repository.Track](connectionDb), "track")
	read.Read("popularity DESC", forms.ChooseUser()[0], "name")
}

func updateTrack() {
	track := relational.NewTable[repository.Track](connectionDb)
	forms := createForms(80, 5, "ID", "Disc Number", "Duration", "Explicit", "Name", "Preview url", "track number", "Popularity", "isPlyable")
	forms.Draw()
	result := forms.ChooseUser()

	convertStr := func(str string) int {
		result, _ := strconv.Atoi(str)
		return result
	}

	isPlayable, _ := strconv.ParseBool(result[8])

	track.Update(result[0], repository.Track { 
		ID: result[0],
		Name: result[1],	
		Disc_number: convertStr(result[2]),	
		Duration: convertStr(result[3]),
		Explicit: convertStr(result[4]),
		Audio_feature_id: "",
		Preview_url: result[5],
		Track_number: convertStr(result[6]),
		Popularity: convertStr(result[7]),
		Is_playable: isPlayable,
	})
}

func createTrack(){
	track := relational.NewTable[repository.Track](connectionDb)
	forms := createForms(80, 5, "ID", "Disc Number", "Duration", "Explicit", "Name", "Preview url", "track number", "Popularity", "isPlyable")
	forms.Draw()
	result := forms.ChooseUser()

	convertStr := func(str string) int {
		result, _ := strconv.Atoi(str)
		return result
	}

	isPlayable, _ := strconv.ParseBool(result[8])

	track.Create(repository.Track { 
		ID: result[0],
		Name: result[1],	
		Disc_number: convertStr(result[2]),	
		Duration: convertStr(result[3]),
		Explicit: convertStr(result[4]),
		Audio_feature_id: "",
		Preview_url: result[5],
		Track_number: convertStr(result[6]),
		Popularity: convertStr(result[7]),
		Is_playable: isPlayable,
	})
}

func deleteTrack() {
	track := relational.NewTable[repository.Track](connectionDb)
	forms := createForms(80, 5, "id")
	forms.Draw()
	id := forms.ChooseUser()[0]
	track.Delete(id, repository.Track{})

}