package player

import "time"

type Rank struct {
	ID         string                `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	RankName   string                `gorm:"type:varchar(255);not null"`
	Conditions []PlayerRankCondition `gorm:"foreignKey:PlayerID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt  time.Time             `gorm:"not null;autoCreateTime"`
	UpdatedAt  time.Time             `gorm:"not null;autoUpdateTime"`
}

type PlayerRankCondition struct {
	PlayerID        string `gorm:"type:uuid;not null"`
	RankConditionID string `gorm:"type:uuid;not null"`
	CurrentProgress int    `gorm:"type:integer;not null"`
	IsCompleted     bool   `gorm:"type:boolean;not null;default:false"`
	CreatedAt       string `gorm:"type:timestamp;default:now()"`
	UpdatedAt       string `gorm:"type:timestamp;default:now()"`
}

type RankCondition struct {
	ID        string `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	RankID    string `gorm:"type:uuid;not null"`
	Objective string `gorm:"type:varchar(255);not null"`
	Goal      int    `gorm:"type:integer;not null"`
}
