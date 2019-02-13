package models

type Article struct {
	BaseModel

	Slug    string `gorm:"unique; not null" json:"slug"`
	Title   string `gorm:"not null" json:"title"`
	Content string `gorm:"not null" json:"content"`

	UserId int `json:"user_id"`
}
