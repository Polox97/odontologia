package paciente

import (
	"errors"

	"github.com/Polox97/odontologia/internal/domain"
	"github.com/Polox97/odontologia/pkg/store"
)

type Repository interface {
	GetAll() ([]domain.Paciente, error)
	GetByID(id int) (domain.Paciente, error)
	Create(p domain.Paciente) (domain.Paciente, error)
	Update(id int, p domain.Paciente) (domain.Paciente, error)
	Delete(id int) error
}

type repository struct {
	storage store.StoreInterfacePaciente
}

func (r *repository) Create(d domain.Paciente) (domain.Paciente, error) {
	if r.storage.Exists(d.DNI) {
		return domain.Paciente{}, errors.New("dni already exists")
	}
	err := r.storage.Create(d)
	if err != nil {
		return domain.Paciente{}, errors.New("error creating patient")
	}
	return d, nil
}

func (r *repository) Delete(id int) error {
	err := r.storage.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) GetAll() ([]domain.Paciente, error) {
	pacientes, err := r.storage.ReadAll()
	if err != nil {
		return []domain.Paciente{}, err
	}
	return pacientes, nil
}

func (r *repository) GetByID(id int) (domain.Paciente, error) {
	paciente, err := r.storage.Read(id)
	if err != nil {
		return domain.Paciente{}, errors.New("patient not found")
	}
	return paciente, nil
}

func (r *repository) Update(id int, d domain.Paciente) (domain.Paciente, error) {
	if r.storage.Exists(d.DNI) {
		return domain.Paciente{}, errors.New("dni not exists")
	}
	err := r.storage.Update(d)
	if err != nil {
		return domain.Paciente{}, errors.New("error updating patient")
	}
	return d, nil
}

func NewRepository(storage store.StoreInterfacePaciente) Repository {
	return &repository{storage}
}
