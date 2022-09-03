package migrations

import (
	"github.com/thompsonmss/minificador-url-golang-with-gorm/entities"
	"gorm.io/gorm"
)

func RunAutoMigrations(db *gorm.DB) {

	db.AutoMigrate(&entities.User{})
	db.AutoMigrate(&entities.Url{})
	db.AutoMigrate(&entities.ClickUrl{})

}
