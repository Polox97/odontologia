package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"database/sql"

	"github.com/Polox97/odontologia/handler"
	"github.com/Polox97/odontologia/pkg/store"
	service "github.com/Polox97/odontologia/service"
	repo "github.com/Polox97/odontologia/repository"
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
	dRepo := repo.NewDentistaRepo(dStorage)
	dService := service.NewDentistaService(dRepo)
	dentistatHandler := handler.NewDentistaHandler(dService)

	pStorage := store.NewSqlStorePaciente(db)
	pRepo := repo.NewPacienteRepository(pStorage)
	pService := service.NewPacienteService(pRepo)
	pacienteHandler := handler.NewPacienteHandler(pService)

	tStorage := store.NewSqlStoreTurno(db)
	tRepo := repo.NewRepository(tStorage)
	tService := service.NewService(tRepo)
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
		turnos.GET("/paciente/:id", turnoHandler.GetPaciente())
		turnos.POST("", turnoHandler.Post())
		turnos.DELETE(":id", turnoHandler.Delete())
		turnos.PATCH(":id", turnoHandler.Patch())
		turnos.PUT(":id", turnoHandler.Put())
	}

	r.Run(":8080")
}
