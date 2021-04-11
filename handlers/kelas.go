package handlers

import (
	"net/http"

	"github.com/dionarya23/sipgri-backend/helper"
	"github.com/dionarya23/sipgri-backend/kelas"
	"github.com/gin-gonic/gin"
)

type kelasHandler struct {
	kelasService kelas.Service
}

func NewKelasHandler(kelasService kelas.Service) *kelasHandler {
	return &kelasHandler{kelasService}
}

func (h *kelasHandler) CreateKelas(c *gin.Context) {
	var input kelas.InputNewKelas

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Create new kelas failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newKelas, err := h.kelasService.CreateKelas(input)
	if err != nil {
		response := helper.APIResponse("Create new kelas failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := kelas.FormatKelasDetail(newKelas)
	response := helper.APIResponse("Success create new kelas", http.StatusCreated, "success", formatter)
	c.JSON(http.StatusCreated, response)
}

func (h *kelasHandler) GetAll(c *gin.Context) {
	listKelas, err := h.kelasService.GetAll()

	if err != nil {
		response := helper.APIResponse("Get list kelas failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := kelas.FormatKelasList(listKelas)
	response := helper.APIResponse("List kelas", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *kelasHandler) GetById(c *gin.Context) {
	var input kelas.InputIDKelas
	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Get list kelas failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	kelas_, err := h.kelasService.GetById(input)
	if err != nil {
		response := helper.APIResponse("Get list kelas failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if kelas_.IDKelas == 0 {
		response := helper.APIResponse("Kelas not found", http.StatusNotFound, "error", nil)
		c.JSON(http.StatusNotFound, response)
		return
	}

	formatter := kelas.FormatKelasDetail(kelas_)
	response := helper.APIResponse("Success get kelas", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *kelasHandler) GetByNipWali(c *gin.Context) {
	var input kelas.InputNipWali

	err := c.ShouldBindUri(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Get peserta didik failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	kelas_, err := h.kelasService.GetByNipWali(input)
	if err != nil {
		response := helper.APIResponse("Get peserta didik failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if kelas_.IDKelas == 0 {
		response := helper.APIResponse("Nip Wali not found", http.StatusNotFound, "error", nil)
		c.JSON(http.StatusNotFound, response)
		return
	}

	formatter := kelas.FormatKelasDetail(kelas_)
	response := helper.APIResponse("Success get peserta didik", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *kelasHandler) UpdateById(c *gin.Context) {
	var inputID kelas.InputIDKelas
	var inputData kelas.InputNewKelas

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Update kelas failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err_ := c.ShouldBindJSON(&inputData)
	if err_ != nil {
		errors := helper.FormatValidationError(err_)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Update kelas failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	updatedKelas, err := h.kelasService.UpdateById(inputID, inputData)
	if err != nil {
		response := helper.APIResponse("Update kelas failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := kelas.FormatKelasDetail(updatedKelas)
	response := helper.APIResponse("Success update kelas", http.StatusOK, "error", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *kelasHandler) DeleteById(c *gin.Context) {
	var inputID kelas.InputIDKelas

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Delete kelas failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err_ := h.kelasService.DeleteById(inputID)

	if err_ != nil {
		response := helper.APIResponse("Delete kelas failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Delete kelas success", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)

}
