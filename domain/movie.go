package domain

import (
	"time"
)

type Movie struct {
	ID           int64         `gorm:"primary_key;auto_increment" json:"id"`
	Title        string        `gorm:"type:varchar(100)" json:"title" binding:"min=2,max=10,required"`
	Description  string        `gorm:"type:varchar(100)" json:"description" binding:"max=20"`
	Year         int           `json:"year" binding:"min=100"`
	Language     string        `gorm:"type:varchar(20)" json:"language" binding:"min=2,max=10"`
	ImdbScore    float32       `json:"imdbScore" binding:"min=2,max=10"`
	TrailerLinks []TrailerLink `gorm:"foreignKey:MovieID" json:"trailerLinks,omitempty"`
	Directors    []*Director   `gorm:"many2many:movies_directors" json:"directors,omitempty"`
	Actors       []*Actor      `gorm:"many2many:movies_actors" json:"actors,omitempty"`
	CreatedAt    time.Time     `json:"created_at"`
	UpdatedAt    time.Time     `json:"updated_at"`
}

type TrailerLink struct {
	ID      int64  `json:"id"`
	Link    string `json:"link"`
	Site    string `json:"site"`
	MovieID int64  `json:"movieID"`
}
