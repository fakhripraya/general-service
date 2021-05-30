package entities

import (
	"time"
)

// ChatRoom is an entity to communicate with the chat room client side
type ChatRoom struct {
	ID         uint      `json:"id"`
	RoomDesc   string    `json:"room_desc"`
	IsActive   bool      `json:"is_active"`
	Created    time.Time `json:"created"`
	CreatedBy  string    `json:"created_by"`
	Modified   time.Time `json:"modified"`
	ModifiedBy string    `json:"modified_by"`
}

// ChatRoomMembers is an entity to communicate with the chat room members client side
type ChatRoomMembers struct {
	ID             uint      `json:"id"`
	RoomID         uint      `json:"room_id"`
	UserID         uint      `json:"user_id"`
	Displayname    string    `json:"displayname"`
	ProfilePicture string    `json:"profile_picture"`
	SocketID       uint      `json:"socket_id"`
	IsActive       bool      `json:"is_active"`
	Created        time.Time `json:"created"`
	CreatedBy      string    `json:"created_by"`
	Modified       time.Time `json:"modified"`
	ModifiedBy     string    `json:"modified_by"`
}

// ChatRoomChats is an entity to communicate with the chat room chats client side
type ChatRoomChats struct {
	ID         uint      `json:"id"`
	RoomID     uint      `json:"room_id"`
	SenderID   uint      `json:"sender_id"`
	ChatBody   string    `json:"chat_body"`
	Attachment byte      `json:"attachment"`
	Pic_URL    string    `json:"pic_url"`
	IsActive   bool      `json:"is_active"`
	Created    time.Time `json:"created"`
	CreatedBy  string    `json:"created_by"`
	Modified   time.Time `json:"modified"`
	ModifiedBy string    `json:"modified_by"`
}
