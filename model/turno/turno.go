package turno

import (
	dentista "github.com/Polox97/odontologia/model/dentista"
	paciente "github.com/Polox97/odontologia/model/paciente"
)

type Turno struct {
	ID          int    `json:"id"`
	PacienteID  int    `json:"paciente_id"`
	DentistaID  int    `json:"dentista_id"`
	FechaHora   string `json:"fecha_hora"`
	Descripcion string `json:"descripcion"`
}

type TurnoResponse struct {
	ID          int                       `json:"id"`
	FechaHora   string                    `json:"fecha_hora"`
	Descripcion string                    `json:"descripcion"`
	PacienteID  paciente.PacienteResponse `json:"paciente"`
	DentistaID  dentista.DentistaResponse `json:"dentista"`
}
