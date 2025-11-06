package model

type PublicKeys struct {
	User1PublicKey string `json:"user1_public_key"`
	User2PublicKey string `json:"user2_public_key"`
}

type CreateSessionRequest struct {
	Email string `json:"email"`
	A1    string `json:"a1"`
	A2    string `json:"a2"`
}

type RequestSessionRequest struct {
	Email string `json:"email"`
}
