package query

type Result_artist_genre struct {
	ID string `env:"ID"`
	Name string `env:"NAME"`
	Popularity int `env:"POPULARITY"`
}

type Result_artist_track struct {
	Name string
	Cont int
}

type Result_artist_album struct {
	Name string
	Amount int
}

type Result_top_tracks struct {
	Name string
	Popularity int
}

type Result_track_album struct {
	Name string
	Popularity int
}