package database

import "time"

// MasterEvent will migrate a master event table with the given specification into the database
type MasterEvent struct {
	ID         uint      `gorm:"primary_key;autoIncrement;not null" json:"id"`
	EventName  string    `gorm:"not null" json:"event_name"`
	Thumbnail  string    `json:"thumbnail"`
	IsActive   bool      `gorm:"not null;default:true" json:"is_active"`
	Created    time.Time `gorm:"type:datetime" json:"created"`
	CreatedBy  string    `json:"created_by"`
	Modified   time.Time `gorm:"type:datetime" json:"modified"`
	ModifiedBy string    `json:"modified_by"`
}

// MasterEventDetail will migrate a master event detail table with the given specification into the database
type MasterEventDetail struct {
	ID         uint      `gorm:"primary_key;autoIncrement;not null" json:"id"`
	EventID    uint      `gorm:"not null" json:"event_id"`
	KostID     uint      `gorm:"not null" json:"kost_id"`
	IsActive   bool      `gorm:"not null;default:true" json:"is_active"`
	Created    time.Time `gorm:"type:datetime" json:"created"`
	CreatedBy  string    `json:"created_by"`
	Modified   time.Time `gorm:"type:datetime" json:"modified"`
	ModifiedBy string    `json:"modified_by"`
}

// MasterEventTable set the migrated struct table name
func (masterEvent *MasterEvent) MasterEventTable() string {
	return "dbMasterEvent"
}

// MasterEventDetailTable set the migrated struct table name
func (masterEventDetail *MasterEventDetail) MasterEventDetailTable() string {
	return "dbMasterEventDetail"
}
