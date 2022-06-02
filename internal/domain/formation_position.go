package domain

type FormationPosition struct {
	ID          string `json:"id"`
	FormationID string `json:"formation_id"`
	PositionID  string `json:"position_id"`
}
