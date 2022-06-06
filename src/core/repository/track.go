package repository

import "github.com/lipefebn/spotify-crud/src/adapter/relational"

type TrackTable relational.TableI[Track]

type Track struct {
	ID 					string
	Disc_number 		int
	Duration 			int
	Explicit 			int
	Audio_feature_id 	string
	Name 				string
	Preview_url 		string
	Track_number 		int
	Popularity 			int
	Is_playable 		bool
}
