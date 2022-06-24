package player

import (
	"context"
	"database/sql"

	"github.com/GermanChrystan-MeLi/team_manager/internal/domain"
)

//=================================================================================//
type PlayerRepository interface {
	// GetAllPlayersAvailable(ctx context.Context) ([]domain.Player, error)
	// GetOwnPlayers(ctx context.Context) ([]domain.Player, error)
	// GetPlayerById(ctx context.Context) ([]domain.Player, error)

	// UpdatePlayer(ctx context.Context) (domain.Player, error)

	CreatePlayerBasicData(ctx context.Context, newPlayer domain.Player) error
	CreatePlayerPhysicalData(ctx context.Context, newPlayerPhysicalData domain.PlayerPhysicalData) error
	CreatePlayerBaseStats(ctx context.Context, newPlayerStatsData domain.PlayerBaseStatsData) error
}

//=================================================================================//
type repository struct {
	db *sql.DB
}

//=================================================================================//
func NewRepository(db *sql.DB) PlayerRepository {
	return &repository{
		db: db,
	}
}

//=================================================================================//
func (r *repository) CreatePlayerBasicData(ctx context.Context, newPlayer domain.Player) error {
	insertPlayerQuery := "INSERT into players (id, first_name, last_name, country) VALUES(?,?,?,?)"
	stmt, err := r.db.Prepare(insertPlayerQuery)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(
		newPlayer.ID,
		newPlayer.FirstName,
		newPlayer.LastName,
		newPlayer.Country,
	)

	if err != nil {
		return err
	}
	return nil
}

//=================================================================================//
func (r *repository) CreatePlayerPhysicalData(ctx context.Context, newPlayerPhysicalData domain.PlayerPhysicalData) error {
	// TODO: First check the is not an existent field with the same player_id
	insertPlayerPDQuery := "INSERT into players_physical_data (id, player_id, height, age, physical_state, footedness, talent, base_position) VALUES (?,?,?,?,?,?,?,?)"
	stmt, err := r.db.Prepare(insertPlayerPDQuery)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(
		newPlayerPhysicalData.ID,
		newPlayerPhysicalData.PlayerID,
		newPlayerPhysicalData.Height,
		newPlayerPhysicalData.Age,
		newPlayerPhysicalData.PhysicalState,
		newPlayerPhysicalData.Footedness,
		newPlayerPhysicalData.Talent,
		newPlayerPhysicalData.BasePosition,
	)
	if err != nil {
		return err
	}
	return nil
}

//=================================================================================//
func (r *repository) CreatePlayerBaseStats(ctx context.Context, newPlayerStats domain.PlayerBaseStatsData) error {
	insertPlayerStatsQuery := "INSERT into player_base_skills_data (id, player_id, charisma, intelligence,endurance, accuracy, strength, agility, ball_handling,blocking) VALUES(?,?,?,?,?,?,?,?,?,?)"
	stmt, err := r.db.Prepare(insertPlayerStatsQuery)
	if err != nil {
		return nil
	}
	_, err = stmt.Exec(
		newPlayerStats.ID,
		newPlayerStats.PlayerID,
		newPlayerStats.Charisma,
		newPlayerStats.Intelligence,
		newPlayerStats.Endurance,
		newPlayerStats.Accuracy,
		newPlayerStats.Strength,
		newPlayerStats.Agility,
		newPlayerStats.BallHandling,
		newPlayerStats.Blocking,
	)
	if err != nil {
		return err
	}
	return nil
}

//=================================================================================//
