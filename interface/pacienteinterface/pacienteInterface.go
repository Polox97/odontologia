package pacienteinterface

import (
	pacienteModel "github.com/Polox97/odontologia/model/paciente"
)


type PacienteI interface {
	GetAll() ([]pacienteModel.Paciente, error)
	GetByID(id int) (pacienteModel.Paciente, error)
	Create(p pacienteModel.Paciente) (pacienteModel.Paciente, error)
	Update(id int, p pacienteModel.Paciente) (pacienteModel.Paciente, error)
	Delete(id int) error
}