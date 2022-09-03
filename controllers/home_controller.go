package controllers

import (
	"net/http"

	repository "github.com/thompsonmss/minificador-url-golang-with-gorm/repositories/postgres"
	"github.com/thompsonmss/minificador-url-golang-with-gorm/services"
)

func InitHomeController(w http.ResponseWriter, req *http.Request) {

	homeService := services.HomeService{
		Response: w,
		Request:  req,
	}

	homeService.Index()
}

func RegisterHashHomeController(w http.ResponseWriter, req *http.Request) {

	repo := repository.UrlRepository{}

	homeService := services.HomeService{
		Response: w,
		Request:  req,
		RepoUrl:  repo,
	}

	homeService.Register()
}
