package model

import "time"

type Post struct {
	ID        uint64    `bun:"id"`
	FirstName string    `bun:"first_name"`
	LastName  string    `bun:"last_name"`
	Email     string    `bun:"email"`
	Gender    string    `bun:"gender"`
	IPAddress string    `bun:"ip_address"`
	CreatedAt time.Time `bun:"created_at"`
	UpdatedAt time.Time `bun:"updated_at"`
}
