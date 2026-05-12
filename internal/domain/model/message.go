package model

type MessageRequest struct {
	Message string `json:"message"`
	UserID  string `json:"userId"`
}
