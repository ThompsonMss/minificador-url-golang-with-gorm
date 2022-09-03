package controllers

import (
	"net/http"

	repository "github.com/thompsonmss/minificador-url-golang-with-gorm/repositories/postgres"
	"github.com/thompsonmss/minificador-url-golang-with-gorm/services"
)

func RedirectController(w http.ResponseWriter, req *http.Request) {

	repo := repository.UrlRepository{}
	repoClickUrl := repository.ClickUrlRepository{}

	redirectService := services.RedirectService{
		Response:     w,
		Request:      req,
		Repo:         repo,
		RepoClickUrl: repoClickUrl,
	}

	redirectService.Index()
}
