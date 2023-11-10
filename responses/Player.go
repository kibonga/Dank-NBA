package responses

type Player struct {
	ID           int    `json:"id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	HeightFeet   int    `json:"height_feet"`
	HeightInches int    `json:"height_inches"`
	Position     string `json:"position"`
	Team         Team   `json:"team"`
}
