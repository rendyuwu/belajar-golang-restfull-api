package web

type ProductUpdateRequest struct {
	Id          int    `validate:"required" json:"id"`
	Name        string `validate:"required,min=1,max=200" json:"name"`
	Category    string `validate:"required,min=1,max=200" json:"category"`
	Description string `validate:"required,min=1,max=1000" json:"description"`
	ImageURL    string `validate:"url,min=1,max=500" json:"image_url"`
}
