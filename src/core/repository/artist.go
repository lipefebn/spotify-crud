package repository

import "github.com/lipefebn/spotify-crud/src/adapter/relational"

type ArtistTable relational.TableI[Artist]

type Artist struct {
	ID				string
	Name			string
	Popularity		int
	Followers		int	
}