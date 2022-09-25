package store

import (
	"database/sql"

	pacienteModel "github.com/Polox97/odontologia/model/paciente"
)

type StoreInterfacePaciente interface {
	// Read devuelve un paciente por su id
	Read(id int) (pacienteModel.Paciente, error)
	// ReadAll devuelve todos los pacientes
	ReadAll() ([]pacienteModel.Paciente, error)
	// Create agrega un nuevo paciente
	Create(paciente pacienteModel.Paciente) error
	// Update actualiza un paciente
	Update(paciente pacienteModel.Paciente) error
	// Delete elimina un paciente
	Delete(id int) error
	// Exists verifica si un paciente existe
	Exists(dni string) bool
}

type sqlStoreP struct {
	db *sql.DB
}

func NewSqlStorePaciente(db *sql.DB) StoreInterfacePaciente {
	return &sqlStoreP{
		db: db,
	}
}

func (s *sqlStoreP) Read(id int) (pacienteModel.Paciente, error) {
	var paciente pacienteModel.Paciente
	query := "SELECT * FROM pacientes WHERE id = ?;"
	row := s.db.QueryRow(query, id)
	err := row.Scan(&paciente.ID, &paciente.DNI, &paciente.Nombre, &paciente.Apellido, &paciente.Domicilio, &paciente.FechaAlta)
	if err != nil {
		return pacienteModel.Paciente{}, err
	}
	return paciente, nil
}

func (s *sqlStoreP) ReadAll() ([]pacienteModel.Paciente, error) {
	query := "SELECT * FROM pacientes"
	rows, err := s.db.Query(query)
	var pacientes []pacienteModel.Paciente

	for rows.Next() {
		paciente := pacienteModel.Paciente{}
		err = rows.Scan(&paciente.ID, &paciente.DNI, &paciente.Nombre, &paciente.Apellido, &paciente.Domicilio, &paciente.FechaAlta)
		pacientes = append(pacientes, paciente)
	}
	if err != nil {
		return []pacienteModel.Paciente{}, err
	}
	return pacientes, nil
}

func (s *sqlStoreP) Create(paciente pacienteModel.Paciente) error {
	query := "INSERT INTO pacientes (dni, nombre, apellido, domicilio) VALUES (?, ?, ?, ?);"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return err
	}
	res, err := stmt.Exec(paciente.DNI, paciente.Nombre, paciente.Apellido, paciente.Domicilio)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

func (s *sqlStoreP) Update(paciente pacienteModel.Paciente) error {
	query := "UPDATE pacientes SET dni = ?, nombre = ?, apellido = ?, domicilio = ?, fecha_alta = ? WHERE id = ?;"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return err
	}
	res, err := stmt.Exec(paciente.DNI, paciente.Nombre, paciente.Apellido, paciente.Domicilio, paciente.FechaAlta, paciente.ID)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

func (s *sqlStoreP) Delete(id int) error {
	query := "DELETE FROM pacientes WHERE id = ?;"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return err
	}
	res, err := stmt.Exec(id)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

func (s *sqlStoreP) Exists(dni string) bool {
	var exists bool
	var id int
	query := "SELECT id FROM dentistas WHERE dni = ?;"
	row := s.db.QueryRow(query, dni)
	err := row.Scan(&id)
	if err != nil {
		return false
	}
	if id > 0 {
		exists = true
	}
	return exists
}
