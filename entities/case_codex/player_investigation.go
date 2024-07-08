package case_codex

type PlayerInvestigation struct {
	PlayerID        string `gorm:"type:uuid;not null"`
	CaseID          string `gorm:"type:uuid;not null"`
	InvestigationID string `gorm:"type:uuid;not null"`
	Status          string `gorm:"type:varchar(255);not null"` // open, closed, solved, envidence, etc.
	CreatedAt       string `gorm:"type:timestamp;default:now()"`
	UpdatedAt       string `gorm:"type:timestamp;default:now()"`
}
