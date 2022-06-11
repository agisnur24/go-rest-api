package web

type ProductUpdateRequest struct {
	Id         int    `validate:"required"`
	Name       string `validate:"required,min=1,max=100" json:"name"`
	Price      int    `validate:"required" json:"price"`
	CategoryId int    `validate:"required" json:"category_id"`
}
