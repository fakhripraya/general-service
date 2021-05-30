package data

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/fakhripraya/general-service/config"
	"github.com/fakhripraya/general-service/database"
	"github.com/hashicorp/go-hclog"
	"github.com/srinathgs/mysqlstore"
)

// Claims determine the current user token holder
type Claims struct {
	Username string
	jwt.StandardClaims
}

// Chat defines a struct for chat flow
type Chat struct {
	logger hclog.Logger
}

// NewChat is a function to create new Chat struct
func NewChat(newLogger hclog.Logger) *Chat {
	return &Chat{newLogger}
}

// GetCurrentUser will get the current user login info
func (kost *Chat) GetCurrentUser(rw http.ResponseWriter, r *http.Request, store *mysqlstore.MySQLStore) (*database.MasterUser, error) {

	// Get a session (existing/new)
	session, err := store.Get(r, "session-name")
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)

		return nil, err
	}

	// check the logged in user from the session
	// if user available, get the user info from the session
	if session.Values["userLoggedin"] == nil {
		rw.WriteHeader(http.StatusUnauthorized)

		return nil, fmt.Errorf("Error 401")
	}

	// work with database
	// look for the current user logged in in the db
	var currentUser database.MasterUser
	if err := config.DB.Where("username = ?", session.Values["userLoggedin"].(string)).First(&currentUser).Error; err != nil {
		rw.WriteHeader(http.StatusUnauthorized)

		return nil, err
	}

	return &currentUser, nil

}
