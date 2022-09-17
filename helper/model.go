package helper

import (
	"github.com/rendyuwu/belajar-golang-restfull-api/model/domain"
	"github.com/rendyuwu/belajar-golang-restfull-api/model/web"
)

func ToCategoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}

func ToCategoryResponses(categories []domain.Category) []web.CategoryResponse {
	var categoryResponse []web.CategoryResponse
	for _, category := range categories {
		categoryResponse = append(categoryResponse, ToCategoryResponse(category))
	}
	return categoryResponse
}

func ToProductResponse(product domain.Product) web.ProductResponse {
	return web.ProductResponse{
		Id:          product.Id,
		Name:        product.Name,
		Category:    product.Category,
		Description: product.Description,
		ImageURL:    product.ImageURL,
	}
}

func ToProductResponses(products []domain.Product) []web.ProductResponse {
	var productResponse []web.ProductResponse
	for _, product := range products {
		productResponse = append(productResponse, ToProductResponse(product))
	}
	return productResponse
}
