package service

import (
	"encoding/json"
	"example/go-web-gin/config"
	"example/go-web-gin/model"
	"example/go-web-gin/repositories"
	"log"
	"time"
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
	cacheKey := "album:" + id
	// 1️⃣ Try cache
	cached, err := config.RedisClient.Get(config.Ctx, cacheKey).Result()
	if err == nil {
		log.Println("✅ Cache hit for album ID:", id)
		var album model.Album
		json.Unmarshal([]byte(cached), &album)
		return album, nil
	}
	log.Println("⚠️ Cache miss for album ID:", id)

	log.Println("Fetching album ID", id, "from DB")
	// 2️⃣ Cache miss → DB
	album, err := s.repo.FindByID(id)
	if err != nil {
		return model.Album{}, err
	}

	// 3️⃣ Save to cache
	data, _ := json.Marshal(album)
	err = config.RedisClient.Set(
		config.Ctx,
		cacheKey,
		data,
		60*time.Minute,
	).Err()
	if err != nil {
		log.Println("❌ Failed to set album in cache:", err)
	} else {
		log.Println("✅ Album cached with key:", cacheKey)
	}

	return album, nil
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
	err := s.repo.Delete(id)
	if err == nil {
		s.ClearAlbumCache(id)
	}
	return err

}

// PatchAlbum partially updates an album by ID
func (s *AlbumService) PatchAlbum(id string, album model.Album) (model.Album, error) {
	return s.repo.Patch(id, album)
}

func (s *AlbumService) ClearAlbumCache(id string) {
	cacheKey := "album:" + id
	err := config.RedisClient.Del(config.Ctx, cacheKey).Err()
	if err != nil {
		log.Println("❌ Failed to clear album cache for ID", id, ":", err)
	} else {
		log.Println("✅ Cleared album cache for ID:", id)
	}
}
