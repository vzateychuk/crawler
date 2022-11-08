package dao

import (
	"time"

	"github.com/google/uuid"
)

type Edge struct {
	ID        uuid.UUID
	Src       uuid.UUID
	Dst       uuid.UUID
	UpdatedAt time.Time
}
