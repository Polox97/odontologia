package repository

import (
	"errors"

	dentistaI "github.com/Polox97/odontologia/interface/dentistainterface"
	dentistaModel "github.com/Polox97/odontologia/model/dentista"
	"github.com/Polox97/odontologia/pkg/store"
)

type dentistaRepository struct {
	storage store.StoreInterfaceDentista
}

func NewDentistaRepo(storage store.StoreInterfaceDentista) dentistaI.DentistaI {
	return &dentistaRepository{storage}
}

func (r *dentistaRepository) Create(d dentistaModel.Dentista) (dentistaModel.Dentista, error) {
	if r.storage.Exists(d.Matricula) {
		return dentistaModel.Dentista{}, errors.New("code value already exists")
	}
	err := r.storage.Create(d)
	if err != nil {
		return dentistaModel.Dentista{}, errors.New("error creating dentist")
	}
	return d, nil
}

func (r *dentistaRepository) Delete(id int) error {
	err := r.storage.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (r *dentistaRepository) GetAll() ([]dentistaModel.Dentista, error) {
	dentistas, err := r.storage.ReadAll()
	if err != nil {
		return []dentistaModel.Dentista{}, err
	}
	return dentistas, nil
}

func (r *dentistaRepository) GetByID(id int) (dentistaModel.Dentista, error) {
	dentista, err := r.storage.Read(id)
	if err != nil {
		return dentistaModel.Dentista{}, errors.New("dentist not found")
	}
	return dentista, nil
}

func (r *dentistaRepository) Update(id int, d dentistaModel.Dentista) (dentistaModel.Dentista, error) {
	if !r.storage.Exists(d.Matricula) {
		return dentistaModel.Dentista{}, errors.New("code value not exists")
	}
	err := r.storage.Update(d)
	if err != nil {
		return dentistaModel.Dentista{}, errors.New("error updating dentist")
	}
	return d, nil
}
