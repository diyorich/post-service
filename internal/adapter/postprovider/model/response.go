package model

type PostJSON struct {
	ID        uint64 `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Gender    string `json:"gender"`
	IPAddress string `json:"address"`
}

type Meta struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
	Total  int `json:"total"`
}

type Response struct {
	Posts []PostJSON `json:"data"`
	Meta  Meta       `json:"meta"`
}
