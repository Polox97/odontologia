package domain

type Turno struct {
	ID          int    `json:"id"`
	PacienteID  int    `json:"paciente_id"`
	DentistaID  int    `json:"dentista_id"`
	FechaHora   string `json:"fecha_hora"`
	Descripcion string `json:"descripcion"`
}
