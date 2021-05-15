package database

import "time"

// MasterUOM is an entity that directly communicate with the MasterUOM table in the database
type MasterUOM struct {
	ID         uint      `gorm:"primary_key;autoIncrement;not null" json:"id"`
	UOMType    string    `gorm:"not null" json:"uom_type"`
	UOMDesc    string    `gorm:"not null" json:"uom_desc"`
	UOMRate    float64   `gorm:"not null" json:"uom_rate"`
	IsActive   bool      `gorm:"not null;default:true" json:"is_active"`
	Created    time.Time `gorm:"type:datetime" json:"created"`
	CreatedBy  string    `json:"created_by"`
	Modified   time.Time `gorm:"type:datetime" json:"modified"`
	ModifiedBy string    `json:"modified_by"`
}

// MasterUOMTable set the migrated struct table name
func (masterUOM *MasterUOM) MasterUOMTable() string {
	return "dbMasterUOM"
}
