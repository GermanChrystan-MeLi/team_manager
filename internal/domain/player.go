package domain

import "github.com/GermanChrystan-MeLi/team_manager/pkg/utils/constants"

type Player struct {
	ID        string            `json:"id"`
	FirstName string            `json:"first_name"`
	LastName  string            `json:"last_name"`
	Country   constants.Country `json:"country"`
}
