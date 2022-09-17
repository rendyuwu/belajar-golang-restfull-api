package web

type ProductCreateRequest struct {
	Name        string `validate:"required,min=1,max=200"`
	Category    string `validate:"required,min=1,max=200"`
	Description string `validate:"required,min=1,max=1000"`
	ImageURL    string `validate:"url,min=1,max=500"`
}
