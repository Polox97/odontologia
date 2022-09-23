package handler

import (
	"errors"
	"os"
	"strconv"

	"github.com/Polox97/odontologia/internal/domain"
	"github.com/Polox97/odontologia/internal/turno"
	"github.com/Polox97/odontologia/pkg/web"
	"github.com/gin-gonic/gin"
)

type turnoHandler struct {
	s turno.Service
}

// NewProductHandler crea un nuevo controller de turno
func NewTurnoHandler(s turno.Service) *turnoHandler {
	return &turnoHandler{
		s: s,
	}
}

// GetAll obtiene todos los turnos
func (h *turnoHandler) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		turnos, err := h.s.GetAll()
		if err != nil {
			web.Failure(c, 404, errors.New("turno not found"))
			return
		}
		web.Success(c, 200, turnos)
	}
}

// Get obtiene un turno por id
func (h *turnoHandler) GetByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		turno, err := h.s.GetByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("turno not found"))
			return
		}
		web.Success(c, 200, turno)
	}
}

// validateEmptys valida que los campos no esten vacios
func validateEmptysTurnos(turno *domain.Turno) (bool, error) {
	switch {
	case turno.PacienteID == 0 || turno.DentistaID == 0 || turno.Descripcion == "" || turno.FechaHora == "":
		return false, errors.New("fields can't be empty")
	}
	return true, nil
}

// Post crea un nuevo turno
func (h *turnoHandler) Post() gin.HandlerFunc {
	return func(c *gin.Context) {
		var turno domain.Turno
		token := c.GetHeader("TOKEN")
		if token == "" {
			web.Failure(c, 401, errors.New("token not found"))
			return
		}
		if token != os.Getenv("TOKEN") {
			web.Failure(c, 401, errors.New("invalid token"))
			return
		}
		err := c.ShouldBindJSON(&turno)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		valid, err := validateEmptysTurnos(&turno)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		p, err := h.s.Create(turno)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		web.Success(c, 201, p)
	}
}

// Delete elimina un turno
func (h *turnoHandler) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("TOKEN")
		if token == "" {
			web.Failure(c, 401, errors.New("token not found"))
			return
		}
		if token != os.Getenv("TOKEN") {
			web.Failure(c, 401, errors.New("invalid token"))
			return
		}
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		err = h.s.Delete(id)
		if err != nil {
			web.Failure(c, 404, err)
			return
		}
		web.Success(c, 204, nil)
	}
}

// Put actualiza un turno
func (h *turnoHandler) Put() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("TOKEN")
		if token == "" {
			web.Failure(c, 401, errors.New("token not found"))
			return
		}
		if token != os.Getenv("TOKEN") {
			web.Failure(c, 401, errors.New("invalid token"))
			return
		}
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		_, err = h.s.GetByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("turno not found"))
			return
		}
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		var turno domain.Turno
		err = c.ShouldBindJSON(&turno)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		valid, err := validateEmptysTurnos(&turno)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		p, err := h.s.Update(id, turno)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, p)
	}
}

// Patch actualiza un turno o alguno de sus campos
func (h *turnoHandler) Patch() gin.HandlerFunc {
	type Request struct {
		PacienteID  int    `json:"paciente_id,omitempty"`
		DentistaID  int    `json:"dentista_id,omitempty"`
		Descripcion string `json:"descripcion,omitempty"`
		FechaHora   string `json:"fecha_hora,omitempty"`
	}
	return func(c *gin.Context) {
		token := c.GetHeader("TOKEN")
		if token == "" {
			web.Failure(c, 401, errors.New("token not found"))
			return
		}
		if token != os.Getenv("TOKEN") {
			web.Failure(c, 401, errors.New("invalid token"))
			return
		}
		var r Request
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		_, err = h.s.GetByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("turno not found"))
			return
		}
		if err := c.ShouldBindJSON(&r); err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		update := domain.Turno{
			PacienteID:  r.PacienteID,
			DentistaID:  r.DentistaID,
			Descripcion: r.Descripcion,
			FechaHora:   r.FechaHora,
		}
		p, err := h.s.Update(id, update)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, p)
	}
}
