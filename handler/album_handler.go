package handler

import (
	"database/sql"
	"net/http"

	"example/go-web-gin/model"
	"example/go-web-gin/service"

	"github.com/gin-gonic/gin"
)

type AlbumHandler struct {
	service *service.AlbumService
}

func NewAlbumHandler(service *service.AlbumService) *AlbumHandler {
	return &AlbumHandler{service: service}
}

// GetAllAlbums godoc
// @Summary Get all albums
// @Description Retrieves a list of all albums
// @Tags albums
// @Accept json
// @Produce json
// @Success 200 {array} model.Album
// @Failure 500 {object} map[string]string
// @Router /albums [get]
func (h *AlbumHandler) GetAllAlbums(c *gin.Context) {
	albums, err := h.service.GetAllAlbums()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if albums == nil {
		albums = []model.Album{}
	}

	c.JSON(http.StatusOK, albums)
}

// GetAlbumByID godoc
// @Summary Get album by ID
// @Description Retrieves an album by its ID
// @Tags albums
// @Accept json
// @Produce json
// @Param id path string true "Album ID"
// @Success 200 {object} model.Album
// @Failure 404 {object} map[string]string
// @Router /albums/{id} [get]
func (h *AlbumHandler) GetAlbumByID(c *gin.Context) {
	id := c.Param("id")

	album, err := h.service.GetAlbumByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "album not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, album)
}

// PostAlbum godoc
// @Summary Create a new album
// @Description Creates a new album with the provided data
// @Tags albums
// @Accept json
// @Produce json
// @Param album body model.Album true "Album data"
// @Success 201 {object} model.Album
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /albums [post]
func (h *AlbumHandler) PostAlbum(c *gin.Context) {
	var newAlbum model.Album
	if err := c.ShouldBindJSON(&newAlbum); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	album, err := h.service.CreateAlbum(newAlbum)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, album)
}

// UpdateAlbum godoc
// @Summary Update an album
// @Description Updates an existing album by ID
// @Tags albums
// @Accept json
// @Produce json
// @Param id path string true "Album ID"
// @Param album body model.Album true "Album data"
// @Success 200 {object} model.Album
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /albums/{id} [put]
func (h *AlbumHandler) UpdateAlbum(c *gin.Context) {
	id := c.Param("id")

	var updatedAlbum model.Album
	if err := c.ShouldBindJSON(&updatedAlbum); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	album, err := h.service.UpdateAlbum(id, updatedAlbum)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "album not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, album)
}

// DeleteAlbum godoc
// @Summary Delete an album
// @Description Deletes an album by ID
// @Tags albums
// @Accept json
// @Produce json
// @Param id path string true "Album ID"
// @Success 200 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /albums/{id} [delete]
func (h *AlbumHandler) DeleteAlbum(c *gin.Context) {
	id := c.Param("id")

	err := h.service.DeleteAlbum(id)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "album not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "album deleted successfully"})
}

// PatchAlbum godoc
// @Summary Partially update an album
// @Description Partially updates an existing album by ID
// @Tags albums
// @Accept json
// @Produce json
// @Param id path string true "Album ID"
// @Param album body model.Album true "Partial album data"
// @Success 200 {object} model.Album
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /albums/{id} [patch]
func (h *AlbumHandler) PatchAlbum(c *gin.Context) {
	id := c.Param("id")

	var patchAlbum model.Album
	if err := c.ShouldBindJSON(&patchAlbum); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	album, err := h.service.PatchAlbum(id, patchAlbum)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "album not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, album)
}
