package turnoucinterface

import (
	turnoModel "github.com/Polox97/odontologia/model/turno"
)

type TurnoUCI interface {
	// GetAll busca todos los turnos
	GetAllTurnos() ([]turnoModel.Turno, error)
	// GetByID busca un turno por su id
	GetTurnoByID(id int) (turnoModel.Turno, error)
	// GetAll busca todos los turnos
	GetTurnoPaciente(dni string) ([]turnoModel.TurnoResponse, error)
	// Create agrega un nuevo turno
	CreateTurno(p turnoModel.Turno) (turnoModel.Turno, error)
	// Delete elimina un turno
	DeleteTurno(id int) error
	// Update actualiza un turno
	UpdateTurno(id int, p turnoModel.Turno) (turnoModel.Turno, error)
}