package domain

type Turno struct {
	ID          int    `json:"id"`
	PacienteID  string `json:"paciente_id"`
	DentistaID  string `json:"dentista_id"`
	FechaHora   string `json:"fecha_hora"`
	Descripcion string `json:"descripcion"`
}
