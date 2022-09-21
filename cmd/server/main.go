package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"database/sql"

	"github.com/Polox97/odontologia/cmd/server/handler"
	"github.com/Polox97/odontologia/internal/dentista"
	"github.com/Polox97/odontologia/pkg/store/dentista"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error al intentar cargar archivo .env")
	}

	user := os.Getenv("USER_DB")
	password := os.Getenv("PASSWORD_DB")
	dataBase := os.Getenv("DB")

	rute := user + ":" + password + "@/" + dataBase

	db, err := sql.Open("mysql", rute)
	if err != nil {
		panic(err.Error())
	}
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	storage := store.NewSqlStore(db)

	repo := dentista.NewRepository(storage)
	service := dentista.NewService(repo)
	dentistatHandler := handler.NewDentistaHandler(service)

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) { c.String(200, "pong") })
	dentistas := r.Group("/dentistas")
	{
		dentistas.GET("", dentistatHandler.GetAll())
		dentistas.GET(":id", dentistatHandler.GetByID())
		dentistas.POST("", dentistatHandler.Post())
		dentistas.DELETE(":id", dentistatHandler.Delete())
		dentistas.PATCH(":id", dentistatHandler.Patch())
		dentistas.PUT(":id", dentistatHandler.Put())
	}

	r.Run(":8080")
}
