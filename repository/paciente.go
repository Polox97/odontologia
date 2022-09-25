package repository

import (
	"errors"

	pacienteI "github.com/Polox97/odontologia/interface/pacienteinterface"
	pacienteModel "github.com/Polox97/odontologia/model/paciente"
	"github.com/Polox97/odontologia/pkg/store"
)

type pacienteRepository struct {
	storage store.StoreInterfacePaciente
}

func NewPacienteRepository(storage store.StoreInterfacePaciente) pacienteI.PacienteI {
	return &pacienteRepository{storage}
}

func (r *pacienteRepository) Create(d pacienteModel.Paciente) (pacienteModel.Paciente, error) {
	if r.storage.Exists(d.DNI) {
		return pacienteModel.Paciente{}, errors.New("dni already exists")
	}
	err := r.storage.Create(d)
	if err != nil {
		return pacienteModel.Paciente{}, errors.New("error creating patient")
	}
	return d, nil
}

func (r *pacienteRepository) Delete(id int) error {
	err := r.storage.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (r *pacienteRepository) GetAll() ([]pacienteModel.Paciente, error) {
	pacientes, err := r.storage.ReadAll()
	if err != nil {
		return []pacienteModel.Paciente{}, err
	}
	return pacientes, nil
}

func (r *pacienteRepository) GetByID(id int) (pacienteModel.Paciente, error) {
	paciente, err := r.storage.Read(id)
	if err != nil {
		return pacienteModel.Paciente{}, errors.New("patient not found")
	}
	return paciente, nil
}

func (r *pacienteRepository) Update(id int, d pacienteModel.Paciente) (pacienteModel.Paciente, error) {
	if r.storage.Exists(d.DNI) {
		return pacienteModel.Paciente{}, errors.New("dni not exists")
	}
	err := r.storage.Update(d)
	if err != nil {
		return pacienteModel.Paciente{}, errors.New("error updating patient")
	}
	return d, nil
}
