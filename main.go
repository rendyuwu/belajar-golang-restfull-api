package main

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"github.com/rendyuwu/belajar-golang-restfull-api/app"
	"github.com/rendyuwu/belajar-golang-restfull-api/controller"
	"github.com/rendyuwu/belajar-golang-restfull-api/helper"
	"github.com/rendyuwu/belajar-golang-restfull-api/repository"
	"github.com/rendyuwu/belajar-golang-restfull-api/service"
)

func main() {
	validate := validator.New()
	db := app.NewDB()

	categoryRepository := repository.NewCategoryRepository()
	productRepository := repository.NewProductRepository()

	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	productService := service.NewProductService(productRepository, db, validate)

	categoryController := controller.NewCategoryController(categoryService)
	productController := controller.NewProductController(productService)

	router := httprouter.New()

	// Route product
	router.GET("/api/v1/products", productController.FindAll)
	router.GET("/api/v1/products/:productId", productController.FindById)
	router.POST("/api/v1/products", productController.Create)
	router.PUT("/api/v1/products/:productId", productController.Update)
	router.DELETE("/api/v1/products/:productId", productController.Delete)

	// Route category
	router.GET("/api/v1/categories", categoryController.FindAll)
	router.GET("/api/v1/categories/:categoryId", categoryController.FindById)
	router.POST("/api/v1/categories", categoryController.Create)
	router.PUT("/api/v1/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/v1/categories/:categoryId", categoryController.Delete)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
