package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"database/sql"

	"github.com/Polox97/odontologia/cmd/server/handler"
	"github.com/Polox97/odontologia/pkg/store"
	"github.com/Polox97/odontologia/internal/dentista"
	"github.com/Polox97/odontologia/internal/paciente"
	"github.com/Polox97/odontologia/internal/turno"
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

	dStorage := store.NewSqlStoreDentista(db)
	pStorage := store.NewSqlStorePaciente(db)
	tStorage := store.NewSqlStoreTurno(db)

	dRepo := dentista.NewRepository(dStorage)
	dService := dentista.NewService(dRepo)
	dentistatHandler := handler.NewDentistaHandler(dService)

	pRepo := paciente.NewRepository(pStorage)
	pService := paciente.NewService(pRepo)
	pacienteHandler := handler.NewPacienteHandler(pService)

	tRepo := turno.NewRepository(tStorage)
	tService := turno.NewService(tRepo)
	turnoHandler := handler.NewTurnoHandler(tService)

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

	pacientes := r.Group("/pacientes")
	{
		pacientes.GET("", pacienteHandler.GetAll())
		pacientes.GET(":id", pacienteHandler.GetByID())
		pacientes.POST("", pacienteHandler.Post())
		pacientes.DELETE(":id", pacienteHandler.Delete())
		pacientes.PATCH(":id", pacienteHandler.Patch())
		pacientes.PUT(":id", pacienteHandler.Put())
	}

	turnos := r.Group("/turnos")
	{
		turnos.GET("", turnoHandler.GetAll())
		turnos.GET(":id", turnoHandler.GetByID())
		turnos.POST("", turnoHandler.Post())
		turnos.DELETE(":id", turnoHandler.Delete())
		turnos.PATCH(":id", turnoHandler.Patch())
		turnos.PUT(":id", turnoHandler.Put())
	}

	r.Run(":8080")
}
