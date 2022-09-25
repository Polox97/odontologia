package service

import (
	pacienteI "github.com/Polox97/odontologia/interface/pacienteinterface"
	pacienteUCI "github.com/Polox97/odontologia/interface/pacienteucinterface"
	pacienteModel "github.com/Polox97/odontologia/model/paciente"
)

type paceinteService struct {
	pacienteI.PacienteI
}

// NewService crea un nuevo servicio
func NewPacienteService(pacienteInterface pacienteI.PacienteI) pacienteUCI.PacienteUCI {
	return &paceinteService{
		pacienteInterface,
	}
}

func (s *paceinteService) GetAllPacientes() ([]pacienteModel.Paciente, error) {
	p, err := s.GetAll()
	if err != nil {
		return []pacienteModel.Paciente{}, err
	}
	return p, nil
}

func (s *paceinteService) GetPacienteByID(id int) (pacienteModel.Paciente, error) {
	p, err := s.GetByID(id)
	if err != nil {
		return pacienteModel.Paciente{}, err
	}
	return p, nil
}

func (s *paceinteService) CreatePaciente(d pacienteModel.Paciente) (pacienteModel.Paciente, error) {
	p, err := s.Create(d)
	if err != nil {
		return pacienteModel.Paciente{}, err
	}
	return p, nil
}
func (s *paceinteService) UpdatePaciente(id int, u pacienteModel.Paciente) (pacienteModel.Paciente, error) {
	p, err := s.GetByID(id)
	if err != nil {
		return pacienteModel.Paciente{}, err
	}
	if u.DNI != "" {
		p.DNI = u.DNI
	}
	if u.Nombre != "" {
		p.Nombre = u.Nombre
	}
	if u.Apellido != "" {
		p.Apellido = u.Apellido
	}
	if u.Domicilio != "" {
		p.Domicilio = u.Domicilio
	}
	p, err = s.Update(id, p)
	if err != nil {
		return pacienteModel.Paciente{}, err
	}
	return p, nil
}

func (s *paceinteService) DeletePaciente(id int) error {
	err := s.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
