package handlers

import (
	"net/http"

	"github.com/dionarya23/sipgri-backend/helper"
	"github.com/dionarya23/sipgri-backend/mengajar"
	"github.com/gin-gonic/gin"
)

type mengajarHandler struct {
	mengajarService mengajar.Service
}

func NewMengajarHandler(mengajarService mengajar.Service) *mengajarHandler {
	return &mengajarHandler{mengajarService}
}

func (h *mengajarHandler) Create(c *gin.Context) {
	var inputData mengajar.InputNewMengajar

	err := c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Create mengajar failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newMengajar, err := h.mengajarService.Create(inputData)
	if err != nil {
		response := helper.APIResponse("Create mengajar failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := mengajar.FormatMengajar(newMengajar)
	response := helper.APIResponse("Success create mengajar", http.StatusCreated, "success", formatter)
	c.JSON(http.StatusCreated, response)
}

func (h *mengajarHandler) GetAll(c *gin.Context) {
	mengajarList, err := h.mengajarService.GetAll()
	if err != nil {
		response := helper.APIResponse("Get list mengajar failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := mengajar.FormatListMengajar(mengajarList)
	response := helper.APIResponse("Get list mengajar failed", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *mengajarHandler) GetByKodeMengajar(c *gin.Context) {
	var inputKodeMengajar mengajar.InputKodeMengajar

	err := c.ShouldBindUri(&inputKodeMengajar)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Get list mengajar failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	mengajar_, err := h.mengajarService.GetByKodeMengajar(inputKodeMengajar)
	if err != nil {
		response := helper.APIResponse("Get Mengajar", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := mengajar.FormatMengajar(mengajar_)
	response := helper.APIResponse("Get Mengajar", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *mengajarHandler) GetByNipGuru(c *gin.Context) {
	var inputNipGuru mengajar.InputNipGuru

	err := c.ShouldBindUri(&inputNipGuru)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Get mengajar failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	mengajar_, err := h.mengajarService.GetByNipGuru(inputNipGuru)
	if err != nil {
		response := helper.APIResponse("Get mengajar failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := mengajar.FormatListMengajar(mengajar_)
	response := helper.APIResponse("Get mengajar success", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *mengajarHandler) UpdateMengajar(c *gin.Context) {
	var inputID mengajar.InputKodeMengajar
	var inputData mengajar.InputNewMengajar

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Update mengajar failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err_ := c.ShouldBindJSON(&inputData)
	if err_ != nil {
		errors := helper.FormatValidationError(err_)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Update mengajar failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	mengajar_, err := h.mengajarService.GetByKodeMengajar(inputID)
	if err != nil {
		response := helper.APIResponse("Update mengajar failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if mengajar_.KodeMengajar == "" {
		response := helper.APIResponse("Mengajar not found", http.StatusNotFound, "error", nil)
		c.JSON(http.StatusNotFound, response)
		return
	}

	updatedMengajar, err := h.mengajarService.UpdateByKodeMengajar(inputID, inputData)
	if err != nil {
		response := helper.APIResponse("Update mengajar failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := mengajar.FormatMengajar(updatedMengajar)
	response := helper.APIResponse("Update mengajar success", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *mengajarHandler) DeleteByKodeMengajar(c *gin.Context) {
	var inputID mengajar.InputKodeMengajar

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Delete mengajar failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	mengajar_, err := h.mengajarService.GetByKodeMengajar(inputID)
	if err != nil {
		response := helper.APIResponse("Delete mengajar failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if mengajar_.KodeMengajar == "" {
		response := helper.APIResponse("Mengajar not found", http.StatusNotFound, "error", nil)
		c.JSON(http.StatusNotFound, response)
		return
	}

	err_ := h.mengajarService.DeleteByKodeMengajar(inputID)
	if err_ != nil {
		response := helper.APIResponse("Delete mengajar failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Delete mengajar success", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}
