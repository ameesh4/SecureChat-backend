package schema

type User struct {
	Id          uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string `gorm:"not null" json:"name"`
	PhoneNumber string `gorm:"unique" json:"phone_number"`
	Email       string `gorm:"unique" json:"email"`
	Password    string `gorm:"not null" json:"password"`
	CreatedAt   string `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   string `gorm:"autoUpdateTime" json:"updated_at"`
}
