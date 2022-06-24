package dto

import "github.com/GermanChrystan-MeLi/team_manager/internal/domain"

type Team struct {
	Name    string          `json:"name"`
	Players []domain.Player `json:"players"`
	// GoalKeeper  BasicPlayerDTO   `json:"goal_keeper"`

	// Defenders   []BasicPlayerDTO `json:"defenders"`
	// Midfielders []BasicPlayerDTO `json:"midfielders"`
	// Forwards    []BasicPlayerDTO `json:"forwards"`

	// GoalKeeperSubstitutes  []BasicPlayerDTO `json:"goal_keeper_substitutes"`
	// DefendersSubstitutes   []BasicPlayerDTO `json:"defenders_substitutes"`
	// MidfieldersSubstitutes []BasicPlayerDTO `json:"midfielders_substitutes"`
	// ForwardsSubstitures    []BasicPlayerDTO `json:"forwards_substitutes"`
}
