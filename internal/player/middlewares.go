package player

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type PlayerMiddleware interface {
	CreatePlayerBasicData() gin.HandlerFunc
	CreatePlayerPhysicalData() gin.HandlerFunc
	CreatePlayerBaseStatsData() gin.HandlerFunc
}

func respondWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"error": message})
}

//=================================================================================//
type middleware struct {
	playerService PlayerService
}

//=================================================================================//
func NewPlayerMiddleware(ps PlayerService) PlayerMiddleware {
	return &middleware{
		playerService: ps,
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

		newID, err := m.playerService.CreatePlayerBasicData(ctx, country)
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
		playerID := ctx.Request.Header.Get("player_id")

		basePosition, err := m.playerService.CreatePlayerPhysicalData(ctx, playerID)
		if err != nil {
			respondWithError(ctx, 400, err.Error())
		}

		ctx.Writer.Header().Set("base_position", strconv.Itoa(basePosition))
		ctx.Next()
	}
}

//=================================================================================//
func (m *middleware) CreatePlayerBaseStatsData() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		playerID := ctx.Request.Header.Get("player_id")

		basePositionHeader := ctx.Request.Header.Get("base_position")

		err := m.playerService.CreatePlayerBaseStatsData(ctx, playerID, basePositionHeader)
		if err != nil {
			respondWithError(ctx, 400, err.Error())
		}
		ctx.Next()
	}
}

//=================================================================================//
