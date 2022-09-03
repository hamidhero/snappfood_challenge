package forms

type AddOrder struct {
	Id    int64  `json:"id" binding:"required"`
	Price int64  `json:"price" binding:"required"`
	Title string `json:"title" binding:"required"`
}
