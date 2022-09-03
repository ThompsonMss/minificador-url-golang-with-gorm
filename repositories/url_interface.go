package repositories

import "github.com/thompsonmss/minificador-url-golang-with-gorm/entities"

type UrlRepositoryInterface interface {
	Store(url entities.Url) (entities.Url, error)
	Delete(id uint) (entities.Url, error)
	Update(url entities.Url) error
	Get(id uint) (entities.Url, error)
	GetAll() ([]entities.Url, error)
	FindHash(hash string) (entities.Url, error)
}
