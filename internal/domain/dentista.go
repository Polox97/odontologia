package domain

type Dentista struct {
	ID        int    `json:"id"`
	Matricula string `json:"matricula"`
	Nombre    string `json:"nombre"`
	Apellido  string `json:"apellido"`
}

type DentistaResponse struct {
	Matricula string `json:"matricula"`
	Nombre    string `json:"nombre"`
	Apellido  string `json:"apellido"`
}
