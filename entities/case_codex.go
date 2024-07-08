package entities

type CaseCodex struct {
	ID         string   `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	BookID     string   `gorm:"type:uuid;not null"`
	Players    []Player `gorm:"many2many:cases;"`
	Status     string   `gorm:"type:varchar(255);not null"` // drop, on hold in progress, completed
	Difficulty string   `gorm:"type:varchar(255);not null"` // easy, medium, hard, expert
	StartedAt  string   `gorm:"type:timestamp;default:now()"`
	EndedAt    string   `gorm:"type:timestamp;default:now()"`
	CreatedAt  string   `gorm:"type:timestamp;default:now()"`
	UpdatedAt  string   `gorm:"type:timestamp;default:now()"`
}
