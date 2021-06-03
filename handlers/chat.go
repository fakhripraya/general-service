package handlers

import (
	"github.com/fakhripraya/general-service/data"

	"github.com/hashicorp/go-hclog"
	"github.com/srinathgs/mysqlstore"
)

// KeyUser is a key used for the User object in the context
type KeyUser struct{}

// KeyChat is a key used for the User object in the context
type KeyChat struct{}

// ChatHandler is a handler struct for chat changes
type ChatHandler struct {
	logger hclog.Logger
	kost   *data.Chat
	store  *mysqlstore.MySQLStore
}

// NewChatHandler returns a new Chat handler with the given logger
func NewChatHandler(newLogger hclog.Logger, newChat *data.Chat, newStore *mysqlstore.MySQLStore) *ChatHandler {
	return &ChatHandler{newLogger, newChat, newStore}
}

// GenericError is a generic error message returned by a server
type GenericError struct {
	Message string `json:"message"`
}
