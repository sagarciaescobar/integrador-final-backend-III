package handler

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sagarciaescobar/integrador-final-backend-III/internal/appointment"
	"github.com/sagarciaescobar/integrador-final-backend-III/pkg/web"
)

type appointmentHandler struct {
	s appointment.Service
}

func NewAppointmentHandler() *appointmentHandler {
	return &appointmentHandler{s: appointment.NewService()}
}

func (h appointmentHandler) AGetByID() gin.HandlerFunc {
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

func (h appointmentHandler) AAdd() gin.HandlerFunc {
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

func (h appointmentHandler) AUpdate() gin.HandlerFunc {
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

func (h appointmentHandler) AChangeTimeById() gin.HandlerFunc {
	return func(c *gin.Context) {
		var m map[string]string
		if err := c.BindJSON(&m); err != nil {
			web.Failure(c, 400, err)
			return
		}
		err := h.s.ChangeTimeById(m)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		web.Success(c, 206, nil)
		return
	}
}

func (h appointmentHandler) ADelete() gin.HandlerFunc {
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
