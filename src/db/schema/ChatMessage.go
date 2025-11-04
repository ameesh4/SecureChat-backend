package schema

type ChatMessage struct {
	Id        uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	SessionId uint   `gorm:"not null" json:"session_id"`
	SenderId  uint   `gorm:"not null" json:"sender_id"`
	Content   []byte `gorm:"not null" json:"content"`
	// MessageType string `gorm:"default:'text'" json:"message_type"`
	IsRead    bool  `gorm:"default:false" json:"is_read"`
	CreatedAt int64 `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt int64 `gorm:"autoUpdateTime" json:"updated_at"`

	// Foreign key relationships
	Session ChatSession `gorm:"foreignKey:SessionId;references:Id"`
	Sender  User        `gorm:"foreignKey:SenderId;references:Id"`
}
