package database

import "time"

// DBChatRoom will migrate a db chat room table with the given specification into the database
type DBChatRoom struct {
	ID         uint      `gorm:"primary_key;autoIncrement;not null" json:"id"`
	RoomDesc   string    `gorm:"not null" json:"room_desc"`
	IsActive   bool      `gorm:"not null;default:true" json:"is_active"`
	Created    time.Time `gorm:"type:datetime" json:"created"`
	CreatedBy  string    `json:"created_by"`
	Modified   time.Time `gorm:"type:datetime" json:"modified"`
	ModifiedBy string    `json:"modified_by"`
}

// DBChatRoomMembers will migrate a db chat room members table with the given specification into the database
type DBChatRoomMembers struct {
	ID         uint      `gorm:"primary_key;autoIncrement;not null" json:"id"`
	RoomID     uint      `gorm:"not null" json:"room_id"`
	UserID     uint      `gorm:"not null" json:"user_id"`
	SocketID   uint      `gorm:"not null" json:"socket_id"`
	IsActive   bool      `gorm:"not null;default:true" json:"is_active"`
	Created    time.Time `gorm:"type:datetime" json:"created"`
	CreatedBy  string    `json:"created_by"`
	Modified   time.Time `gorm:"type:datetime" json:"modified"`
	ModifiedBy string    `json:"modified_by"`
}

// DBChatRoomChats will migrate a db chat room chats table with the given specification into the database
type DBChatRoomChats struct {
	ID         uint      `gorm:"primary_key;autoIncrement;not null" json:"id"`
	RoomID     uint      `gorm:"not null" json:"room_id"`
	SenderID   uint      `gorm:"not null" json:"sender_id"`
	ChatBody   string    `json:"chat_body"`
	Attachment byte      `json:"attachment"`
	Pic_URL    string    `json:"pic_url"`
	IsActive   bool      `gorm:"not null;default:true" json:"is_active"`
	Created    time.Time `gorm:"type:datetime" json:"created"`
	CreatedBy  string    `json:"created_by"`
	Modified   time.Time `gorm:"type:datetime" json:"modified"`
	ModifiedBy string    `json:"modified_by"`
}

// DBChatRoomTable set the migrated struct table name
func (dbChatRoom *DBChatRoom) DBChatRoomTable() string {
	return "dbChatRoom"
}

// DBChatRoomMembersTable set the migrated struct table name
func (dbChatRoomMembers *DBChatRoomMembers) DBChatRoomMembersTable() string {
	return "dbChatRoomMembers"
}

// dbChatRoomChatsable set the migrated struct table name
func (dbChatRoomChats *DBChatRoomChats) DBChatRoomChatsTable() string {
	return "dbChatRoomChats"
}
