package main

import (
	"database/sql"
	"os"

	"github.com/GermanChrystan-MeLi/team_manager/cmd/server/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// name, _ := names.CreateFullNameByNationality(constants.Brazil)
	// fmt.Println(name)
	_ = godotenv.Load()
	db, _ := sql.Open("mysql", os.Getenv("DB_URL"))

	r := gin.Default()
	router := routes.NewRouter(r, db)
	router.MapRoutes()

	if err := r.Run(); err != nil {
		panic(err)
	}
}
