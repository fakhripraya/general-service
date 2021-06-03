package handlers

import (
	"net/http"

	"github.com/fakhripraya/general-service/config"
	"github.com/fakhripraya/general-service/data"
	"github.com/fakhripraya/general-service/database"
	"github.com/fakhripraya/general-service/entities"
)

// PatchReadChat is a method to patch the read value of the chat based on the given parameter info
func (chatHandler *ChatHandler) PatchReadChat(rw http.ResponseWriter, r *http.Request) {

	// get the user via context
	user := r.Context().Value(KeyUser{}).(*entities.UserEntity)

	// get the requestChatRoom via context
	requestChatRoom := r.Context().Value(KeyUser{}).(*entities.RequestChatRoom)

	var allChatRoomChats []database.DBChatRoomChats
	if err := config.DB.Where("room_id = ? && user_id != ? && chat_read = ?", requestChatRoom.ChatRoom.ID, user.ID, false).Find(&allChatRoomChats).Error; err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)

		return
	}

	return
}
