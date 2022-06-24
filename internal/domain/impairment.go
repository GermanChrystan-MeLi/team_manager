package domain

import "github.com/GermanChrystan-MeLi/team_manager/pkg/utils/constants"

type Impairment struct {
	ID             int                         `json:"id"`
	PlayerID       int                         `json:"player_id"`
	ImpairmentType constants.ImpairmentType    `json:"impairment_type"`
	Gravity        constants.ImpairmentGravity `json:"gravity"`
	Cause          string                      `json:"cause"`
	ExpirationDate int                         `json:"expiration_date"`
}
