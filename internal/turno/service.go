package turno

import (
	"github.com/Polox97/odontologia/internal/domain"
)

type Service interface {
	// GetAll busca todos los turnos
	GetAll() ([]domain.Turno, error)
	// GetByID busca un turno por su id
	GetByID(id int) (domain.Turno, error)
	// GetAll busca todos los turnos
	GetPaciente(dni string) ([]domain.TurnoResponse, error)
	// Create agrega un nuevo turno
	Create(p domain.Turno) (domain.Turno, error)
	// Delete elimina un turno
	Delete(id int) error
	// Update actualiza un turno
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
	t, err := s.r.GetAll()
	if err != nil {
		return []domain.Turno{}, err
	}
	return t, nil
}

func (s *service) GetByID(id int) (domain.Turno, error) {
	t, err := s.r.GetByID(id)
	if err != nil {
		return domain.Turno{}, err
	}
	return t, nil
}

func (s *service) GetPaciente(dni string) ([]domain.TurnoResponse, error) {
	t, err := s.r.GetPaciente(dni)
	if err != nil {
		return []domain.TurnoResponse{}, err
	}
	return t, nil
}

func (s *service) Create(d domain.Turno) (domain.Turno, error) {
	t, err := s.r.Create(d)
	if err != nil {
		return domain.Turno{}, err
	}
	return t, nil
}
func (s *service) Update(id int, u domain.Turno) (domain.Turno, error) {
	t, err := s.r.GetByID(id)
	if err != nil {
		return domain.Turno{}, err
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
	t, err = s.r.Update(id, t)
	if err != nil {
		return domain.Turno{}, err
	}
	return t, nil
}

func (s *service) Delete(id int) error {
	err := s.r.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
