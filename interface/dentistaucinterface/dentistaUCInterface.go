package dentistaucinterface

import (
	dentistaModel "github.com/Polox97/odontologia/model/dentista"
)

type DentistaUCI interface {
	GetAllDentistas() ([]dentistaModel.Dentista, error)
	GetDentistaByID(id int) (dentistaModel.Dentista, error)
	CreateDentista(p dentistaModel.Dentista) (dentistaModel.Dentista, error)
	DeleteDentista(id int) error
	UpdateDentista(id int, p dentistaModel.Dentista) (dentistaModel.Dentista, error)
}