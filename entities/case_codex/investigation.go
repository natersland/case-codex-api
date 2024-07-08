package case_codex

type Investigation struct {
	ID      string `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Title   string `gorm:"type:varchar(255);not null"`
	Details string `gorm:"type:text;"`
}
