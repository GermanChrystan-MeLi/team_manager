package player

import (
	"context"
	"database/sql"
	"errors"

	"github.com/GermanChrystan-MeLi/team_manager/internal/domain"
	"github.com/GermanChrystan-MeLi/team_manager/internal/dto"
)

//=================================================================================//
type PlayerRepository interface {
	// GetAllPlayersAvailable(ctx context.Context) ([]domain.Player, error)
	// GetOwnPlayers(ctx context.Context) ([]domain.Player, error)
	GetPlayerById(ctx context.Context, id string) (dto.PlayerCardDTO, error)

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
func (r *repository) GetPlayerById(ctx context.Context, id string) (dto.PlayerCardDTO, error) {
	var basicData dto.PlayerBasicDataDTO
	basicDataQuery := "SELECT (id, first_name, last_name, country) FROM players WHERE id=?"
	row := r.db.QueryRow(basicDataQuery, id)
	err := row.Scan(
		&basicData.ID,
		&basicData.FirstName,
		&basicData.LastName,
		&basicData.Country,
	)
	if err != nil {
		return dto.PlayerCardDTO{}, errors.New("could not retrieve basic data")
	}

	var physicalData dto.PlayerPhysicalDataDTO
	physicalDataQuery := "SELECT(height, position, age, physical_state, footedness) FROM players_physical_data WHERE player_id = ?"
	row = r.db.QueryRow(physicalDataQuery, id)
	err = row.Scan(
		&physicalData.Height,
		&physicalData.BasePosition,
		&physicalData.Age,
		&physicalData.PhysicalState,
		&physicalData.Footedness,
	)
	if err != nil {
		return dto.PlayerCardDTO{}, errors.New("could not retrieve physical data")
	}

	var baseStatsData dto.PlayerBaseStatsDataDTO
	baseStatsDataQuery := "SELECT (charisma, intelligence, endurance, accuracy, strength, agility, ball_handling, blocking) FROM players_base_stats_data WHERE player_id = ?"
	row = r.db.QueryRow(baseStatsDataQuery, id)
	err = row.Scan(
		&baseStatsData.Charisma,
		&baseStatsData.Intelligence,
		&baseStatsData.Endurance,
		&baseStatsData.Accuracy,
		&baseStatsData.Strength,
		&baseStatsData.Agility,
		&baseStatsData.BallHandling,
		&baseStatsData.Blocking,
	)
	if err != nil {
		return dto.PlayerCardDTO{}, errors.New("could not retrieve base stats data")
	}

	return dto.PlayerCardDTO{
		BasicData:       basicData,
		PhysicalData:    physicalData,
		PlayerBaseStats: baseStatsData,
	}, nil
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
	insertPlayerStatsQuery := "INSERT into players_base_stats_data (id, player_id, charisma, intelligence,endurance, accuracy, strength, agility, ball_handling,blocking) VALUES(?,?,?,?,?,?,?,?,?,?)"
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
