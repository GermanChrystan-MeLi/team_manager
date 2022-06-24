package dto

import "github.com/GermanChrystan-MeLi/team_manager/pkg/utils/constants"

// Not used for now
type BasicPlayerDTO struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
	Country   int    `json:"country"`
	Injured   bool   `json:"injured"`
	Suspended bool   `json:"suspended"`
}

// Not used for now
type PlayerStatisticsDataDTO struct {
	PlayerID    int `json:"player_id"`
	Appearances int `json:"appearances"`
	Goals       int `json:"goals"`
	YellowCards int `json:"yellow_cards"`
	RedCards    int `json:"red_cards"`
	Assists     int `json:"assists"`
	Passes      int `json:"passes"`
	Suspensions int `json:"suspensions"`
	Injuries    int `json:"injuries"`
}

type PlayerCardDTO struct {
	BasicData       PlayerBasicDataDTO     `json:"basic_data"`
	PhysicalData    PlayerPhysicalDataDTO  `json:"physical_data"`
	PlayerBaseStats PlayerBaseStatsDataDTO `json:"base_stats"`
	// Calculated contract data
	// ContractLength int     `json:"contract_length"`
	// ContractPrice  float32 `json:"contract_price"`
}

type PlayerBasicDataDTO struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Country   int    `json:"country"`
}
type PlayerPhysicalDataDTO struct {
	Height        float32                 `json:"height"`
	BasePosition  constants.BasePosition  `json:"position"`
	Age           int                     `json:"age"`
	PhysicalState constants.PhysicalState `json:"physical_state"`
	Footedness    constants.Footedness    `json:"footedness"`
}

type PlayerBaseStatsDataDTO struct {
	Charisma     int `json:"charisma"`
	Intelligence int `json:"intelligence"`
	Endurance    int `json:"endurance"`
	Accuracy     int `json:"accuracy"`
	Strength     int `json:"strength"`
	Agility      int `json:"agility"`
	BallHandling int `json:"ball_handling"`
	Blocking     int `json:"blocking"`
}
