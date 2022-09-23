package turno

import (
	"github.com/Polox97/odontologia/internal/domain"
)

type Service interface {
	GetAll() ([]domain.Turno, error)
	// GetByID busca un paciente por su id
	GetByID(id int) (domain.Turno, error)
	// Create agrega un nuevo paciente
	Create(p domain.Turno) (domain.Turno, error)
	// Delete elimina un paciente
	Delete(id int) error
	// Update actualiza un paciente
	Update(id int, p domain.Turno) (domain.Turno, error)
}

type service struct {
	r Repository
}

// NewService crea un nuevo servicio
func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) GetAll() ([]domain.Turno, error) {
	p, err := s.r.GetAll()
	if err != nil {
		return []domain.Turno{}, err
	}
	return p, nil
}

func (s *service) GetByID(id int) (domain.Turno, error) {
	p, err := s.r.GetByID(id)
	if err != nil {
		return domain.Turno{}, err
	}
	return p, nil
}

func (s *service) Create(d domain.Turno) (domain.Turno, error) {
	p, err := s.r.Create(d)
	if err != nil {
		return domain.Turno{}, err
	}
	return p, nil
}
func (s *service) Update(id int, u domain.Turno) (domain.Turno, error) {
	p, err := s.r.GetByID(id)
	if err != nil {
		return domain.Turno{}, err
	}
	if u.PacienteID > 0 {
		p.PacienteID = u.PacienteID
	}
	if u.DentistaID > 0 {
		p.DentistaID = u.DentistaID
	}
	if u.Descripcion != "" {
		p.Descripcion = u.Descripcion
	}
	if u.FechaHora != "" {
		p.FechaHora = u.FechaHora
	}
	p, err = s.r.Update(id, p)
	if err != nil {
		return domain.Turno{}, err
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
