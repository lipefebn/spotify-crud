package repository

import "github.com/lipefebn/spotify-crud/src/adapter/relational"

type GenreTable relational.TableI[Genre]

type Genre struct {
	ID string
}
