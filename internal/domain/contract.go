package domain

type Contract struct {
	ID             int `json:"id"`
	PlayerID       int `json:"player_id"`
	TeamID         int `json:"team_id"`
	ExpirationDate int `json:"expiration_date"`
}
