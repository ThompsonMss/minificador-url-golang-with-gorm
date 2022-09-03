package entities

type ClickUrl struct {
	ID    uint   `gorm:"primaryKey"`
	IP    string `gorm:"type:varchar(255)"`
	UrlID uint
	Url   Url
	Base
}

func (c *ClickUrl) newClickUrl(idUrl uint, ip string) *ClickUrl {
	return &ClickUrl{
		UrlID: idUrl,
		IP:    ip,
	}
}
