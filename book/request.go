package book

type BookRequest struct {
	Title string      `json:"title" binding:"required"`
	Price interface{} `json:"price" binding:"required,number"`
}
