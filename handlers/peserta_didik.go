package handlers

import (
	"net/http"

	"github.com/dionarya23/sipgri-backend/helper"
	"github.com/dionarya23/sipgri-backend/peserta_didik"
	"github.com/gin-gonic/gin"
)

type pesertaDidikHandler struct {
	pesertaDidikService peserta_didik.Service
}

func NewPesertaDidikHandler(pesertaDidikService peserta_didik.Service) *pesertaDidikHandler {
	return &pesertaDidikHandler{pesertaDidikService}
}

func (h *pesertaDidikHandler) CreatePesertaDidik(c *gin.Context) {
	var input peserta_didik.InputDataPesertaDidik

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Create peserta didik failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newSiswa, err := h.pesertaDidikService.CreatePesertaDidik(input)
	if err != nil {
		response := helper.APIResponse("Create peserta didik failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := peserta_didik.FormatPesertaDidikDetail(newSiswa)
	response := helper.APIResponse("Success create peserta didik", http.StatusCreated, "success", formatter)
	c.JSON(http.StatusCreated, response)
}
