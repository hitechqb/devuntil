package models

type Photo struct {
	BaseModel
	URL  string `gorm:"not null" json:"url"`
	Path string `json:"-"`
}
