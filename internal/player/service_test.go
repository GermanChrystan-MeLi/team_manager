package player

import (
	"context"
	"testing"

	"github.com/GermanChrystan-MeLi/team_manager/internal/domain"
	"github.com/GermanChrystan-MeLi/team_manager/internal/dto"
	"github.com/GermanChrystan-MeLi/team_manager/pkg/utils/constants"
	"github.com/GermanChrystan-MeLi/team_manager/pkg/utils/error_vars"

	"github.com/stretchr/testify/assert"
)

var mockDB = mockPlayerDB{
	PlayerBasicData: []domain.Player{
		{
			ID:        "test",
			FirstName: "test_first_name",
			LastName:  "test_last_name",
			Country:   constants.Argentina,
		},
	},
	PlayerPhysicalData: []domain.PlayerPhysicalData{
		{
			ID:            "some_id",
			PlayerID:      "test",
			Age:           20,
			Height:        1.70,
			PhysicalState: constants.GoodShape,
			Footedness:    constants.Left,
			BasePosition:  constants.Midfielders,
		},
	},
	PlayerBaseStatsData: []domain.PlayerBaseStatsData{
		{
			ID:           "some_id",
			PlayerID:     "test",
			Charisma:     3,
			Intelligence: 3,
			Endurance:    3,
			Accuracy:     3,
			Strength:     3,
			Agility:      3,
			BallHandling: 3,
			Blocking:     3,
		},
	},
}

var r = NewMockRepository(&mockDB)
var s = NewService(r)

//=================================================================================//
func TestGetPlayerByID_Ok(t *testing.T) {
	ctx := context.Background()
	result, err := s.GetPlayerById(ctx, "test")

	want := dto.PlayerCardDTO{
		BasicData: dto.PlayerBasicDataDTO{
			ID:        "test",
			FirstName: "test_first_name",
			LastName:  "test_last_name",
			Country:   0,
		},
		PhysicalData: dto.PlayerPhysicalDataDTO{
			Age:           20,
			Height:        1.70,
			PhysicalState: constants.GoodShape,
			Footedness:    constants.Left,
			BasePosition:  constants.Midfielders,
		},
		PlayerBaseStats: dto.PlayerBaseStatsDataDTO{
			Charisma:     3,
			Intelligence: 3,
			Endurance:    3,
			Accuracy:     3,
			Strength:     3,
			Agility:      3,
			BallHandling: 3,
			Blocking:     3,
		},
	}

	assert.Zero(t, err)
	assert.Equal(t, want, result)
}

//=================================================================================//
func TestGetPlayerByID_Error(t *testing.T) {
	ctx := context.Background()
	_, err := s.GetPlayerById(ctx, "trest")
	assert.Equal(t, error_vars.XNotFound("player"), err)
}

//=================================================================================//
func TestCreatePlayerBasicData_Ok(t *testing.T) {
	ctx := context.Background()
	id, err := s.CreatePlayerBasicData(ctx, int(constants.Argentina))
	assert.Zero(t, err)
	assert.NotEqual(t, "", id)
}

//=================================================================================//
func TestCreatePlayerBasicData_Error(t *testing.T) {
	ctx := context.Background()
	id, err := s.CreatePlayerBasicData(ctx, 45)
	assert.Error(t, err)
	assert.Equal(t, "", id)
}

//=================================================================================//
func TestCreatePlayerPhysicalData_Ok(t *testing.T) {
	ctx := context.Background()
	playerID := "some_new_id"

	basePosition, err := s.CreatePlayerPhysicalData(ctx, playerID)
	assert.Zero(t, err)
	assert.Less(t, basePosition, 4)
}

//=================================================================================//
func TestCreatePlayerBaseStatsData_Ok(t *testing.T) {
	ctx := context.Background()
	playerId := "some_new_id"
	basePosition := "1"

	err := s.CreatePlayerBaseStatsData(ctx, playerId, basePosition)

	assert.Zero(t, err)
}

//=================================================================================//
func TestCreatePlayerBaseStatsData_Error(t *testing.T) {
	ctx := context.Background()
	playerId := "some_new_id"
	basePosition := "qfq2f"

	err := s.CreatePlayerBaseStatsData(ctx, playerId, basePosition)

	assert.NotNil(t, err)
}
