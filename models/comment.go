package models

type Comment struct {
	BaseModel
	Content string `json:"content"`

	UserID string `json:"user_id"`
	User   User   `gorm:"foreign_key:UserID" json:"user_id"`

	ArticleID string `gorm:"not null" json:"user_id"`
}
