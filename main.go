package main

import (
	"github.com/go-playground/validator"
	_ "github.com/go-sql-driver/mysql"
	"github.com/zaenalarifin12/golang-restfull-api/app"
	"github.com/zaenalarifin12/golang-restfull-api/controller"
	"github.com/zaenalarifin12/golang-restfull-api/helper"
	"github.com/zaenalarifin12/golang-restfull-api/middleware"
	"github.com/zaenalarifin12/golang-restfull-api/repository"
	"github.com/zaenalarifin12/golang-restfull-api/service"
	"net/http"
)

func main()  {

	db := app.NewDB()
	validate :=validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr: "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
