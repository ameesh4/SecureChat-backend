package schema

type User struct {
	Id           uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name         string `gorm:"not null" json:"name"`
	PhoneNumber  string `gorm:"unique" json:"phone_number"`
	Email        string `gorm:"unique" json:"email"`
	Password     string `gorm:"not null" json:"password"`
	RefreshToken string `gorm:"" json:"refresh_token"`
	IsAdmin      bool   `gorm:"default:false" json:"is_admin"`
	CreatedAt    int64  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    int64  `gorm:"autoUpdateTime" json:"updated_at"`
}
