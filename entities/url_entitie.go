package entities

type Url struct {
	ID       uint   `gorm:"primaryKey"`
	Original string `gorm:"type:varchar(255);not null"`
	Hash     string `gorm:"type:varchar(255);not null;unique"`
	Base
}

func NewUrl(original string, hash string) Url {
	return Url{
		Original: original,
		Hash:     hash,
	}
}
