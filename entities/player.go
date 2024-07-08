package entities

type Player struct {
	ID          string      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	DisplayName string      `gorm:"type:varchar(255);not null;unique"`
	Email       string      `gorm:"type:varchar(255);not null"`
	Password    string      `gorm:"type:varchar(255);not null"`
	Role        string      `gorm:"type:varchar(255);not null"` // admin, user
	Rank        string      `gorm:"type:varchar(255);not null"` // newbie, beginner, intermediate, advanced, expert
	IsActivated bool        `gorm:"type:boolean;default:false"`
	Level       int         `gorm:"type:int;default:1"`
	ExpPoints   int         `gorm:"type:int;default:0"`
	Coins       int         `gorm:"type:int;default:0"`
	Avatar      string      `gorm:"type:varchar(255);"`
	Books       []Book      `gorm:"many2many:books;"`
	Cases       []CaseCodex `gorm:"many2many:cases;"`
	CreatedAt   string      `gorm:"type:timestamp;default:now()"`
	UpdatedAt   string      `gorm:"type:timestamp;default:now()"`
}
