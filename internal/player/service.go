package player

import (
	"context"

	"github.com/GermanChrystan-MeLi/team_manager/internal/domain"
)

//=================================================================================//
type PlayerService interface {
	GetAllPlayersAvailable(ctx context.Context) ([]domain.Player, error)
	GetOwnPlayers(ctx context.Context) ([]domain.Player, error)
	GetPlayerById(ctx context.Context) ([]domain.Player, error)

	CreatePlayer(ctx context.Context) (domain.Player, error)
	UpdatePlayer(ctx context.Context) (domain.Player, error)
}

//=================================================================================//
type service struct {
	repository PlayerRepository
}

//=================================================================================//
func NewService(repository PlayerRepository) PlayerService {
	return &service{
		repository: repository,
	}
}

//=================================================================================//
func (s *service) GetAllPlayersAvailable(ctx context.Context) ([]domain.Player, error) {
	return s.repository.GetAllPlayersAvailable(ctx)

}

//=================================================================================//
func (s *service) GetOwnPlayers(ctx context.Context) ([]domain.Player, error) {

}

//=================================================================================//
func (s *service) GetPlayerById(ctx context.Context) ([]domain.Player, error) {}

//=================================================================================//
func (s *service) CreatePlayer(ctx context.Context, nationality) (domain.Player, error) {

}

//=================================================================================//
func (s *service) UpdatePlayer(ctx context.Context) (domain.Player, error) {}

//=================================================================================//
