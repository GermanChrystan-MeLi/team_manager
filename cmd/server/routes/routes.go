package routes

import (
	"database/sql"

	"github.com/GermanChrystan-MeLi/team_manager/cmd/server/handler"
	"github.com/GermanChrystan-MeLi/team_manager/internal/player"
	"github.com/GermanChrystan-MeLi/team_manager/internal/user"
	"github.com/GermanChrystan-MeLi/team_manager/pkg/utils/middlewares"
	"github.com/gin-gonic/gin"
)

//=================================================================================//
type Router interface {
	MapRoutes()
}

//=================================================================================//
type router struct {
	r  *gin.Engine
	rg *gin.RouterGroup
	db *sql.DB
}

//=================================================================================//
func NewRouter(r *gin.Engine, db *sql.DB) Router {
	return &router{
		r:  r,
		db: db,
	}
}

//=================================================================================//
func (r *router) MapRoutes() {
	r.setGroup()
}

func (r *router) setGroup() {
	r.rg = r.r.Group("/api/v1")

	r.userRoutes()
	r.playerRoutes()
}

//=================================================================================//
func (r *router) userRoutes() {
	userRepo := user.NewRepository(r.db)
	userService := user.NewService(userRepo)
	userHandler := handler.NewUser(userService)

	userMiddleware := middlewares.NewMiddlewareRepository(r.db)
	r.rg.POST("/user/login", userHandler.Login())
	r.rg.POST("/user/register", userHandler.Register())
	r.rg.PATCH("/user", userMiddleware.IsUserSession(), userHandler.EditUser())
	r.rg.DELETE("/user/:id", userMiddleware.IsAdminUserSession(), userHandler.DeleteUser())
}

//=================================================================================//
func (r *router) playerRoutes() {
	playerRepo := player.NewRepository(r.db)
	playerService := player.NewService(playerRepo)
	playerHandler := handler.NewPlayer(playerService)

	playerMiddleware := player.NewPlayerMiddleware(playerService)

	r.rg.POST("/player",
		playerMiddleware.CreatePlayerBasicData(),
		playerMiddleware.CreatePlayerPhysicalData(),
		playerMiddleware.CreatePlayerBaseStatsData(),
		playerHandler.GetPlayerById(),
	)

	r.rg.GET("/player/:id", playerHandler.GetPlayerById())
}

//=================================================================================//
