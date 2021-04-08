package handlers

import (
	"net/http"

	"github.com/dionarya23/sipgri-backend/auth"
	"github.com/dionarya23/sipgri-backend/guru"
	"github.com/dionarya23/sipgri-backend/helper"
	"github.com/gin-gonic/gin"
)

type guruHandler struct {
	guruService guru.Service
	authService auth.Service
}

func NewGuruHandler(guruService guru.Service, authService auth.Service) *guruHandler {
	return &guruHandler{guruService, authService}
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

	formatter := guru.FormatGuru(newGuru, "")
	response := helper.APIResponse("Success create guru", http.StatusCreated, "success", formatter)
	c.JSON(http.StatusCreated, response)
}

func (h *guruHandler) Login(c *gin.Context) {
	var input guru.LoginGuruInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Login Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
	}

	guru_, err := h.guruService.Login(input)
	if err != nil {
		response := helper.APIResponse("Login failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
	}

	token, err := h.authService.GenerateToken(guru_.Nip)
	if err != nil {
		response := helper.APIResponse("Login failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := guru.FormatGuru(guru_, token)
	response := helper.APIResponse("Login success", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}
