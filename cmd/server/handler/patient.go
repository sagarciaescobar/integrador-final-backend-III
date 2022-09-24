package handler

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sagarciaescobar/integrador-final-backend-III/internal/patient"
	"github.com/sagarciaescobar/integrador-final-backend-III/pkg/web"
)

type patientHandler struct {
	s patient.Service
}

func NewPatientHandler() *patientHandler {
	return &patientHandler{s: patient.NewService()}
}

func (h patientHandler) PGetByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		patient, err := h.s.GetByID(id)
		if err != nil {
			web.Failure(c, 404, err)
			return
		}
		web.Success(c, 200, patient)
		return
	}
}

func (h patientHandler) PAdd() gin.HandlerFunc {
	return func(c *gin.Context) {
		var m map[string]string
		if err := c.BindJSON(&m); err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		res, err := h.s.Create(m)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		web.Success(c, 201, res)
		return
	}
}

func (h patientHandler) PUpdate() gin.HandlerFunc {
	return func(c *gin.Context) {
		var m map[string]string
		if err := c.BindJSON(&m); err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		res, err := h.s.UpdateById(m)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		web.Success(c, 200, res)
		return
	}
}

func (h patientHandler) PChangeAddresById() gin.HandlerFunc {
	return func(c *gin.Context) {
		var m map[string]string
		if err := c.BindJSON(&m); err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		err := h.s.ChangeAddresById(m)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		web.Success(c, 206, nil)
		return
	}
}

func (h patientHandler) PDelete() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		err = h.s.DeleteById(id)
		if err != nil {
			web.Failure(c, 404, err)
			return
		}
		web.Success(c, 204, nil)
		return
	}
}
