package handlers

import (
	"net/http"

	"github.com/dionarya23/sipgri-backend/helper"
	"github.com/dionarya23/sipgri-backend/jadwal"
	"github.com/gin-gonic/gin"
)

type jadwalHandler struct {
	jadwalService jadwal.Service
}

func NewJadwalHandler(jadwalService jadwal.Service) *jadwalHandler {
	return &jadwalHandler{jadwalService}
}

func (h *jadwalHandler) CreateNewData(c *gin.Context) {
	var inputData jadwal.InputJadwal

	err := c.ShouldBindJSON(&inputData)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Create jadwal failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newJadwal, err := h.jadwalService.Create(inputData)

	if err != nil {
		response := helper.APIResponse("Create jadwal failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := jadwal.FormatJadwalDetail(newJadwal)
	response := helper.APIResponse("Success create new jadwal", http.StatusCreated, "success", formatter)
	c.JSON(http.StatusCreated, response)
}

func (h *jadwalHandler) FindAllJadwal(c *gin.Context) {
	listJadwal, err := h.jadwalService.FindAll()

	if err != nil {
		response := helper.APIResponse("Create jadwal failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := jadwal.FormatJadwalList(listJadwal)
	response := helper.APIResponse("List jadwal", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *jadwalHandler) FindOneByIdJadwal(c *gin.Context) {
	var inputID jadwal.InputParamsIDJadwal

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Get jadwal failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	jadwal_, err := h.jadwalService.FindByIdJadwal(inputID)

	if err != nil {
		response := helper.APIResponse("Get jadwal failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if jadwal_.IDJadwal == 0 {
		response := helper.APIResponse("Jadwal not found", http.StatusNotFound, "error", nil)
		c.JSON(http.StatusNotFound, response)
		return
	}

	formatter := jadwal.FormatJadwalDetail(jadwal_)
	response := helper.APIResponse("Get jadwal", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *jadwalHandler) UpdateById(c *gin.Context) {
	var inputID jadwal.InputParamsIDJadwal
	var inputData jadwal.InputJadwal

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Update jadwal failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err_ := c.ShouldBindJSON(&inputData)

	if err_ != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Update jadwal failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	jadwal_, err := h.jadwalService.FindByIdJadwal(inputID)

	if err != nil {
		response := helper.APIResponse("Update jadwal failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if jadwal_.IDJadwal == 0 {
		response := helper.APIResponse("Jadwal not found", http.StatusNotFound, "error", nil)
		c.JSON(http.StatusNotFound, response)
		return
	}

	updatedJadwal, err := h.jadwalService.UpdateById(inputID, inputData)

	if err != nil {
		response := helper.APIResponse("Update jadwal failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := jadwal.FormatJadwalDetail(updatedJadwal)
	response := helper.APIResponse("Update jadwal success", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *jadwalHandler) DeleteById(c *gin.Context) {
	var inputID jadwal.InputParamsIDJadwal

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Delete jadwal failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	jadwal_, err := h.jadwalService.FindByIdJadwal(inputID)

	if err != nil {
		response := helper.APIResponse("Delete jadwal failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if jadwal_.IDJadwal == 0 {
		response := helper.APIResponse("Jadwal not found", http.StatusNotFound, "error", nil)
		c.JSON(http.StatusNotFound, response)
		return
	}

	err_ := h.jadwalService.Delete(inputID)
	if err_ != nil {
		response := helper.APIResponse("Delete jadwal failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success delete jadwal", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}
