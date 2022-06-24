package dto

import "github.com/GermanChrystan-MeLi/team_manager/pkg/utils/constants"

type BasicPlayerDTO struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
	Country   int    `json:"country"`
	Injured   bool   `json:"injured"`
	Suspended bool   `json:"suspended"`
}

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
	// Data from BasicData
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Country   int    `json:"country"`
	// Data from PhysicalData
	Height        float32                 `json:"height"`
	BasePosition  constants.BasePosition  `json:"position"`
	Age           int                     `json:"age"`
	PhysicalState constants.PhysicalState `json:"physical_state"`
	Footedness    constants.Footedness    `json:"footedness"`
	// Data from BaseStats
	Charisma     int `json:"charisma"`
	Intelligence int `json:"intelligence"`
	Endurance    int `json:"endurance"`
	Accuracy     int `json:"accuracy"`
	Strength     int `json:"strength"`
	Agility      int `json:"agility"`
	BallHandling int `json:"ball_handling"`
	Blocking     int `json:"blocking"`
	// Calculated contract data
	ContractLength int     `json:"contract_length"`
	ContractPrice  float32 `json:"contract_price"`
}
