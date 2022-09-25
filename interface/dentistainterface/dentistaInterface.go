package dentistainterface

import (
	dentistaModel "github.com/Polox97/odontologia/model/dentista"
)

type DentistaI interface {
	GetAll() ([]dentistaModel.Dentista, error)
	GetByID(id int) (dentistaModel.Dentista, error)
	Create(p dentistaModel.Dentista) (dentistaModel.Dentista, error)
	Update(id int, p dentistaModel.Dentista) (dentistaModel.Dentista, error)
	Delete(id int) error
}