package pacienteucinterface

import (
	pacienteModel "github.com/Polox97/odontologia/model/paciente"
)

type PacienteUCI interface {
	GetAllPacientes() ([]pacienteModel.Paciente, error)
	GetPacienteByID(id int) (pacienteModel.Paciente, error)
	CreatePaciente(p pacienteModel.Paciente) (pacienteModel.Paciente, error)
	DeletePaciente(id int) error
	UpdatePaciente(id int, p pacienteModel.Paciente) (pacienteModel.Paciente, error)
}