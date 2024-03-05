package models

// User used to get the invite information
type Invitation struct {
	EventID uint   `json:"event_id"`
	UserID  uint   `json:"user_id"`
	Message string `json:"message"`
} //name @Invitation

type AcceptRequest struct {
	EventID uint `json:"event_id" binding:"required"`
	UserID  uint `json:"user_id" binding:"required"`
} //name @AcceptRequest

type OutPutOfInvite struct {
	Message string `json:"message"`
} //name @OutPutOfInvite
