package dentista

import (
	"errors"

	"github.com/Polox97/odontologia/internal/domain"
	"github.com/Polox97/odontologia/pkg/store/dentista"
)

type Repository interface {
	GetAll() ([]domain.Dentista, error)
	GetByID(id int) (domain.Dentista, error)
	Create(p domain.Dentista) (domain.Dentista, error)
	Update(id int, p domain.Dentista) (domain.Dentista, error)
	Delete(id int) error
}

type repository struct {
	storage store.StoreInterface
}

func (r *repository) Create(d domain.Dentista) (domain.Dentista, error) {
	if r.storage.Exists(d.Matricula) {
		return domain.Dentista{}, errors.New("code value already exists")
	}
	err := r.storage.Create(d)
	if err != nil {
		return domain.Dentista{}, errors.New("error creating dentist")
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

func (r *repository) GetAll() ([]domain.Dentista, error) {
	dentistas, err := r.storage.ReadAll()
	if err != nil {
		return []domain.Dentista{}, err
	}
	return dentistas, nil
}

func (r *repository) GetByID(id int) (domain.Dentista, error) {
	dentista, err := r.storage.Read(id)
	if err != nil {
		return domain.Dentista{}, errors.New("dentist not found")
	}
	return dentista, nil
}

func (r *repository) Update(id int, d domain.Dentista) (domain.Dentista, error) {
	if !r.storage.Exists(d.Matricula) {
		return domain.Dentista{}, errors.New("code value not exists")
	}
	err := r.storage.Update(d)
	if err != nil {
		return domain.Dentista{}, errors.New("error updating dentist")
	}
	return d, nil
}

func NewRepository(storage store.StoreInterface) Repository {
	return &repository{storage}
}
