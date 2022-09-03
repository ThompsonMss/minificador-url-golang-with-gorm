package postgres

import (
	"github.com/thompsonmss/minificador-url-golang-with-gorm/database"
	"github.com/thompsonmss/minificador-url-golang-with-gorm/entities"
)

type ClickUrlRepository struct{}

func (r ClickUrlRepository) Store(clickUrl entities.ClickUrl) (entities.ClickUrl, error) {

	db := database.GetDatabase()
	err := db.Create(&clickUrl).Error

	if err != nil {
		return entities.ClickUrl{}, err
	}

	return clickUrl, nil

}
