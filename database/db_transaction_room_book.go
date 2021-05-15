package database

import "time"

// DBTransactionRoomBook is an entity that directly communicate with the TransactionRoomBook table in the database
type DBTransactionRoomBook struct {
	ID           uint      `gorm:"primary_key;autoIncrement;not null" json:"id"`
	BookerID     uint      `gorm:"not null" json:"booker_id"`
	KostID       uint      `gorm:"not null" json:"kost_id"`
	RoomID       uint      `gorm:"not null" json:"room_id"`
	RoomDetailID uint      `gorm:"not null" json:"room_detail_id"`
	PeriodID     uint      `gorm:"not null" json:"period_id"`
	Status       uint      `gorm:"not null" json:"status"`
	BookCode     string    `gorm:"not null" json:"book_code"`
	BookDate     time.Time `gorm:"not null" json:"book_date"`
	IsActive     bool      `gorm:"not null;default:true" json:"is_active"`
	Created      time.Time `gorm:"type:datetime" json:"created"`
	CreatedBy    string    `json:"created_by"`
	Modified     time.Time `gorm:"type:datetime" json:"modified"`
	ModifiedBy   string    `json:"modified_by"`
}

// DBTransactionRoomBookMember is an entity that directly communicate with the TransactionRoomBookMember table in the database
type DBTransactionRoomBookMember struct {
	ID         uint      `gorm:"primary_key;autoIncrement;not null" json:"id"`
	RoomBookID uint      `gorm:"not null" json:"room_book_id"`
	MemberName string    `gorm:"not null" json:"member_name"`
	Phone      string    `json:"phone"`
	Gender     bool      `gorm:"not null" json:"gender"`
	IsActive   bool      `gorm:"not null;default:true" json:"is_active"`
	Created    time.Time `gorm:"type:datetime" json:"created"`
	CreatedBy  string    `json:"created_by"`
	Modified   time.Time `gorm:"type:datetime" json:"modified"`
	ModifiedBy string    `json:"modified_by"`
}

// DBTransactionVerification is an entity that directly communicate with the DBTransactionVerification table in the database
type DBTransactionVerification struct {
	ID          uint      `gorm:"primary_key;autoIncrement;not null" json:"id"`
	ReferenceID uint      `gorm:"not null" json:"reference_id"`
	PictDesc    string    `gorm:"not null" json:"pict_desc"`
	URL         string    `gorm:"not null" json:"url"`
	IsActive    bool      `gorm:"not null;default:true" json:"is_active"`
	Created     time.Time `gorm:"type:datetime" json:"created"`
	CreatedBy   string    `json:"created_by"`
	Modified    time.Time `gorm:"type:datetime" json:"modified"`
	ModifiedBy  string    `json:"modified_by"`
}

// DBTransactionRoomBookTable set the migrated struct table name
func (dbTransactionRoomBook *DBTransactionRoomBook) DBTransactionRoomBookTable() string {
	return "dbTransactionRoomBook"
}

// DBTransactionRoomBookMemberTable set the migrated struct table name
func (dbTransactionRoomBookMember *DBTransactionRoomBookMember) DBTransactionRoomBookMemberTable() string {
	return "dbTransactionRoomBookMember"
}

// DBTransactionVerificationTable set the migrated struct table name
func (dbTransactionVerification *DBTransactionVerification) DBTransactionVerificationTable() string {
	return "dbTransactionVerification"
}
