package player

import (
	"context"

	"github.com/GermanChrystan-MeLi/team_manager/internal/domain"
)

//=================================================================================//
type PlayerRepository interface {
	GetAllPlayersAvailable(ctx context.Context) ([]domain.Player, error)
	GetOwnPlayers(ctx context.Context) ([]domain.Player, error)
	GetPlayerById(ctx context.Context) ([]domain.Player, error)

	CreatePlayer(ctx context.Context) (domain.Player, error)
	UpdatePlayer(ctx context.Context) (domain.Player, error)
}

//=================================================================================//
