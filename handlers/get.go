package handlers

import (
	"net/http"

	"github.com/fakhripraya/general-service/config"
	"github.com/fakhripraya/general-service/data"
	"github.com/fakhripraya/general-service/database"
	"github.com/fakhripraya/general-service/entities"
)

// GetRoomList is a method to fetch all room list based on the given parameter info
func (chatHandler *ChatHandler) GetRoomList(rw http.ResponseWriter, r *http.Request) {

	// get the user via context
	user := r.Context().Value(KeyUser{}).(*entities.UserEntity)

	// look for the selected kost in the db to fetch all the picts

	var allChatRoomMembers []database.DBChatRoomMembers
	if err := config.DB.Where("user_id = ?", user.ID).Find(&allChatRoomMembers).Error; err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)

		return
	}

	type CustomGetRoomType struct {
		ChatRoom         database.DBChatRoom        `json:"chat_room"`
		ChatRoomLastChat database.DBChatRoomChats   `json:"chat_room_last_chat"`
		ChatRoomMembers  []entities.ChatRoomMembers `json:"chat_room_members"`
	}

	var finalResult []CustomGetRoomType
	for _, member := range allChatRoomMembers {

		if member.UserID == user.ID {

			var chatRoom database.DBChatRoom
			if err := config.DB.Where("id = ?", member.RoomID).First(&chatRoom).Error; err != nil {
				rw.WriteHeader(http.StatusBadRequest)
				data.ToJSON(&GenericError{Message: err.Error()}, rw)

				return
			}

			var chatRoomLastChat database.DBChatRoomChats
			if err := config.DB.Where("room_id = ?", chatRoom.ID).Last(&chatRoomLastChat).Error; err != nil {
				rw.WriteHeader(http.StatusBadRequest)
				data.ToJSON(&GenericError{Message: err.Error()}, rw)

				return
			}

			// var chatRoomMembers []database.DBChatRoomMembers
			// if err := config.DB.Where("room_id = ? && user_id = ?", chatRoom.ID, user.ID).Find(&chatRoomMembers).Error; err != nil {
			// 	rw.WriteHeader(http.StatusBadRequest)
			// 	data.ToJSON(&GenericError{Message: err.Error()}, rw)

			// 	return
			// }

			var chatRoomMembers []entities.ChatRoomMembers
			if err := config.DB.
				Model(&database.DBChatRoomMembers{}).
				Select("db_chat_room_members.id,db_chat_room_members.room_id,db_chat_room_members.user_id,db_chat_room_members.socket_id,master_users.display_name as displayname,master_users.profile_picture").
				Joins("inner join master_users on master_users.id = db_chat_room_members.user_id").
				Where("db_chat_room_members.room_id = ?", chatRoom.ID).
				Scan(&chatRoomMembers).Error; err != nil {
				rw.WriteHeader(http.StatusBadRequest)
				data.ToJSON(&GenericError{Message: err.Error()}, rw)

				return
			}

			finalResult = append(finalResult, CustomGetRoomType{
				ChatRoom:         chatRoom,
				ChatRoomLastChat: chatRoomLastChat,
				ChatRoomMembers:  chatRoomMembers,
			})

		}

	}

	// parse the given instance to the response writer
	err := data.ToJSON(finalResult, rw)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)

		return
	}

	return
}
