package repository

import "github.com/lipefebn/spotify-crud/src/adapter/relational"

type AlbumTable relational.TableI[Album]

type Album struct {
	ID  			string
	Name 			string
	Album_group 	string
	Album_type 		string
	Release_date	int64
	Popularity		int
}
