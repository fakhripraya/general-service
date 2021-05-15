package database

import "time"

// DBKost will migrate a kost table with the given specification into the database
type DBKost struct {
	ID            uint      `gorm:"primary_key;autoIncrement;not null" json:"id"`
	OwnerID       uint      `gorm:"not null" json:"owner_id"`
	TypeID        uint      `gorm:"not null" json:"type_id"`
	Status        uint      `gorm:"not null" json:"status"`
	KostCode      string    `gorm:"not null" json:"kost_code"`
	KostName      string    `gorm:"not null" json:"kost_name"`
	KostDesc      string    `gorm:"not null" json:"kost_desc"`
	Country       string    `gorm:"not null" json:"country"`
	City          string    `gorm:"not null" json:"city"`
	Address       string    `gorm:"not null" json:"address"`
	Latitude      string    `gorm:"not null" json:"latitude"`
	Longitude     string    `gorm:"not null" json:"longitude"`
	UpRate        uint64    `json:"up_rate"`
	UpRateExpired time.Time `json:"up_rate_expired"`
	ThumbnailURL  string    `json:"thumbnail_url"`
	IsVerified    bool      `gorm:"not null;default:false" json:"is_verified"`
	IsActive      bool      `gorm:"not null;default:true" json:"is_active"`
	Created       time.Time `gorm:"type:datetime" json:"created"`
	CreatedBy     string    `json:"created_by"`
	Modified      time.Time `gorm:"type:datetime" json:"modified"`
	ModifiedBy    string    `json:"modified_by"`
}

// DBKostPeriod will migrate a kost period table with the given specification into the database
type DBKostPeriod struct {
	ID         uint      `gorm:"primary_key;autoIncrement;not null" json:"id"`
	KostID     uint      `gorm:"not null" json:"kost_id"`
	PeriodID   uint      `gorm:"not null" json:"period_id"`
	IsActive   bool      `gorm:"not null;default:true" json:"is_active"`
	Created    time.Time `gorm:"type:datetime" json:"created"`
	CreatedBy  string    `json:"created_by"`
	Modified   time.Time `gorm:"type:datetime" json:"modified"`
	ModifiedBy string    `json:"modified_by"`
}

// DBKostPict will migrate a kost pict table with the given specification into the database
type DBKostPict struct {
	ID         uint      `gorm:"primary_key;autoIncrement;not null" json:"id"`
	KostID     uint      `gorm:"not null" json:"kost_id"`
	PictDesc   string    `gorm:"not null" json:"pict_desc"`
	URL        string    `gorm:"not null" json:"url"`
	IsActive   bool      `gorm:"not null;default:true" json:"is_active"`
	Created    time.Time `gorm:"type:datetime" json:"created"`
	CreatedBy  string    `json:"created_by"`
	Modified   time.Time `gorm:"type:datetime" json:"modified"`
	ModifiedBy string    `json:"modified_by"`
}

// DBKostFacilities will migrate a kost facilities table with the given specification into the database
type DBKostFacilities struct {
	ID         uint      `gorm:"primary_key;autoIncrement;not null" json:"id"`
	FacID      uint      `gorm:"not null" json:"fac_id"`
	KostID     uint      `gorm:"not null" json:"kost_id"`
	IsActive   bool      `gorm:"not null;default:true" json:"is_active"`
	Created    time.Time `gorm:"type:datetime" json:"created"`
	CreatedBy  string    `json:"created_by"`
	Modified   time.Time `gorm:"type:datetime" json:"modified"`
	ModifiedBy string    `json:"modified_by"`
}

// DBKostReview will migrate a kost review table with the given specification into the database
type DBKostReview struct {
	ID          uint      `gorm:"primary_key;autoIncrement;not null" json:"id"`
	KostID      uint      `gorm:"not null" json:"owner_id"`
	UserID      uint      `gorm:"not null" json:"user_id"`
	Cleanliness float64   `json:"cleanliness"`
	Convenience float64   `json:"convenience"`
	Security    float64   `json:"security"`
	Facilities  float64   `json:"facilities"`
	Comments    string    `json:"comments"`
	IsActive    bool      `gorm:"not null;default:true" json:"is_active"`
	Created     time.Time `gorm:"type:datetime" json:"created"`
	CreatedBy   string    `json:"created_by"`
	Modified    time.Time `gorm:"type:datetime" json:"modified"`
	ModifiedBy  string    `json:"modified_by"`
}

// DBKostBenchmark will migrate a kost benchmark table with the given specification into the database
type DBKostBenchmark struct {
	ID            uint      `gorm:"primary_key;autoIncrement;not null" json:"id"`
	KostID        uint      `gorm:"not null" json:"kost_id"`
	BenchmarkDesc string    `gorm:"not null" json:"benchmark_desc"`
	IsActive      bool      `gorm:"not null;default:true" json:"is_active"`
	Created       time.Time `gorm:"type:datetime" json:"created"`
	CreatedBy     string    `json:"created_by"`
	Modified      time.Time `gorm:"type:datetime" json:"modified"`
	ModifiedBy    string    `json:"modified_by"`
}

// DBKostAccess will migrate a kost accessibility table with the given specification into the database
type DBKostAccess struct {
	ID                uint      `gorm:"primary_key;autoIncrement;not null" json:"id"`
	KostID            uint      `gorm:"not null" json:"kost_id"`
	AccessibilityDesc string    `gorm:"not null" json:"accessibility_desc"`
	IsActive          bool      `gorm:"not null;default:true" json:"is_active"`
	Created           time.Time `gorm:"type:datetime" json:"created"`
	CreatedBy         string    `json:"created_by"`
	Modified          time.Time `gorm:"type:datetime" json:"modified"`
	ModifiedBy        string    `json:"modified_by"`
}

// DBKostAround will migrate a kost around table with the given specification into the database
type DBKostAround struct {
	ID         uint      `gorm:"primary_key;autoIncrement;not null" json:"id"`
	KostID     uint      `gorm:"not null" json:"kost_id"`
	IconID     uint      `gorm:"not null" json:"icon_id"`
	AroundDesc string    `gorm:"not null" json:"around_desc"`
	IsActive   bool      `gorm:"not null;default:true" json:"is_active"`
	Created    time.Time `gorm:"type:datetime" json:"created"`
	CreatedBy  string    `json:"created_by"`
	Modified   time.Time `gorm:"type:datetime" json:"modified"`
	ModifiedBy string    `json:"modified_by"`
}

// DBKostRoom will migrate a kost room table with the given specification into the database
type DBKostRoom struct {
	ID            uint      `gorm:"primary_key;autoIncrement;not null" json:"id"`
	KostID        uint      `gorm:"not null" json:"kost_id"`
	RoomDesc      string    `gorm:"not null" json:"room_desc"`
	RoomPrice     float64   `gorm:"not null" json:"room_price"`
	RoomPriceUOM  uint      `gorm:"not null" json:"room_price_uom"`
	RoomLength    float64   `gorm:"not null" json:"room_length"`
	RoomWidth     float64   `gorm:"not null" json:"room_width"`
	RoomArea      float64   `gorm:"not null" json:"room_area"`
	RoomAreaUOM   uint      `gorm:"not null" json:"room_area_uom"`
	MaxPerson     uint      `gorm:"not null" json:"max_person"`
	AllowedGender string    `gorm:"not null" json:"allowed_gender"`
	Comments      string    `gorm:"not null" json:"comments"`
	IsActive      bool      `gorm:"not null;default:true" json:"is_active"`
	Created       time.Time `gorm:"type:datetime" json:"created"`
	CreatedBy     string    `json:"created_by"`
	Modified      time.Time `gorm:"type:datetime" json:"modified"`
	ModifiedBy    string    `json:"modified_by"`
}

// DBKostRoomDetail will migrate a kost room table with the given specification into the database
type DBKostRoomDetail struct {
	ID         uint      `gorm:"primary_key;autoIncrement;not null" json:"id"`
	KostID     uint      `gorm:"not null" json:"kost_id"`
	RoomID     uint      `gorm:"not null" json:"room_id"`
	RoomNumber string    `gorm:"not null" json:"room_number"`
	FloorLevel uint      `gorm:"not null" json:"floor_level"`
	IsActive   bool      `gorm:"not null;default:true" json:"is_active"`
	Created    time.Time `gorm:"type:datetime" json:"created"`
	CreatedBy  string    `json:"created_by"`
	Modified   time.Time `gorm:"type:datetime" json:"modified"`
	ModifiedBy string    `json:"modified_by"`
}

// DBKostRoomPict will migrate a kost room pict table with the given specification into the database
type DBKostRoomPict struct {
	ID         uint      `gorm:"primary_key;autoIncrement;not null" json:"id"`
	RoomID     uint      `gorm:"not null" json:"room_id"`
	PictDesc   string    `gorm:"not null" json:"pict_desc"`
	URL        string    `gorm:"not null" json:"url"`
	IsActive   bool      `gorm:"not null;default:true" json:"is_active"`
	Created    time.Time `gorm:"type:datetime" json:"created"`
	CreatedBy  string    `json:"created_by"`
	Modified   time.Time `gorm:"type:datetime" json:"modified"`
	ModifiedBy string    `json:"modified_by"`
}

// DBKostRoomFacilities will migrate a room facilities table with the given specification into the database
type DBKostRoomFacilities struct {
	ID         uint      `gorm:"primary_key;autoIncrement;not null" json:"id"`
	FacID      uint      `gorm:"not null" json:"fac_id"`
	RoomID     uint      `gorm:"not null" json:"room_id"`
	IsActive   bool      `gorm:"not null;default:true" json:"is_active"`
	Created    time.Time `gorm:"type:datetime" json:"created"`
	CreatedBy  string    `json:"created_by"`
	Modified   time.Time `gorm:"type:datetime" json:"modified"`
	ModifiedBy string    `json:"modified_by"`
}

// KostTable set the migrated struct table name
func (dbKost *DBKost) KostTable() string {
	return "dbKost"
}

// KostPeriodTable set the migrated struct table name
func (dbKostPeriod *DBKostPeriod) KostPeriodTable() string {
	return "dbKostPeriod"
}

// KostPictTable set the migrated struct table name
func (dbKostPict *DBKostPict) KostPictTable() string {
	return "dbKostPict"
}

// KostFacilitiesTable set the migrated struct table name
func (dbKostFacilities *DBKostFacilities) KostFacilitiesTable() string {
	return "dbKostFacilities"
}

// KostReviewTable set the migrated struct table name
func (dbKostReview *DBKostReview) KostReviewTable() string {
	return "dbKostReview"
}

// KostBenchmarkTable set the migrated struct table name
func (dbKostBenchmark *DBKostBenchmark) KostBenchmarkTable() string {
	return "dbKostBenchmark"
}

// KostAccessTable set the migrated struct table name
func (dbKostAccess *DBKostAccess) KostAccessTable() string {
	return "dbKostAccess"
}

// KostAroundTable set the migrated struct table name
func (dbKostAround *DBKostAround) KostAroundTable() string {
	return "dbKostAround"
}

// KostRoomTable set the migrated struct table name
func (dbKostRoom *DBKostRoom) KostRoomTable() string {
	return "dbKostRoom"
}

// KostRoomDetailTable set the migrated struct table name
func (dbKostRoomDetail *DBKostRoomDetail) KostRoomDetailTable() string {
	return "dbKostRoomDetail"
}

// KostRoomPictTable set the migrated struct table name
func (dbKostRoomPict *DBKostRoomPict) KostRoomPictTable() string {
	return "dbKostRoomPict"
}

// KostRoomFacilitiesTable set the migrated struct table name
func (dbKostRoomFacilities *DBKostRoomFacilities) KostRoomFacilitiesTable() string {
	return "dbKostRoomFacilities"
}
