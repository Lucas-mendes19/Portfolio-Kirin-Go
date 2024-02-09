package entity

import (
	"github.com/google/uuid"
)

type Id = uuid.UUID

func NewId() Id {
	return uuid.New()
}
