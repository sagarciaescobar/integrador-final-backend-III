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

func (h patientHandler) GetByID() gin.HandlerFunc {
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

func (h patientHandler) Add() gin.HandlerFunc {
	return func(c *gin.Context) {
		var m map[string]string
		if err := c.BindJSON(&m); err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		p, err := patient.Mapper(m)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		res, err := h.s.Create(p)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		web.Success(c, 201, res)
		return
	}
}

func (h patientHandler) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		var m map[string]string
		if err := c.BindJSON(&m); err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		if m["id"] == "" {
			web.Failure(c, 400, errors.New("id must be defined"))
			return
		}
		p, err := patient.Mapper(m)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		res, err := h.s.UpdateById(p)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		web.Success(c, 200, res)
		return
	}
}

func (h patientHandler) ChangeAddresById() gin.HandlerFunc {
	return func(c *gin.Context) {
		var m map[string]string
		if err := c.BindJSON(&m); err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		if m["id"] == "" {
			web.Failure(c, 400, errors.New("id must be defined"))
			return
		}
		if m["address"] == "" {
			web.Failure(c, 400, errors.New("address must be defined"))
			return
		}
		id, err := strconv.Atoi(m["id"])
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		err = h.s.ChangeAddresById(id, m["address"])
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		web.Success(c, 206, nil)
		return
	}
}

func (h patientHandler) Delete() gin.HandlerFunc {
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
