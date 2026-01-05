package repositories

import (
	"database/sql"
	"example/go-web-gin/model"
)

// repo interface
type AlbumRepository interface {
	FindAll() ([]model.Album, error)
	FindByID(id string) (model.Album, error)
	Create(album model.Album) (model.Album, error)
	Update(id string, album model.Album) (model.Album, error)
	Delete(id string) error
	Patch(id string, album model.Album) (model.Album, error)
}

// implementation
type AlbumRepoImpl struct {
	db *sql.DB
}

// Create implements AlbumRepository.
func (a *AlbumRepoImpl) Create(album model.Album) (model.Album, error) {
	query := `
		INSERT INTO albums (title, artist, price)
		VALUES ($1, $2, $3)
		RETURNING id
	`

	err := a.db.QueryRow(
		query,
		album.Title,
		album.Artist,
		album.Price,
	).Scan(&album.ID)

	return album, err
}

// Delete implements AlbumRepository.
func (a *AlbumRepoImpl) Delete(id string) error {
	query := `DELETE FROM albums WHERE id = $1`

	_, err := a.db.Exec(query, id)
	return err
}

// FindAll implements AlbumRepository.
func (a *AlbumRepoImpl) FindAll() ([]model.Album, error) {
	query := `
		SELECT id, title, artist, price
		FROM albums
	`

	rows, err := a.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var albums []model.Album

	for rows.Next() {
		var album model.Album
		if err := rows.Scan(
			&album.ID,
			&album.Title,
			&album.Artist,
			&album.Price,
		); err != nil {
			return nil, err
		}

		albums = append(albums, album)
	}

	return albums, nil
}

// FindByID implements AlbumRepository.
func (a *AlbumRepoImpl) FindByID(id string) (model.Album, error) {
	query := `
		SELECT id, title, artist, price
		FROM albums
		WHERE id = $1
	`

	var album model.Album

	err := a.db.QueryRow(query, id).Scan(
		&album.ID,
		&album.Title,
		&album.Artist,
		&album.Price,
	)

	if err == sql.ErrNoRows {
		return model.Album{}, err
	}

	return album, err
}

// Update implements AlbumRepository.
func (a *AlbumRepoImpl) Update(id string, album model.Album) (model.Album, error) {
	query := `
		UPDATE albums
		SET title = $1,
		    artist = $2,
		    price = $3
		WHERE id = $4
		RETURNING id, title, artist, price
	`

	err := a.db.QueryRow(
		query,
		album.Title,
		album.Artist,
		album.Price,
		id,
	).Scan(
		&album.ID,
		&album.Title,
		&album.Artist,
		&album.Price,
	)

	return album, err
}

// Patch implements AlbumRepository.
func (a *AlbumRepoImpl) Patch(id string, album model.Album) (model.Album, error) {
	query := `
		UPDATE albums
		SET
			title  = COALESCE(NULLIF($1, ''), title),
			artist = COALESCE(NULLIF($2, ''), artist),
			price  = COALESCE(NULLIF($3, 0), price)
		WHERE id = $4
		RETURNING id, title, artist, price
	`

	err := a.db.QueryRow(
		query,
		album.Title,
		album.Artist,
		album.Price,
		id,
	).Scan(
		&album.ID,
		&album.Title,
		&album.Artist,
		&album.Price,
	)

	return album, err
}
func NewAlbumRepoImpl(db *sql.DB) AlbumRepository {
	return &AlbumRepoImpl{db: db}
}
