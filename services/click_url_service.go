package services

import (
	"net/http"

	"github.com/thompsonmss/minificador-url-golang-with-gorm/repositories"
)

type ClickUrlService struct {
	Response http.ResponseWriter
	Request  *http.Request
	Repo     repositories.ClickUrlRepositoryInterface
}

func (service *ClickUrlService) RegisterClick(idUrl uint) error {
	return nil
}
