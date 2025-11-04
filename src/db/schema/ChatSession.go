package schema

type ChatSession struct {
	Id           uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Participant1 uint   `gorm:"not null" json:"participant1"`
	Participant2 uint   `gorm:"not null" json:"participant2"`
	P1PublicKey  string `gorm:"not null" json:"p1_public_key"`
	P2PublicKey  string `gorm:"not null" json:"p2_public_key"`
	CreatedAt    int64  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    int64  `gorm:"autoUpdateTime" json:"updated_at"`

	User1 User `gorm:"foreignKey:Participant1;references:Id"`
	User2 User `gorm:"foreignKey:Participant2;references:Id"`
}
