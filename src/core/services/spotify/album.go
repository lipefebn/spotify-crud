package services

import (
	repository "github.com/lipefebn/spotify-crud/src/core/repository/spotify"
)

type AlbumService struct {
	ServiceDTO[repository.Album]
}

func InitAlbumService(data []byte) AlbumService {
	return AlbumService{ initDTO[repository.Album](data) }
}

func (a AlbumService) Select() []repository.Album{
	return a.table.Select("popularity DESC", "name LIKE ?", a.data.Name+"%")
}
