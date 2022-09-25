package handler

import (
	"errors"
	"os"
	"strconv"

	dentistaUCI "github.com/Polox97/odontologia/interface/dentistaucinterface"
	dentistaModel "github.com/Polox97/odontologia/model/dentista"
	"github.com/Polox97/odontologia/pkg/web"
	"github.com/gin-gonic/gin"
)

type dentistaHandler struct {
	dentistaUCI.DentistaUCI
}

// NewProductHandler crea un nuevo controller de dentista
func NewDentistaHandler(dentistaUCInterface dentistaUCI.DentistaUCI) *dentistaHandler {
	return &dentistaHandler{
		dentistaUCInterface,
	}
}
// GetAll obtiene todos los dentistas
func (h *dentistaHandler) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		dentistas, err := h.GetAllDentistas()
		if err != nil {
			web.Failure(c, 404, errors.New("dentist not found"))
			return
		}
		web.Success(c, 200, dentistas)
	}
}

// Get obtiene un dentista por id
func (h *dentistaHandler) GetByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		product, err := h.GetDentistaByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("dentist not found"))
			return
		}
		web.Success(c, 200, product)
	}
}

// validateEmptys valida que los campos no esten vacios
func validateEmptyDentista(dentista *dentistaModel.Dentista) (bool, error) {
	switch {
	case dentista.Matricula == "" || dentista.Nombre == "" || dentista.Apellido == "":
		return false, errors.New("fields can't be empty")
	}
	return true, nil
}

// Post crea un nuevo dentista
func (h *dentistaHandler) Post() gin.HandlerFunc {
	return func(c *gin.Context) {
		var dentista dentistaModel.Dentista
		token := c.GetHeader("TOKEN")
		if token == "" {
			web.Failure(c, 401, errors.New("token not found"))
			return
		}
		if token != os.Getenv("TOKEN") {
			web.Failure(c, 401, errors.New("invalid token"))
			return
		}
		err := c.ShouldBindJSON(&dentista)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		valid, err := validateEmptyDentista(&dentista)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		p, err := h.CreateDentista(dentista)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		web.Success(c, 201, p)
	}
}

// Delete elimina un dentista
func (h *dentistaHandler) Delete() gin.HandlerFunc {
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
		err = h.DeleteDentista(id)
		if err != nil {
			web.Failure(c, 404, err)
			return
		}
		web.Success(c, 204, nil)
	}
}

// Put actualiza un dentista
func (h *dentistaHandler) Put() gin.HandlerFunc {
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
		_, err = h.GetDentistaByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("dentist not found"))
			return
		}
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		var dentista dentistaModel.Dentista
		err = c.ShouldBindJSON(&dentista)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		valid, err := validateEmptyDentista(&dentista)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		p, err := h.UpdateDentista(id, dentista)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, p)
	}
}

// Patch actualiza un dentista o alguno de sus campos
func (h *dentistaHandler) Patch() gin.HandlerFunc {
	type Request struct {
		Matricula string `json:"matricula,omitempty"`
		Nombre    string `json:"nombre,omitempty"`
		Apellido  string `json:"apellido,omitempty"`
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
		_, err = h.GetDentistaByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("dentist not found"))
			return
		}
		if err := c.ShouldBindJSON(&r); err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		update := dentistaModel.Dentista{
			Matricula: r.Matricula,
			Nombre:    r.Nombre,
			Apellido:  r.Apellido,
		}
		p, err := h.UpdateDentista(id, update)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, p)
	}
}
