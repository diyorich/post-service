package model

import "time"

const (
	MaleGender      = "Male"
	FemaleGender    = "Female"
	NonBinaryGender = "Non-binary"
)

type Post struct {
	ID        uint64
	FirstName string
	LastName  string
	Email     string
	Gender    string
	IPAddress string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type PostJSON struct {
	ID        uint64 `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Gender    string `json:"gender"`
	IPAddress string `json:"ip_address"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func (p *Post) IsValidGender() bool {
	if p.Gender != MaleGender && p.Gender != FemaleGender {
		return false
	}

	return true
}
