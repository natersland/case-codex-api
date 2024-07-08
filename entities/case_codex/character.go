package case_codex

type Character struct {
	ID             string                `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	CaseId         string                `gorm:"type:uuid;not null"`
	Name           string                `gorm:"type:varchar(255);not null"`
	Age            int                   `gorm:"type:int;not null"`
	Image          string                `gorm:"type:varchar(255);"`
	Notes          string                `gorm:"type:text;"`
	Type           string                `gorm:"type:varchar(255);not null"` // criminal, witness, victim, detective, suspect, police, protagonist, etc.
	Investigations []PlayerInvestigation `gorm:"foreignKey:CharacterID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt      string                `gorm:"type:timestamp;default:now()"`
	UpdatedAt      string                `gorm:"type:timestamp;default:now()"`
}
