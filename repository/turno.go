package repository

import (
	"errors"

	turnoI "github.com/Polox97/odontologia/interface/turnointerface"
	turnoModel "github.com/Polox97/odontologia/model/turno"
	"github.com/Polox97/odontologia/pkg/store"
)


type turnoRepository struct {
	storage store.StoreInterfaceTurno
}

func NewRepository(storage store.StoreInterfaceTurno) turnoI.TurnoI {
	return &turnoRepository{storage}
}

func (r *turnoRepository) Create(t turnoModel.Turno) (turnoModel.Turno, error) {
	if !r.storage.ExistsPaciente(t.PacienteID) {
		return turnoModel.Turno{}, errors.New("paciente no existe")
	}
	if !r.storage.ExistsDentista(t.DentistaID) {
		return turnoModel.Turno{}, errors.New("dentista no existe")
	}
	err := r.storage.Create(t)
	if err != nil {
		return turnoModel.Turno{}, errors.New("error creating turno")
	}
	return t, nil
}

func (r *turnoRepository) Delete(id int) error {
	err := r.storage.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (r *turnoRepository) GetAll() ([]turnoModel.Turno, error) {
	turnos, err := r.storage.ReadAll()
	if err != nil {
		return []turnoModel.Turno{}, err
	}
	return turnos, nil
}

func (r *turnoRepository) GetByID(id int) (turnoModel.Turno, error) {
	turno, err := r.storage.Read(id)
	if err != nil {
		return turnoModel.Turno{}, errors.New("turno not found")
	}
	return turno, nil
}

func (r *turnoRepository) GetPaciente(dni string) ([]turnoModel.TurnoResponse, error) {
	turnos, err := r.storage.ReadPaciente(dni)
	if err != nil {
		return []turnoModel.TurnoResponse{}, err
	}
	return turnos, nil
}

func (r *turnoRepository) Update(id int, d turnoModel.Turno) (turnoModel.Turno, error) {
	if r.storage.Exists(d.ID) {
		return turnoModel.Turno{}, errors.New("turno not exists")
	}
	err := r.storage.Update(d)
	if err != nil {
		return turnoModel.Turno{}, errors.New("error updating turno")
	}
	return d, nil
}
