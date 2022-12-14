package handler

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sagarciaescobar/integrador-final-backend-III/internal/dentist"
	"github.com/sagarciaescobar/integrador-final-backend-III/pkg/web"
)

type dentistHandler struct {
	s dentist.Service
}

func NewDentistHandler() *dentistHandler {
	return &dentistHandler{s: dentist.NewService()}
}

func (h dentistHandler) DGetByID() gin.HandlerFunc {
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

func (h dentistHandler) DAdd() gin.HandlerFunc {
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

func (h dentistHandler) DUpdate() gin.HandlerFunc {
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

func (h dentistHandler) DChangeRegistrationIdById() gin.HandlerFunc {
	return func(c *gin.Context) {
		var m map[string]string
		if err := c.BindJSON(&m); err != nil {
			web.Failure(c, 400, errors.New("invalid request"))
			return
		}
		err := h.s.ChangeRegistrationIdById(m)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		web.Success(c, 206, nil)
		return
	}
}

func (h dentistHandler) DDelete() gin.HandlerFunc {
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
