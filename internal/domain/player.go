package domain

import "github.com/GermanChrystan-MeLi/team_manager/utils/constants"

type Player struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Country   string `json:"country"`
}

type PlayerPhysicalData struct {
	PlayerID int `json:"player_id"`

	Height        float32                 `json:"height"`
	Position      constants.Position      `json:"position"`
	Age           int                     `json:"age"`
	PhysicalState constants.PhysicalState `json:"physical_state"`

	Footedness constants.Footedness `json:"footedness"`
	Talent     constants.Talent     `json:"talent"`
}

type PlayerBaseSkillsData struct {
	PlayerID     int `json:"player_id"`
	Charisma     int `json:"charisma"`
	Endurance    int `json:"endurance"`
	Accuracy     int `json:"accuracy"`
	Strength     int `json:"strength"`
	Agility      int `json:"agility"`
	BallHandling int `json:"ball_handling"`
	Blocking     int `json:"blocking"`
	Intelligence int `json:"intelligence"`
}

type PlayerTeamPosition struct {
	PlayerID     int                `json:"player_id"`
	Position     constants.Position `json:"position"`
	IsMainPlayer bool               `json:"is_main_player"`
}

type Impairment struct {
	ID             int                         `json:"id"`
	ImpairmentType constants.ImpairmentType    `json:"impairment_type"`
	PlayerID       int                         `json:"player_id"`
	Gravity        constants.ImpairmentGravity `json:"gravity"`
	Cause          string                      `json:"cause"`
	ExpirationDate int                         `json:"expiration_date"`
}

type Contract struct {
	ID             int `json:"id"`
	PlayerID       int `json:"player_id"`
	ManagerID      int `json:"manager_id"`
	ExpirationDate int `json:"expiration_date"`
}
