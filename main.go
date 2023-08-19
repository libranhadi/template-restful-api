package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/go-playground/validator/v10"
	"template-restful-api/repository"
	"template-restful-api/service"
	"template-restful-api/controller"
	"template-restful-api/app"
	// "template-restful-api/helper"
	"template-restful-api/exception"
	"template-restful-api/middleware"
	"net/http"
)

func main() {
	validate := validator.New();
	db := app.NewDb()
	categoryRepository := repository.NewCategoryRepository();
	categoryService := service.NewCategoryService(categoryRepository, db, validate);
	categoryController := controller.NewCategoryController(categoryService);

	router := httprouter.New()

	router.GET("/api/categories", categoryController.Get)
	router.POST("/api/categories", categoryController.Create)
	router.GET("/api/category/:id", categoryController.FindById)
	router.PUT("/api/category/:id", categoryController.Update)
	router.DELETE("/api/category/:id", categoryController.Delete)
	router.PanicHandler = exception.ErrorHandler
	server := http.Server{
		Addr : "localhost:8000",
		Handler : middleware.NewSimpleAuthMiddleware(router),
	}

	server.ListenAndServe()
	// helper.PanicIfError(err);
}