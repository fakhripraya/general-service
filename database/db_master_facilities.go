package database

import "time"

// MasterFacilities will migrate a master facilities table with the given specification into the database
type MasterFacilities struct {
	ID          uint      `gorm:"primary_key;autoIncrement;not null" json:"id"`
	FacCategory uint      `gorm:"not null" json:"fac_category"` // fasilitas umum , fasilitas ruangan , fasilitas kosan
	FacName     string    `gorm:"not null" json:"fac_name"`
	IsActive    bool      `gorm:"not null;default:true" json:"is_active"`
	Created     time.Time `gorm:"type:datetime" json:"created"`
	CreatedBy   string    `json:"created_by"`
	Modified    time.Time `gorm:"type:datetime" json:"modified"`
	ModifiedBy  string    `json:"modified_by"`
}

// MasterFacilitiesTable set the migrated struct table name
func (masterFacilities *MasterFacilities) MasterFacilitiesTable() string {
	return "dbMasterFacilities"
}
