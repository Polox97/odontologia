package dentista

import (
	"github.com/Polox97/odontologia/internal/domain"
)

type Service interface {
	GetAll() ([]domain.Dentista, error)
	// GetByID busca un dentista por su id
	GetByID(id int) (domain.Dentista, error)
	// Create agrega un nuevo dentista
	Create(p domain.Dentista) (domain.Dentista, error)
	// Delete elimina un dentista
	Delete(id int) error
	// Update actualiza un dentista
	Update(id int, p domain.Dentista) (domain.Dentista, error)
}

type service struct {
	r Repository
}

// NewService crea un nuevo servicio
func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) GetAll() ([]domain.Dentista, error) {
	d, err := s.r.GetAll()
	if err != nil {
		return []domain.Dentista{}, err
	}
	return d, nil
}

func (s *service) GetByID(id int) (domain.Dentista, error) {
	d, err := s.r.GetByID(id)
	if err != nil {
		return domain.Dentista{}, err
	}
	return d, nil
}

func (s *service) Create(d domain.Dentista) (domain.Dentista, error) {
	d, err := s.r.Create(d)
	if err != nil {
		return domain.Dentista{}, err
	}
	return d, nil
}
func (s *service) Update(id int, u domain.Dentista) (domain.Dentista, error) {
	d, err := s.r.GetByID(id)
	if err != nil {
		return domain.Dentista{}, err
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
	d, err = s.r.Update(id, d)
	if err != nil {
		return domain.Dentista{}, err
	}
	return d, nil
}

func (s *service) Delete(id int) error {
	err := s.r.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
