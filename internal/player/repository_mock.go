package player

import (
	"context"

	"github.com/GermanChrystan-MeLi/team_manager/internal/domain"
	"github.com/GermanChrystan-MeLi/team_manager/internal/dto"
	"github.com/GermanChrystan-MeLi/team_manager/pkg/utils/error_vars"
)

type mockPlayerDB struct {
	PlayerBasicData     []domain.Player
	PlayerPhysicalData  []domain.PlayerPhysicalData
	PlayerBaseStatsData []domain.PlayerBaseStatsData
}

type mockRepository struct {
	db         *mockPlayerDB
	wasWritten bool
	wasRead    bool
}

//=================================================================================//
func NewMockRepository(db *mockPlayerDB) PlayerRepository {
	return &mockRepository{
		db:         db,
		wasWritten: false,
		wasRead:    false,
	}
}

//=================================================================================//
func (r *mockRepository) resetValues() {
	r.wasWritten = false
	r.wasRead = false
}

//=================================================================================//
func (r *mockRepository) GetPlayerById(ctx context.Context, id string) (dto.PlayerCardDTO, error) {
	r.resetValues()
	var result dto.PlayerCardDTO

	for _, p := range r.db.PlayerBasicData {
		if p.ID == id {
			r.wasRead = true
			result.BasicData.ID = p.ID
			result.BasicData.FirstName = p.FirstName
			result.BasicData.LastName = p.LastName
			result.BasicData.Country = int(p.Country)
		}
		if !r.wasRead {
			return result, error_vars.XNotFound("player")
		}
	}

	for _, p := range r.db.PlayerPhysicalData {
		if p.PlayerID == id {
			result.PhysicalData.Height = p.Height
			result.PhysicalData.BasePosition = p.BasePosition
			result.PhysicalData.Age = p.Age
			result.PhysicalData.PhysicalState = p.PhysicalState
			result.PhysicalData.Footedness = p.Footedness
		}
	}

	for _, p := range r.db.PlayerBaseStatsData {
		if p.PlayerID == id {
			result.PlayerBaseStats.Charisma = p.Charisma
			result.PlayerBaseStats.Intelligence = p.Intelligence
			result.PlayerBaseStats.Endurance = p.Endurance
			result.PlayerBaseStats.Accuracy = p.Accuracy
			result.PlayerBaseStats.Strength = p.Strength
			result.PlayerBaseStats.Agility = p.Agility
			result.PlayerBaseStats.BallHandling = p.BallHandling
			result.PlayerBaseStats.Blocking = p.Blocking
		}
	}
	return result, nil
}

//=================================================================================//
func (r *mockRepository) CreatePlayerBasicData(ctx context.Context, newPlayer domain.Player) error {
	r.resetValues()

	r.db.PlayerBasicData = append(r.db.PlayerBasicData, newPlayer)
	r.wasWritten = true
	return nil
}

//=================================================================================//
func (r *mockRepository) CreatePlayerPhysicalData(ctx context.Context, newPlayerPhysicalData domain.PlayerPhysicalData) error {
	r.resetValues()
	r.db.PlayerPhysicalData = append(r.db.PlayerPhysicalData, newPlayerPhysicalData)
	r.wasWritten = true
	return nil
}

//=================================================================================//
func (r *mockRepository) CreatePlayerBaseStats(ctx context.Context, newPlayerStatsData domain.PlayerBaseStatsData) error {
	r.resetValues()
	r.db.PlayerBaseStatsData = append(r.db.PlayerBaseStatsData, newPlayerStatsData)
	r.wasWritten = true
	return nil
}
