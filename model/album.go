package model

// album represents data about a record album.
type Album struct {
	ID     string  `json:"id" example:"1"`
	Title  string  `json:"title" binding:"required,min=1,max=50" example:"My Album"`
	Artist string  `json:"artist" binding:"required" example:"John Doe"`
	Price  float64 `json:"price" binding:"required,gt=0" example:"19.99"`
}

// albums slice to seed record album data.
var Albums = []Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}
