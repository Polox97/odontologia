package service

import (
	turnoI "github.com/Polox97/odontologia/interface/turnointerface"
	turnoUCI "github.com/Polox97/odontologia/interface/turnoucinterface"
	tuenoModel "github.com/Polox97/odontologia/model/turno"
)

type turnoService struct {
	turnoI.TurnoI
}

// NewService crea un nuevo servicio
func NewService(turnoInterface turnoI.TurnoI) turnoUCI.TurnoUCI {
	return &turnoService{
		turnoInterface,
	}
}

func (s *turnoService) GetAllTurnos() ([]tuenoModel.Turno, error) {
	t, err := s.GetAll()
	if err != nil {
		return []tuenoModel.Turno{}, err
	}
	return t, nil
}

func (s *turnoService) GetTurnoByID(id int) (tuenoModel.Turno, error) {
	t, err := s.GetByID(id)
	if err != nil {
		return tuenoModel.Turno{}, err
	}
	return t, nil
}

func (s *turnoService) GetTurnoPaciente(dni string) ([]tuenoModel.TurnoResponse, error) {
	t, err := s.GetPaciente(dni)
	if err != nil {
		return []tuenoModel.TurnoResponse{}, err
	}
	return t, nil
}

func (s *turnoService) CreateTurno(d tuenoModel.Turno) (tuenoModel.Turno, error) {
	t, err := s.Create(d)
	if err != nil {
		return tuenoModel.Turno{}, err
	}
	return t, nil
}
func (s *turnoService) UpdateTurno(id int, u tuenoModel.Turno) (tuenoModel.Turno, error) {
	t, err := s.GetByID(id)
	if err != nil {
		return tuenoModel.Turno{}, err
	}
	if u.PacienteID > 0 {
		t.PacienteID = u.PacienteID
	}
	if u.DentistaID > 0 {
		t.DentistaID = u.DentistaID
	}
	if u.Descripcion != "" {
		t.Descripcion = u.Descripcion
	}
	if u.FechaHora != "" {
		t.FechaHora = u.FechaHora
	}
	t, err = s.Update(id, t)
	if err != nil {
		return tuenoModel.Turno{}, err
	}
	return t, nil
}

func (s *turnoService) DeleteTurno(id int) error {
	err := s.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
