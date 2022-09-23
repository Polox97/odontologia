package turno

import (
	"errors"

	"github.com/Polox97/odontologia/internal/domain"
	"github.com/Polox97/odontologia/pkg/store"
)

type Repository interface {
	GetAll() ([]domain.Turno, error)
	GetByID(id int) (domain.Turno, error)
	Create(p domain.Turno) (domain.Turno, error)
	Update(id int, p domain.Turno) (domain.Turno, error)
	Delete(id int) error
}

type repository struct {
	storage store.StoreInterfaceTurno
}

func (r *repository) Create(t domain.Turno) (domain.Turno, error) {
	if !r.storage.ExistsPaciente(t.PacienteID) {
		return domain.Turno{}, errors.New("paciente no existe")
	}
	if !r.storage.ExistsDentista(t.DentistaID) {
		return domain.Turno{}, errors.New("dentista no existe")
	}
	err := r.storage.Create(t)
	if err != nil {
		return domain.Turno{}, errors.New("error creating turno")
	}
	return t, nil
}

func (r *repository) Delete(id int) error {
	err := r.storage.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) GetAll() ([]domain.Turno, error) {
	turnos, err := r.storage.ReadAll()
	if err != nil {
		return []domain.Turno{}, err
	}
	return turnos, nil
}

func (r *repository) GetByID(id int) (domain.Turno, error) {
	turno, err := r.storage.Read(id)
	if err != nil {
		return domain.Turno{}, errors.New("turno not found")
	}
	return turno, nil
}

func (r *repository) Update(id int, d domain.Turno) (domain.Turno, error) {
	if r.storage.Exists(d.ID) {
		return domain.Turno{}, errors.New("turno not exists")
	}
	err := r.storage.Update(d)
	if err != nil {
		return domain.Turno{}, errors.New("error updating turno")
	}
	return d, nil
}

func NewRepository(storage store.StoreInterfaceTurno) Repository {
	return &repository{storage}
}
