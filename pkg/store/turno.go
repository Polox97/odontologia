package store

import (
	"database/sql"

	turnoModel "github.com/Polox97/odontologia/model/turno"
)

type StoreInterfaceTurno interface {
	// Read devuelve un turno por su id
	Read(id int) (turnoModel.Turno, error)
	// Read devuelve los turnos de un paciente por su dni
	ReadPaciente(dni string) ([]turnoModel.TurnoResponse, error)
	// ReadAll devuelve todos los turnos
	ReadAll() ([]turnoModel.Turno, error)
	// Create agrega un nuevo turno
	Create(paciente turnoModel.Turno) error
	// Update actualiza un turno
	Update(paciente turnoModel.Turno) error
	// Delete elimina un turno
	Delete(id int) error
	// Exists verifica si un turno existe
	Exists(idTurno int) bool
	// Exists verifica si un paciente existe
	ExistsPaciente(id int) bool
	// Exists verifica si un dentista existe
	ExistsDentista(id int) bool
}

type sqlStoreT struct {
	db *sql.DB
}

func NewSqlStoreTurno(db *sql.DB) StoreInterfaceTurno {
	return &sqlStoreT{
		db: db,
	}
}

func (s *sqlStoreT) Read(id int) (turnoModel.Turno, error) {
	var turno turnoModel.Turno
	query := "SELECT * FROM turnos WHERE id = ?;"
	row := s.db.QueryRow(query, id)
	err := row.Scan(&turno.ID, &turno.PacienteID, &turno.DentistaID, &turno.FechaHora, &turno.Descripcion)
	if err != nil {
		return turnoModel.Turno{}, err
	}
	return turno, nil
}

func (s *sqlStoreT) ReadPaciente(dni string) ([]turnoModel.TurnoResponse, error) {
	query := "SELECT t.id, t.fecha_hora, t.descripcion, p.dni, p.nombre, p.apellido, p.domicilio, d.matricula, d.nombre, d.apellido from turnos t INNER JOIN pacientes p on p.id = t.paciente_id INNER JOIN dentistas d on d.id = t.dentista_id WHERE p.dni = ?;"
	rows, err := s.db.Query(query, dni)
	var turnos []turnoModel.TurnoResponse

	for rows.Next() {
		turno := turnoModel.TurnoResponse{}
		err = rows.Scan(
			&turno.ID,
			&turno.FechaHora, 
			&turno.Descripcion, 
			&turno.PacienteID.DNI, 
			&turno.PacienteID.Nombre, 
			&turno.PacienteID.Apellido, 
			&turno.PacienteID.Domicilio,
			&turno.DentistaID.Matricula, 
			&turno.DentistaID.Nombre, 
			&turno.DentistaID.Apellido)

		turnos = append(turnos, turno)
	}
	if err != nil {
		return []turnoModel.TurnoResponse{}, err
	}
	return turnos, nil
}

func (s *sqlStoreT) ReadAll() ([]turnoModel.Turno, error) {
	query := "SELECT * FROM turnos"
	rows, err := s.db.Query(query)
	var turnos []turnoModel.Turno

	for rows.Next() {
		turno := turnoModel.Turno{}
		err = rows.Scan(&turno.ID, &turno.PacienteID, &turno.DentistaID, &turno.FechaHora, &turno.Descripcion)
		turnos = append(turnos, turno)
	}
	if err != nil {
		return []turnoModel.Turno{}, err
	}
	return turnos, nil
}

func (s *sqlStoreT) Create(turno turnoModel.Turno) error {
	query := "INSERT INTO turnos (paciente_id, dentista_id, fecha_hora, descripcion) VALUES (?, ?, ?, ?);"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return err
	}
	res, err := stmt.Exec(turno.PacienteID, turno.DentistaID, turno.FechaHora, turno.Descripcion)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

func (s *sqlStoreT) Update(turno turnoModel.Turno) error {
	query := "UPDATE turnos SET paciente_id = ?, dentista_id = ?, fecha_hora = ?, descripcion = ? WHERE id = ?;"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return err
	}
	res, err := stmt.Exec(turno.PacienteID, turno.DentistaID, turno.FechaHora, turno.Descripcion, turno.ID)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

func (s *sqlStoreT) Delete(id int) error {
	query := "DELETE FROM turnos WHERE id = ?;"
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

func (s *sqlStoreT) Exists(idTurno int) bool {
	var exists bool
	var id int
	query := "SELECT id FROM turnos WHERE id = ?;"
	row := s.db.QueryRow(query, idTurno)
	err := row.Scan(&id)
	if err != nil {
		return false
	}
	if id > 0 {
		exists = true
	}
	return exists
}

func (s *sqlStoreT) ExistsPaciente(idPaciente int) bool {
	var exists bool
	var id int
	query := "SELECT id FROM pacientes WHERE id = ?;"
	row := s.db.QueryRow(query, idPaciente)
	err := row.Scan(&id)
	if err != nil {
		return false
	}
	if id > 0 {
		exists = true
	}
	return exists
}

func (s *sqlStoreT) ExistsDentista(idDentista int) bool {
	var exists bool
	var id int
	query := "SELECT id FROM dentistas WHERE id = ?;"
	row := s.db.QueryRow(query, idDentista)
	err := row.Scan(&id)
	if err != nil {
		return false
	}
	if id > 0 {
		exists = true
	}
	return exists
}
