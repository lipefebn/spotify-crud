package main

import (
	"fmt"

	"github.com/lipefebn/spotify-crud/src/adapter/relational"
	"github.com/lipefebn/spotify-crud/src/core/repository"
	"github.com/lipefebn/spotify-crud/src/infra/config"
)

func main() {
	sla := relational.NewTable[repository.Artist](config.Connect("..."))
	result := sla.Select("name = ?", "Eminem")
	fmt.Println(result)
}