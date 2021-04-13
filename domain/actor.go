package domain

import (
	"context"
	"time"
)

type Actor struct {
	ID        int64     `gorm:"primary_key;auto_increment" json:"id"`
	Name      string    `gorm:"type:varchar(100)" json:"name" binding:"required"`
	Bio       string    `json:"bio"`
	Image     string    `gorm:"type:varchar(256);UNIQUE" json:"imageUrl" binding:"required,url"`
	Awards    string    `json:"awards"`
	Movie     []*Movie  `gorm:"many2many:movies_actors" json:"movies"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ActorRepository interface {
	GetByID(ctx context.Context, id int64) (Actor, error)
}
