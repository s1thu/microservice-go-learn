package model

// album represents data about a record album.
type Album struct {
	ID     string  `json:"id" example:"1"`
	Title  string  `json:"title" binding:"required,min=1,max=50" example:"My Album"`
	Artist string  `json:"artist" binding:"required" example:"John Doe"`
	Price  float64 `json:"price" binding:"required,gt=0" example:"19.99"`
}
