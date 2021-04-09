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

func (h *pesertaDidikHandler) GetAllPesertaDidik(c *gin.Context) {
	pesertaDidik, err := h.pesertaDidikService.GetAllPesertaDidik()

	if err != nil {
		response := helper.APIResponse("Get All peserta didik failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := peserta_didik.FormatPesertaDidikList(pesertaDidik)
	response := helper.APIResponse("List Peserta Didik", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *pesertaDidikHandler) GetOnePesertaDidik(c *gin.Context) {
	queryParams := c.Request.URL.Query()

	if len(queryParams) != 1 {
		response := helper.APIResponse("Get peserta didik failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	pesertaDidik, err := h.pesertaDidikService.GetOnePesertaDidik(queryParams)

	if err != nil {
		response := helper.APIResponse("Get peserta didik failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if pesertaDidik.Nisn == "" {
		response := helper.APIResponse("peserta didik not found", http.StatusNotFound, "error", nil)
		c.JSON(http.StatusNotFound, response)
		return
	}

	formatter := peserta_didik.FormatPesertaDidikDetail(pesertaDidik)
	response := helper.APIResponse("Success get peserta didik", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *pesertaDidikHandler) UpdatePesertaDidik(c *gin.Context) {
	var inputNisn peserta_didik.InputNisn
	var inputData peserta_didik.InputDataPesertaDidik

	err := c.ShouldBindUri(&inputNisn)
	if err != nil {
		response := helper.APIResponse("Update peserta didik failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	err_ := c.ShouldBindJSON(&inputData)
	if err_ != nil {
		response := helper.APIResponse("Update peserta didik failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	updatedPesertaDidik, err := h.pesertaDidikService.UpdatePesertaDidikByNisn(inputNisn, inputData)
	formatter := peserta_didik.FormatPesertaDidikDetail(updatedPesertaDidik)
	response := helper.APIResponse("Success get peserta didik", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *pesertaDidikHandler) DeleteByNisn(c *gin.Context) {
	var inputNisn peserta_didik.InputNisn

	err := c.ShouldBindUri(&inputNisn)
	if err != nil {
		response := helper.APIResponse("Delete peserta didik failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	params := make(map[string][]string)
	params["nisn"] = append(params["nisn"], inputNisn.Nisn)

	pesertaDidik, err := h.pesertaDidikService.GetOnePesertaDidik(params)
	if err != nil {
		response := helper.APIResponse("Delete peserta didik failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if pesertaDidik.Nisn == "" {
		response := helper.APIResponse("peserta didik not found", http.StatusNotFound, "error", nil)
		c.JSON(http.StatusNotFound, response)
		return
	}

	err_ := h.pesertaDidikService.DeleteByNisn(inputNisn)
	if err_ != nil {
		response := helper.APIResponse("Delete peserta didik failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success delete peserta didik", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}
