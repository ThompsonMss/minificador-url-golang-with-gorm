package postgres

import (
	"github.com/thompsonmss/minificador-url-golang-with-gorm/database"
	"github.com/thompsonmss/minificador-url-golang-with-gorm/entities"
)

type UrlRepository struct{}

func (r UrlRepository) Store(url entities.Url) (entities.Url, error) {

	db := database.GetDatabase()
	err := db.Create(&url).Error

	if err != nil {
		return entities.Url{}, err
	}

	return url, nil

}

func (r UrlRepository) Delete(id uint) (entities.Url, error) {
	return entities.Url{}, nil
}

func (r UrlRepository) Update(url entities.Url) error {
	return nil
}

func (r UrlRepository) Get(id uint) (entities.Url, error) {
	return entities.Url{}, nil
}

func (r UrlRepository) GetAll() ([]entities.Url, error) {
	return []entities.Url{}, nil
}

func (r UrlRepository) FindHash(hash string) (entities.Url, error) {
	db := database.GetDatabase()

	urlModel := entities.Url{}
	db.First(&urlModel, "hash = ?", hash)

	return urlModel, nil
}
