package repositories

import "github.com/thompsonmss/minificador-url-golang-with-gorm/entities"

type ClickUrlRepositoryInterface interface {
	Store(clickUrl entities.ClickUrl) (entities.ClickUrl, error)
}
