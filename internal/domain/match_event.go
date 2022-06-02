package domain

type MatchEvent struct {
	ID       string `json:"id"`
	MatchID  string `json:"match_id"`
	TeamID   string `json:"team_id"`
	PlayerID string `json:"player_id"`
}
