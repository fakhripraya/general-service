package handlers

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/fakhripraya/general-service/data"
	"github.com/fakhripraya/general-service/entities"
	"github.com/gorilla/mux"
)

// MiddlewareValidateAuth validates the request and calls next if ok
func (chatHandler *ChatHandler) MiddlewareValidateAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

		// Get a session (existing/new)
		session, err := chatHandler.store.Get(r, "session-name")
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			data.ToJSON(&GenericError{Message: err.Error()}, rw)

			return
		}

		// check the token from the session
		// if token available, get the token from the session
		if session.Values["token"] == nil {
			rw.WriteHeader(http.StatusUnauthorized)
			return
		}

		// determine the cookie value that holds the token
		tokenString := session.Values["token"].(string)

		if tokenString != "" {

			// Initialize a new instance of claims
			claims := &data.Claims{}

			token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Error while parsing the token with claims")
				}

				return []byte(data.MySigningKey), nil
			})

			if err != nil {
				if err == jwt.ErrSignatureInvalid {
					rw.WriteHeader(http.StatusUnauthorized)
					data.ToJSON(&GenericError{Message: err.Error()}, rw)

					return
				}

				rw.WriteHeader(http.StatusBadRequest)
				data.ToJSON(&GenericError{Message: err.Error()}, rw)

				return
			}

			if token.Valid {

				// create a new token for the current use, with a renewed expiration time
				expirationTime := time.Now().Add(time.Second * 86400 * 7)
				claims.StandardClaims.ExpiresAt = expirationTime.Unix()
				token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
				tokenString, err := token.SignedString([]byte(data.MySigningKey))

				if err != nil {
					rw.WriteHeader(http.StatusInternalServerError)
					data.ToJSON(&GenericError{Message: err.Error()}, rw)

					return
				}

				// renew the token in the session
				session.Options.MaxAge = 86400 * 7
				session.Values["token"] = tokenString
				session.Values["userLoggedin"] = claims.Username
				session.Save(r, rw)

				next.ServeHTTP(rw, r)
			} else {
				rw.WriteHeader(http.StatusUnauthorized)
				data.ToJSON(&GenericError{Message: "Token invalid"}, rw)

				return
			}
		} else {
			rw.WriteHeader(http.StatusUnauthorized)
			data.ToJSON(&GenericError{Message: "Token invalid"}, rw)

			return
		}
	})
}

// MiddlewareParseUserGetRequest parses the currently logged in user payload from the query parameter
func (chatHandler *ChatHandler) MiddlewareParseUserGetRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		id, err := strconv.ParseUint(vars["user_id"], 10, 32)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			data.ToJSON(&GenericError{Message: "Unable to convert id"}, rw)

			return
		}

		// create the user instance
		user := &entities.UserEntity{
			ID: uint(id),
		}

		// add the user to the context
		ctx := context.WithValue(r.Context(), KeyUser{}, user)
		r = r.WithContext(ctx)

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(rw, r)
	})
}
