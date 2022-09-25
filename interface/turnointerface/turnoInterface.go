package turnointerface

import (
	turnoModel "github.com/Polox97/odontologia/model/turno"
)

type TurnoI interface {
	GetAll() ([]turnoModel.Turno, error)
	GetByID(id int) (turnoModel.Turno, error)
	GetPaciente(dni string) ([]turnoModel.TurnoResponse, error)
	Create(p turnoModel.Turno) (turnoModel.Turno, error)
	Update(id int, p turnoModel.Turno) (turnoModel.Turno, error)
	Delete(id int) error
}