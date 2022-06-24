package domain

type Team struct {
	ID          string `json:"id"`
	UserID      string `json:"user_id"`
	FormationID string `json:"formation_id"`
	Name        string `json:"name"`
}
