package paciente

import (
	"github.com/Polox97/odontologia/internal/domain"
)

type Service interface {
	GetAll() ([]domain.Paciente, error)
	// GetByID busca un paciente por su id
	GetByID(id int) (domain.Paciente, error)
	// Create agrega un nuevo paciente
	Create(p domain.Paciente) (domain.Paciente, error)
	// Delete elimina un paciente
	Delete(id int) error
	// Update actualiza un paciente
	Update(id int, p domain.Paciente) (domain.Paciente, error)
}

type service struct {
	r Repository
}

// NewService crea un nuevo servicio
func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) GetAll() ([]domain.Paciente, error) {
	p, err := s.r.GetAll()
	if err != nil {
		return []domain.Paciente{}, err
	}
	return p, nil
}

func (s *service) GetByID(id int) (domain.Paciente, error) {
	p, err := s.r.GetByID(id)
	if err != nil {
		return domain.Paciente{}, err
	}
	return p, nil
}

func (s *service) Create(d domain.Paciente) (domain.Paciente, error) {
	p, err := s.r.Create(d)
	if err != nil {
		return domain.Paciente{}, err
	}
	return p, nil
}
func (s *service) Update(id int, u domain.Paciente) (domain.Paciente, error) {
	p, err := s.r.GetByID(id)
	if err != nil {
		return domain.Paciente{}, err
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
	p, err = s.r.Update(id, p)
	if err != nil {
		return domain.Paciente{}, err
	}
	return p, nil
}

func (s *service) Delete(id int) error {
	err := s.r.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
