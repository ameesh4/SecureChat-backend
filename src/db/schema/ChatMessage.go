package schema

type ChatMessage struct {
	Id         uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	SenderId   uint   `gorm:"not null" json:"sender_id"`
	ReceiverId uint   `gorm:"not null" json:"receiver_id"`
	Content    string `gorm:"not null" json:"content"`
	Iv         string `gorm:"not null" json:"iv"`
	SessionId  uint   `gorm:"not null" json:"session_id"`
	// MessageType string `gorm:"default:'text'" json:"message_type"`
	IsRead    bool  `gorm:"default:false" json:"is_read"`
	CreatedAt int64 `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt int64 `gorm:"autoUpdateTime" json:"updated_at"`

	// Foreign key relationships
	Sender   User `gorm:"foreignKey:SenderId;references:Id"`
	Receiver User `gorm:"foreignKey:ReceiverId;references:Id"`
}
