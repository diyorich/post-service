package model

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
}

func (p *Post) IsValidGender() bool {
	if p.Gender != MaleGender && p.Gender != FemaleGender {
		return false
	}

	return true
}
