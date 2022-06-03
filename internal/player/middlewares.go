package player

import (
	"database/sql"
	"math/rand"
	"strconv"
	"time"

	"github.com/GermanChrystan-MeLi/team_manager/internal/domain"
	"github.com/GermanChrystan-MeLi/team_manager/utils/constants"
	"github.com/GermanChrystan-MeLi/team_manager/utils/names"
	"github.com/GermanChrystan-MeLi/team_manager/utils/talent"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PlayerMiddleware interface {
	CreatePlayerBasicData() gin.HandlerFunc
	CreatePlayerPhysicalData() gin.HandlerFunc
}

func respondWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"error": message})
}

//=================================================================================//
type middleware struct {
	db *sql.DB
}

//=================================================================================//
func NewPlayerMiddleware(db *sql.DB) PlayerMiddleware {
	return &middleware{
		db: db,
	}
}

//=================================================================================//
func (m *middleware) CreatePlayerBasicData() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		requestCountry := ctx.Request.Header.Get("country")
		country, err := strconv.Atoi(requestCountry)
		if err != nil {
			respondWithError(ctx, 404, err.Error())
		}

		newID := uuid.New().String()

		PlayerFullName, err := names.CreateFullNameByNationality(constants.Country(country))
		if err != nil {
			respondWithError(ctx, 404, err.Error())
		}
		// Creating Player struct
		NewPlayer := domain.Player{
			ID:        newID,
			FirstName: PlayerFullName.FirstName,
			LastName:  PlayerFullName.LastName,
			Country:   constants.Country(country),
		}

		insertPlayerQuery := "INSERT into players (id, first_name, last_name, country) VALUES(?,?,?,?)"
		stmt, err := m.db.Prepare(insertPlayerQuery)
		if err != nil {
			respondWithError(ctx, 404, err.Error())
		}
		_, err = stmt.Exec(NewPlayer.ID, NewPlayer.FirstName, NewPlayer.LastName, NewPlayer.Country)
		if err != nil {
			respondWithError(ctx, 400, err.Error())
		}

		ctx.Writer.Header().Set("player_id", newID)
		ctx.Next()
	}
}

//=================================================================================//
func (m *middleware) CreatePlayerPhysicalData() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		rand.Seed(time.Now().UTC().UnixNano())
		playerID := ctx.Request.Header.Get("player_id")

		height := constants.MinHeight + rand.Float32()*(constants.MaxHeight-constants.MinHeight)
		age := constants.MinAge + rand.Int()*(constants.MaxAge-constants.MinAge)
		footedness := rand.Intn(2)
		talent, err := talent.GetRandomTalent()
		if err != nil {
			respondWithError(ctx, 400, err.Error())
		}

		NewPlayerPhysicalData := domain.PlayerPhysicalData{
			ID:            uuid.New().String(),
			PlayerID:      playerID,
			Height:        height,
			Age:           age,
			PhysicalState: constants.PhysicalState(5),
			Footedness:    constants.Footedness(footedness),
			Talent:        talent,
		}

		insertPlayerPDQuery := "INSERT into players_physical_data (id, player_id, height, age, physical_state, footedness, talent) VALUES(?,?,?,?,?,?,?)"
		stmt, err := m.db.Prepare(insertPlayerPDQuery)
		if err != nil {
			respondWithError(ctx, 404, err.Error())
		}
		_, err = stmt.Exec(
			NewPlayerPhysicalData.ID,
			NewPlayerPhysicalData.PlayerID,
			NewPlayerPhysicalData.Height,
			NewPlayerPhysicalData.Age,
			NewPlayerPhysicalData.PhysicalState,
			NewPlayerPhysicalData.Footedness,
			NewPlayerPhysicalData.Talent,
		)
		if err != nil {
			respondWithError(ctx, 400, err.Error())
		}

		ctx.Next()
	}
}

//=================================================================================//
