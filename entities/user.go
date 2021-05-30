package entities

import (
	"time"
)

// UserEntity is an entity to communicate with the user entity on the client side
type UserEntity struct {
	ID             uint      `json:"id"`
	RoleID         uint      `json:"role_id"`
	Username       string    `json:"username"`
	DisplayName    string    `json:"displayname"`
	Password       []byte    `json:"password"`
	Email          string    `json:"email"`
	Phone          string    `json:"phone"`
	ProfilePicture string    `json:"profile_picture"`
	Country        string    `json:"country"`
	City           string    `json:"city"`
	Address        string    `json:"address"`
	Latitude       string    `json:"latitude"`
	Longitude      string    `json:"longitude"`
	LoginFailCount uint      `gorm:"default:0"`
	IsVerified     bool      `gorm:"not null;default:false" json:"is_verified"`
	IsActive       bool      `gorm:"not null;default:true" json:"is_active"`
	Created        time.Time `gorm:"type:datetime" json:"created"`
	CreatedBy      string    `json:"created_by"`
	Modified       time.Time `gorm:"type:datetime" json:"modified"`
	ModifiedBy     string    `json:"modified_by"`
}
