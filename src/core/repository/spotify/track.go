package repository


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

func (t Track) GetId() string{
	return t.ID
}

func (t Track) OrderBy() string { return "popularity" }