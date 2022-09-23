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
		id, err := strconv.ParseInt(idParam, 0, 0)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		patient, err := h.s.GetByID(int(id))
		if err != nil {
			web.Failure(c, 404, errors.New("product not found"))
			return
		}
		web.Success(c, 200, patient)
	}
}
