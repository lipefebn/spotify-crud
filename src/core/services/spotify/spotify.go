package services

import (
	"encoding/json"
	"log"

	"github.com/lipefebn/spotify-crud/src/adapter/relational"
	"github.com/lipefebn/spotify-crud/src/core/repository/spotify"
	"github.com/lipefebn/spotify-crud/src/infra/config"
)

type Model repository.SpotifyRepository

type ServiceDTO[M Model] struct {
	table relational.Table[M]
	data  M
}

func initDTO[typeI Model](data []byte) ServiceDTO[typeI] {
	dataParsed := parser[typeI](data)
	return ServiceDTO[typeI]{
		table: *relational.NewTable[typeI](config.Connect()),
		data: dataParsed,
	}
}

func parser[typeI Model](data []byte) (result typeI){
	if err := json.Unmarshal([]byte(data), &result); err != nil {
		//TODO
		log.Fatal(err, "in parser")
	}
	return
}

func (s ServiceDTO[M]) Insert() error {
	if err := s.table.Create(s.data); err != nil {
		//TODO
		log.Fatal(err, "in insert")
	}
	return nil
}

func (s ServiceDTO[M]) Update() error {
	if err := s.table.Update(s.data.GetId(), s.data); err != nil {
		//TODO
		log.Fatal(err, "in Update")
	}
	return nil
}

func (s ServiceDTO[M]) Delete() error {
	if err := s.table.Delete(s.data.GetId(), s.data); err != nil {
		//TODO
		log.Fatal(err, "in Delete")
	}
	return nil
}