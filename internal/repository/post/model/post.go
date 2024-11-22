package model

import (
	"github.com/uptrace/bun"
	"time"
)

type Post struct {
	ID            uint64    `bun:"id"`
	FirstName     string    `bun:"first_name"`
	LastName      string    `bun:"last_name"`
	Email         string    `bun:"email"`
	Gender        string    `bun:"gender"`
	IPAddress     string    `bun:"ip_address"`
	CreatedAt     time.Time `bun:"created_at,nullzero,notnull,default:current_timestamp"`
	UpdatedAt     time.Time `bun:"updated_at,nullzero,notnull,default:current_timestamp"`
	bun.BaseModel `bun:"table:post"`
}
