package schema

type ChatSession struct {
	Id           uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Participant1 uint   `gorm:"not null" json:"participant1"`
	Participant2 uint   `gorm:"not null" json:"participant2"`
	A1           string `gorm:"not null" json:"a1"` // Encrypted AES key for user 1
	A2           string `gorm:"not null" json:"a2"` // Encrypted AES key for user 2
	CreatedAt    int64  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    int64  `gorm:"autoUpdateTime" json:"updated_at"`

	User1 User `gorm:"foreignKey:Participant1;references:Id"`
	User2 User `gorm:"foreignKey:Participant2;references:Id"`
}
