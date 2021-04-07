package handlers

import (
	"net/http"

	"github.com/dionarya23/sipgri-backend/guru"
	"github.com/dionarya23/sipgri-backend/helper"
	"github.com/gin-gonic/gin"
)

type guruHandler struct {
	guruService guru.Service
}

func NewGuruHandler(guruService guru.Service) *guruHandler {
	return &guruHandler{guruService}
}

func (h *guruHandler) RegisterGuru(c *gin.Context) {
	var input guru.RegisterGuruInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Register Guru failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
	}

	newGuru, err := h.guruService.RegisterGuru(input)
	if err != nil {
		response := helper.APIResponse("Register Guru failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
	}

	formatter := guru.FormatGuru(newGuru)
	response := helper.APIResponse("Success create guru", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)

}
