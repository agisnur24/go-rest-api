package web

type CustomerCreateRequest struct {
	Name        string `validate:"required,min=1,max=100" json:"name"`
	Address     string `validate:"required,min=1,max=250" json:"address"`
	Email       string `validate:"required,min=1,max=50" json:"email"`
	PhoneNumber string `validate:"required,min=1,max=15" json:"phone_number"`
}
