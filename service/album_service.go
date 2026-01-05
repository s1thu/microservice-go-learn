package service

import (
	"example/go-web-gin/model"
	"example/go-web-gin/repositories"
)

type AlbumService struct {
	repo repositories.AlbumRepository
}

func NewAlbumService(repo repositories.AlbumRepository) *AlbumService {
	return &AlbumService{repo: repo}
}

// GetAllAlbums retrieves all albums from the repository
func (s *AlbumService) GetAllAlbums() ([]model.Album, error) {
	return s.repo.FindAll()
}

// GetAlbumByID retrieves a single album by its ID
func (s *AlbumService) GetAlbumByID(id string) (model.Album, error) {
	return s.repo.FindByID(id)
}

// CreateAlbum creates a new album
func (s *AlbumService) CreateAlbum(album model.Album) (model.Album, error) {
	return s.repo.Create(album)
}

// UpdateAlbum updates an existing album by ID
func (s *AlbumService) UpdateAlbum(id string, album model.Album) (model.Album, error) {
	return s.repo.Update(id, album)
}

// DeleteAlbum deletes an album by ID
func (s *AlbumService) DeleteAlbum(id string) error {
	return s.repo.Delete(id)
}

// PatchAlbum partially updates an album by ID
func (s *AlbumService) PatchAlbum(id string, album model.Album) (model.Album, error) {
	return s.repo.Patch(id, album)
}
