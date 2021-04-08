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
		return
	}

	newGuru, err := h.guruService.RegisterGuru(input)
	if err != nil {
		response := helper.APIResponse("Register Guru failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := guru.FormatAuthGuru(newGuru, "")
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
		return
	}

	guru_, err := h.guruService.Login(input)
	if err != nil {
		response := helper.APIResponse("Login failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	token, err := h.authService.GenerateToken(guru_.Nip)
	if err != nil {
		response := helper.APIResponse("Login failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := guru.FormatAuthGuru(guru_, token)
	response := helper.APIResponse("Login success", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *guruHandler) GetAllGuru(c *gin.Context) {
	listGuru, err := h.guruService.GetAllGuru()

	if err != nil {
		response := helper.APIResponse("Error to get list guru", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("List of guru", http.StatusOK, "success", guru.FormatListGuru(listGuru))
	c.JSON(http.StatusOK, response)
}

func (h *guruHandler) GetOneGuru(c *gin.Context) {
	var input guru.GetGuruInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to get detail of guru", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	guruDetail, err := h.guruService.GetGuruByNip(input.Nip)
	if err != nil {
		response := helper.APIResponse("Failed to get detail of guru", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Guru detail", http.StatusOK, "success", guru.FormatDetailGuru(guruDetail))
	c.JSON(http.StatusOK, response)
}

func (h *guruHandler) UpdateGuru(c *gin.Context) {
	var inputNip guru.GetGuruInput
	var inputData guru.UpdateGuruInput

	err := c.ShouldBindUri(&inputNip)
	if err != nil {
		response := helper.APIResponse("Failed to update guru", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	err_ := c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err_)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Update Guru failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	updatedGuru, err := h.guruService.UpdateGuru(inputNip, inputData)
	if err != nil {
		response := helper.APIResponse("Failed to update guru", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to update guru", http.StatusOK, "success", guru.FormatDetailGuru(updatedGuru))
	c.JSON(http.StatusOK, response)
}

func (h *guruHandler) DeleteGuru(c *gin.Context) {
	var input guru.GetGuruInput
	err := c.ShouldBindUri(&input)

	if err != nil {
		response := helper.APIResponse("Failed to delete guru", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	err_ := h.guruService.DeleteGuruByNip(input.Nip)
	if err_ != nil {
		response := helper.APIResponse("Failed to delete guru", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success delete guru", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}
