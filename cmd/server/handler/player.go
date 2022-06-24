package handler

import (
	"context"

	"github.com/GermanChrystan-MeLi/team_manager/internal/player"
	"github.com/GermanChrystan-MeLi/team_manager/pkg/web"
	"github.com/gin-gonic/gin"
)

//=================================================================================//
type Player struct {
	playerService player.PlayerService
}

//=================================================================================//
func NewPlayer(ps player.PlayerService) *Player {
	return &Player{
		playerService: ps,
	}
}

//=================================================================================//
func (p *Player) GetPlayerById() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()
		playerID := c.Param("id")
		if playerID == "" {
			web.Error(c, 400, "missing player id")
			return
		}
		player, err := p.playerService.GetPlayerById(ctx, playerID)
		if err != nil {
			web.Error(c, 404, err.Error())
			return
		}
		web.Success(c, 200, player)
	}
}

//=================================================================================//
