package entity

import (
	"github.com/google/uuid"
	"time"
	"url-shortener-api/internal/hash"
)

type URL struct {
	ID         uuid.UUID
	TargetLink string
	Hash       string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func NewURL(targetLink string) *URL {
	ID := uuid.New()
	return &URL{
		ID:         ID,
		TargetLink: targetLink,
		Hash:       hash.GenerateHash(targetLink, ID.String()),
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
}
