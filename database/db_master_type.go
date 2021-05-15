package database

import "time"

// MasterKostType is an entity that directly communicate with the MasterKostType table in the database
// kost type adalah tipe sewaan
type MasterKostType struct {
	ID         uint      `gorm:"primary_key;autoIncrement;not null" json:"id"`
	TypeDesc   string    `gorm:"not null" json:"type_desc"`
	IsActive   bool      `gorm:"not null;default:true" json:"is_active"`
	Created    time.Time `gorm:"type:datetime" json:"created"`
	CreatedBy  string    `json:"created_by"`
	Modified   time.Time `gorm:"type:datetime" json:"modified"`
	ModifiedBy string    `json:"modified_by"`
}

// MasterKostTypeTable set the migrated struct table name
func (masterKostType *MasterKostType) MasterKostTypeTable() string {
	return "dbMasterKostType"
}
