package domain

import "github.com/GermanChrystan-MeLi/team_manager/utils/constants"

type PlayerPhysicalData struct {
	PlayerID int `json:"player_id"`

	Height        float32                 `json:"height"`
	Position      constants.Position      `json:"position"`
	Age           int                     `json:"age"`
	PhysicalState constants.PhysicalState `json:"physical_state"`

	Footedness constants.Footedness `json:"footedness"`
	Talent     constants.Talent     `json:"talent"`
}
