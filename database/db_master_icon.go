package database

import "time"

// MasterIcon is an entity that directly communicate with the MasterIcon table in the database
type MasterIcon struct {
	ID         uint      `gorm:"primary_key;autoIncrement;not null" json:"id"`
	IconName   string    `gorm:"not null" json:"icon_name"`
	IsActive   bool      `gorm:"not null;default:true" json:"is_active"`
	Created    time.Time `gorm:"type:datetime" json:"created"`
	CreatedBy  string    `json:"created_by"`
	Modified   time.Time `gorm:"type:datetime" json:"modified"`
	ModifiedBy string    `json:"modified_by"`
}

// MasterIconTable set the migrated struct table name
func (masterIcon *MasterIcon) MasterIconTable() string {
	return "dbMasterIcon"
}
