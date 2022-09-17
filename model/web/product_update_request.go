package web

type ProductUpdateRequest struct {
	Id                                    int
	Name, Category, Description, ImageURL string
}
