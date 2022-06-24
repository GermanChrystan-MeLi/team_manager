package domain

import "github.com/GermanChrystan-MeLi/team_manager/pkg/utils/constants"

type PlayerPhysicalData struct {
	ID            string                  `json:"id"`
	PlayerID      string                  `json:"player_id"`
	Height        float32                 `json:"height"`
	BasePosition  constants.BasePosition  `json:"position"`
	Age           int                     `json:"age"`
	PhysicalState constants.PhysicalState `json:"physical_state"`
	Footedness    constants.Footedness    `json:"footedness"`
	Talent        constants.Talent        `json:"talent"`
}
