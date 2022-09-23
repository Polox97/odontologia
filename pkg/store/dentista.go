package store

import (
	"database/sql"

	"github.com/Polox97/odontologia/internal/domain"
)

type StoreInterfaceDentista interface {
	// Read devuelve un dentista por su id
	Read(id int) (domain.Dentista, error)
	// ReadAll devuelve todos los dentistas
	ReadAll() ([]domain.Dentista, error)
	// Create agrega un nuevo dentista
	Create(dentista domain.Dentista) error
	// Update actualiza un dentista
	Update(dentista domain.Dentista) error
	// Delete elimina un dentista
	Delete(id int) error
	// Exists verifica si un dentista existe
	Exists(matricula string) bool
}

type sqlStoreD struct {
	db *sql.DB
}

func NewSqlStoreDentista(db *sql.DB) StoreInterfaceDentista {
	return &sqlStoreD{
		db: db,
	}
}

func (s *sqlStoreD) Read(id int) (domain.Dentista, error) {
	var dentista domain.Dentista
	query := "SELECT * FROM dentistas WHERE id = ?;"
	row := s.db.QueryRow(query, id)
	err := row.Scan(&dentista.ID, &dentista.Matricula, &dentista.Nombre, &dentista.Apellido)
	if err != nil {
		return domain.Dentista{}, err
	}
	return dentista, nil
}

func (s *sqlStoreD) ReadAll() ([]domain.Dentista, error) {
	query := "SELECT * FROM dentistas"
	rows, err := s.db.Query(query)
	var dentistas []domain.Dentista

	for rows.Next() {
		dentista := domain.Dentista{}
		err = rows.Scan(&dentista.ID, &dentista.Matricula, &dentista.Nombre, &dentista.Apellido)
		dentistas = append(dentistas, dentista)
	}
	if err != nil {
		return []domain.Dentista{}, err
	}
	return dentistas, nil
}

func (s *sqlStoreD) Create(dentista domain.Dentista) error {
	query := "INSERT INTO dentistas (matricula, nombre, apellido) VALUES (?, ?, ?);"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return err
	}
	res, err := stmt.Exec(dentista.Matricula, dentista.Nombre, dentista.Apellido)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

func (s *sqlStoreD) Update(dentista domain.Dentista) error {
	query := "UPDATE dentistas SET matricula = ?, nombre = ?, apellido = ? WHERE id = ?;"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return err
	}
	res, err := stmt.Exec(dentista.Matricula, dentista.Nombre, dentista.Apellido, dentista.ID)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

func (s *sqlStoreD) Delete(id int) error {
	query := "DELETE FROM dentistas WHERE id = ?;"
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

func (s *sqlStoreD) Exists(matricula string) bool {
	var exists bool
	var id int
	query := "SELECT id FROM dentistas WHERE matricula = ?;"
	row := s.db.QueryRow(query, matricula)
	err := row.Scan(&id)
	if err != nil {
		return false
	}
	if id > 0 {
		exists = true
	}
	return exists
}
