package query

import (
	"github.com/lipefebn/spotify-crud/src/core/repository"
	"gorm.io/gorm"
)

type DB struct {
	Db *gorm.DB
}

func (db DB) TopArtistsGenre(genre string, limit int) (result []Result_artist_genre){
	db.Db.Table("r_artist_genre").Select("artists.name, artists.popularity, genres.id").Joins("JOIN artists ON artists.id = r_artist_genre.artist_id").Joins("JOIN genres ON genres.id = r_artist_genre.genre_id").Where("genres.id = ?", genre).Order("artists.popularity DESC").Limit(limit).Scan(&result)
	return
}

func (db DB) ArtistToTracks() (result []Result_artist_track){
	db.Db.Table("r_track_artist").Select("COUNT(track_id) AS cont, artists.name").Group("artist_id").Order("COUNT(track_id) DESC").Limit(10).Joins("JOIN artists ON artists.id = r_track_artist.artist_id").Find(&result)
	return
}

func (db DB) ArtistToAlbums() (result []Result_artist_album){
	db.Db.Table("r_albums_artists").Select("COUNT(album_id) AS amount", "artists.name").Group("artist_id").Order("COUNT(album_id) DESC").Limit(10).Joins("JOIN artists ON artists.id = r_albums_artists.artist_id").Where("artists.name != ?", "Various Artists").Find(&result)
	return
}

func (db DB) TopTracks() (result []Result_top_tracks) {
	db.Db.Table("tracks").Select("name, popularity").Order("popularity DESC").Limit(10).Find(&result)
	return
}

func (db DB) TracksToAlbum(album repository.Album) (result []Result_track_album) {
	db.Db.Table("r_albums_tracks").Select("tracks.name, tracks.popularity").Where("album_id = ?", album.ID).Joins("JOIN tracks on tracks.id = r_albums_tracks.track_id").Order("tracks.popularity DESC").Find(&result)
	return 
}