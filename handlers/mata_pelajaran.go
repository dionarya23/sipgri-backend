package handlers

import (
	"net/http"

	"github.com/dionarya23/sipgri-backend/helper"
	"github.com/dionarya23/sipgri-backend/mata_pelajaran"
	"github.com/gin-gonic/gin"
)

type mata_pelajaranHandler struct {
	mata_pelajaranService mata_pelajaran.Service
}

func NewMataPelajaranHandler(mata_pelajaranService mata_pelajaran.Service) *mata_pelajaranHandler {
	return &mata_pelajaranHandler{mata_pelajaranService}
}

func (h *mata_pelajaranHandler) CreateNewMataPelajaran(c *gin.Context) {
	var input mata_pelajaran.InputNewMataPelajaran

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Create mata pelajaran failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	mataPelajaran, err_ := h.mata_pelajaranService.FindMataPelajaranByName(input.MataPelajaran)

	if err_ != nil {
		response := helper.APIResponse("Create mata pelajaran failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if mataPelajaran.IdMataPelajaran != 0 {
		response := helper.APIResponse("Mata Pelajaran Already Exist", http.StatusConflict, "error", nil)
		c.JSON(http.StatusConflict, response)
		return
	}

	newMataPelajaran, err := h.mata_pelajaranService.CreateNewMataPelajaran(input)

	if err_ != nil {
		response := helper.APIResponse("Create mata pelajaran failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := mata_pelajaran.FormatMataPelajaranDetail(newMataPelajaran)
	response := helper.APIResponse("Success create new mata pelajaran", http.StatusCreated, "success", formatter)
	c.JSON(http.StatusCreated, response)
}
