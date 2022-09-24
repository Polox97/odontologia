package domain

type Paciente struct {
	ID        int    `json:"id"`
	DNI       string `json:"dni"`
	Nombre    string `json:"nombre"`
	Apellido  string `json:"apellido"`
	Domicilio string `json:"domicilio"`
	FechaAlta string `json:"fecha_alta"`
}

type PacienteResponse struct {
	DNI       string `json:"dni"`
	Nombre    string `json:"nombre"`
	Apellido  string `json:"apellido"`
	Domicilio string `json:"domicilio"`
}
