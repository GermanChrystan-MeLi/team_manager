package player

import (
	"context"
	"math/rand"
	"strconv"
	"time"

	"github.com/GermanChrystan-MeLi/team_manager/internal/domain"
	"github.com/GermanChrystan-MeLi/team_manager/internal/dto"
	"github.com/GermanChrystan-MeLi/team_manager/pkg/utils/constants"
	"github.com/GermanChrystan-MeLi/team_manager/pkg/utils/names"
	"github.com/GermanChrystan-MeLi/team_manager/pkg/utils/stats"
	"github.com/GermanChrystan-MeLi/team_manager/pkg/utils/talent"
	"github.com/google/uuid"
)

//=================================================================================//
type PlayerService interface {
	// GetAllPlayersAvailable(ctx context.Context) ([]domain.Player, error)
	// GetOwnPlayers(ctx context.Context) ([]domain.Player, error)
	// GetPlayerById(ctx context.Context) ([]domain.Player, error)

	GetPlayerById(ctx context.Context, id string) (dto.PlayerCardDTO, error)

	CreatePlayerBasicData(ctx context.Context, country int) (string, error)
	CreatePlayerPhysicalData(ctx context.Context, player string) (int, error)
	CreatePlayerBaseStatsData(ctx context.Context, playerID, basePositionStr string) error

	// UpdatePlayer(ctx context.Context) (domain.Player, error)
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
// func (s *service) GetAllPlayersAvailable(ctx context.Context) ([]domain.Player, error) {
// 	// return s.repository.GetAllPlayersAvailable(ctx)

// }

//=================================================================================//
// func (s *service) GetOwnPlayers(ctx context.Context) ([]domain.Player, error) {

// }

//=================================================================================//
func (s *service) GetPlayerById(ctx context.Context, id string) (dto.PlayerCardDTO, error) {
	return s.repository.GetPlayerById(ctx, id)
}

//=================================================================================//
// func (s *service) UpdatePlayer(ctx context.Context) (domain.Player, error) {}

//=================================================================================//
func (s *service) CreatePlayerBasicData(ctx context.Context, country int) (string, error) {
	newID := uuid.New().String()

	PlayerFullName, err := names.CreateFullNameByNationality(constants.Country(country))
	if err != nil {
		return "", err
	}
	// Creating Player struct
	NewPlayer := domain.Player{
		ID:        newID,
		FirstName: PlayerFullName.FirstName,
		LastName:  PlayerFullName.LastName,
		Country:   constants.Country(country),
	}
	err = s.repository.CreatePlayerBasicData(ctx, NewPlayer)
	if err != nil {
		return "", err
	}
	return newID, nil
}

//=================================================================================//
func (s *service) CreatePlayerPhysicalData(ctx context.Context, playerID string) (int, error) {
	rand.Seed(time.Now().UTC().UnixNano())

	height := constants.MinHeight + rand.Float32()*(constants.MaxHeight-constants.MinHeight)
	age := constants.MinAge + rand.Int()*(constants.MaxAge-constants.MinAge)
	footedness := rand.Intn(2)
	talent, err := talent.GetRandomTalent()
	if err != nil {
		return 0, err
	}
	basePosition := rand.Intn(4)

	newPlayerPhysicalData := domain.PlayerPhysicalData{
		ID:            uuid.New().String(),
		PlayerID:      playerID,
		Height:        height,
		Age:           age,
		PhysicalState: constants.PhysicalState(5),
		Footedness:    constants.Footedness(footedness),
		Talent:        talent,
		BasePosition:  constants.BasePosition(basePosition),
	}

	err = s.repository.CreatePlayerPhysicalData(ctx, newPlayerPhysicalData)
	if err != nil {
		return 0, err
	}

	return basePosition, nil
}

//=================================================================================//
func (s *service) CreatePlayerBaseStatsData(ctx context.Context, playerID, basePositionStr string) error {
	basePositionInt, err := strconv.Atoi(basePositionStr)
	if err != nil {
		return err
	}
	basePosition := constants.BasePosition(basePositionInt)

	stats, err := stats.CreateStatsFromBasePosition(basePosition)
	if err != nil {
		return err
	}

	newPlayerStats := domain.PlayerBaseStatsData{
		ID:           uuid.New().String(),
		PlayerID:     playerID,
		Charisma:     stats["charisma"],
		Intelligence: stats["intelligence"],
		Endurance:    stats["endurance"],
		Accuracy:     stats["accuracy"],
		Strength:     stats["strength"],
		Agility:      stats["agility"],
		BallHandling: stats["ball_handling"],
		Blocking:     stats["blocking"],
	}
	err = s.repository.CreatePlayerBaseStats(ctx, newPlayerStats)
	if err != nil {
		return err
	}
	return nil
}

//=================================================================================//
