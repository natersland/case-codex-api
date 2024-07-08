package entities

type Book struct {
	ID         string      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Players    []Player    `gorm:"many2many:books;"`
	Cases      []CaseCodex `gorm:"foreignKey:BookID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Title      string      `gorm:"type:varchar(255);not null"`
	Author     string      `gorm:"type:varchar(255);not null"`
	ISBN       string      `gorm:"type:varchar(255);not null"`
	Year       int         `gorm:"type:int;not null"`
	Pages      int         `gorm:"type:int;not null"`
	Price      float64     `gorm:"type:float;not null"`
	Rating     float64     `gorm:"type:float;not null"`
	CoverImage string      `gorm:"type:varchar(255);"`
	CreatedAt  string      `gorm:"type:timestamp;default:now()"`
	UpdatedAt  string      `gorm:"type:timestamp;default:now()"`
}
