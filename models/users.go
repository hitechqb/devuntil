package models

type User struct {
	BaseModel

	Email       string `gorm:"not null; unique" json:"email"`
	PhoneNumber string `gorm:"not null; unique" json:"phone_number"`

	DisplayName string `gorm:"not null" json:"user_name"`

	PhotoID int   `json:"photo_id"`
	Photo   Photo `gorm:"foreign_key:PhotoID" json:"photo"`

	Active int `json:"active"` // 0 - Pending | 1 - Active | 2 - Blocked
}
