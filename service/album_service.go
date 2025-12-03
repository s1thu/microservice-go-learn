package service

import (
	"errors"
	"example/go-web-gin/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// @BasePath /api/v1

// GetAllAlbums godoc
// @Summary List albums
// @Description get all albums
// @Tags album
// @Produce json
// @Success 200 {array} model.Album
// @Router /albums [get]
func GetAllAlbums(c *gin.Context) {
	c.IndentedJSON(200, model.Albums)
}

// GetAllAlbums godoc
// @Summary List albums
// @Description get all albums
// @Tags album
// @Produce json
// @Success 200 {array} model.Album
// @Router /albums [get]
func GetAllAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range model.Albums {
		if a.ID == id {
			c.IndentedJSON(200, a)
			return
		}
	}
	c.IndentedJSON(404, gin.H{"message": "album not found"})
}

// PostAlbum godoc
// @Summary Create album
// @Description create a new album
// @Tags album
// @Accept json
// @Produce json
// @Param album body model.Album true "Album payload"
// @Success 201 {object} model.Album
// @Router /albums [post]
func PostAlbum(c *gin.Context) {
	var newAlbum model.Album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.ShouldBindJSON(&newAlbum); err != nil {
		// Extract validation errors
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]string, len(ve))
			for i, fe := range ve {
				out[i] = fmt.Sprintf("field %s failed on %s rule", fe.Field(), fe.Tag())
			}
			c.JSON(400, gin.H{"errors": out})
			return
		}

		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Add the new album to the slice.
	model.Albums = append(model.Albums, newAlbum)
	c.IndentedJSON(201, newAlbum)
}

// UpdateAlbum godoc
// @Summary Update album by ID
// @Description update an existing album
// @Tags album
// @Accept json
// @Produce json
// @Param id path string true "Album ID"
// @Param album body model.Album true "Updated album"
// @Success 200 {object} model.Album
// @Failure 404 {object} map[string]string
// @Router /albums/{id} [put]
func UpdateAlbum(c *gin.Context) {
	var updatedAlbum model.Album

	if err := c.BindJSON(&updatedAlbum); err != nil {
		return
	}

	id := c.Param("id")

	for i, a := range model.Albums {
		if a.ID == id {
			model.Albums[i] = updatedAlbum
			c.IndentedJSON(200, updatedAlbum)
			return
		}
	}
	c.IndentedJSON(404, gin.H{"message": "album not found"})
}

// DeleteAlbum godoc
// @Summary Delete album by ID
// @Tags album
// @Param id path string true "Album ID"
// @Success 200 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /albums/{id} [delete]
func DeleteAlbum(c *gin.Context) {
	id := c.Param("id")

	for i, a := range model.Albums {
		if a.ID == id {
			// Remove the album from the slice
			model.Albums = append(model.Albums[:i], model.Albums[i+1:]...)
			c.IndentedJSON(200, gin.H{"message": "album deleted"})
			return
		}
	}
	c.IndentedJSON(404, gin.H{"message": "album not found"})
}

// PatchAlbum godoc
// @Summary Patch album fields
// @Description update one or more fields of an album
// @Tags album
// @Accept json
// @Produce json
// @Param id path string true "Album ID"
// @Param album body map[string]interface{} true "Fields to update"
// @Success 200 {object} model.Album
// @Failure 404 {object} map[string]string
// @Router /albums/{id} [patch]
func PatchAlbum(c *gin.Context) {
	id := c.Param("id")

	var albumUpdates map[string]interface{}
	if err := c.BindJSON(&albumUpdates); err != nil {
		return
	}

	for i, a := range model.Albums {
		if a.ID == id {
			// Apply updates to the album
			if title, ok := albumUpdates["title"].(string); ok {
				model.Albums[i].Title = title
			}
			if artist, ok := albumUpdates["artist"].(string); ok {
				model.Albums[i].Artist = artist
			}
			if price, ok := albumUpdates["price"].(float64); ok {
				model.Albums[i].Price = price
			}
			c.IndentedJSON(200, model.Albums[i])
			return
		}
	}
	c.IndentedJSON(404, gin.H{"message": "album not found"})
}

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /example/helloworld [get]
func Helloworld(g *gin.Context) {
	g.JSON(http.StatusOK, "helloworld")
}
