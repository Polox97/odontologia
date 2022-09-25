package service

import (
	dentistaI "github.com/Polox97/odontologia/interface/dentistainterface"
	dentistaUCI "github.com/Polox97/odontologia/interface/dentistaucinterface"
	dentistaModel "github.com/Polox97/odontologia/model/dentista"
)

type dentistaService struct {
	dentistaI.DentistaI
}

// NewService crea un nuevo servicio
func NewDentistaService(dentistaInterface dentistaI.DentistaI) dentistaUCI.DentistaUCI {
	return &dentistaService{
		dentistaInterface,
	}
}

func (s *dentistaService) GetAllDentistas() ([]dentistaModel.Dentista, error) {
	d, err := s.GetAll()
	if err != nil {
		return []dentistaModel.Dentista{}, err
	}
	return d, nil
}

func (s *dentistaService) GetDentistaByID(id int) (dentistaModel.Dentista, error) {
	d, err := s.GetByID(id)
	if err != nil {
		return dentistaModel.Dentista{}, err
	}
	return d, nil
}

func (s *dentistaService) CreateDentista(d dentistaModel.Dentista) (dentistaModel.Dentista, error) {
	d, err := s.Create(d)
	if err != nil {
		return dentistaModel.Dentista{}, err
	}
	return d, nil
}
func (s *dentistaService) UpdateDentista(id int, u dentistaModel.Dentista) (dentistaModel.Dentista, error) {
	d, err := s.GetByID(id)
	if err != nil {
		return dentistaModel.Dentista{}, err
	}
	if u.Matricula != "" {
		d.Matricula = u.Matricula
	}
	if u.Nombre != "" {
		d.Nombre = u.Nombre
	}
	if u.Apellido != "" {
		d.Apellido = u.Apellido
	}
	d, err = s.Update(id, d)
	if err != nil {
		return dentistaModel.Dentista{}, err
	}
	return d, nil
}

func (s *dentistaService) DeleteDentista(id int) error {
	err := s.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
