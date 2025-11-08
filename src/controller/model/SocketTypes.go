package model

type AuthRequest struct {
	Token string `json:"token"`
}

type Message struct {
	SenderId   uint   `json:"sender_id"`
	ReceiverId uint   `json:"receiver_id"`
	Content    string `json:"content"`
	Iv         string `json:"iv"`
}
