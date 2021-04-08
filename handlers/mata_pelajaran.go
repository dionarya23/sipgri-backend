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

func (h *mata_pelajaranHandler) GetAll(c *gin.Context) {
	mataPelajaran, err := h.mata_pelajaranService.GetAllMataPelajaran()
	if err != nil {
		response := helper.APIResponse("Get mata pelajaran failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := mata_pelajaran.FormatListMataPelajaran(mataPelajaran)
	response := helper.APIResponse("List get mata pelajaran", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *mata_pelajaranHandler) GetOne(c *gin.Context) {
	var input mata_pelajaran.InputIDMataPelajaran

	err := c.ShouldBindUri(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Get mata pelajaran failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	mataPelajaran, err := h.mata_pelajaranService.GetOneMataPelajaran(input.IdMataPelajaran)
	if err != nil {
		response := helper.APIResponse("Get mata pelajaran failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formatter := mata_pelajaran.FormatMataPelajaranDetail(mataPelajaran)
	response := helper.APIResponse("Success get mata pelajaran", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *mata_pelajaranHandler) UpdatedMataPelajaran(c *gin.Context) {
	var inputID mata_pelajaran.InputIDMataPelajaran
	var inpuData mata_pelajaran.InputNewMataPelajaran
	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.APIResponse("Failed to update guru", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	err_ := c.ShouldBindJSON(&inpuData)
	if err_ != nil {
		errors := helper.FormatValidationError(err_)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Update Guru failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	updateMataPelajaran, err := h.mata_pelajaranService.UpdateMataPelajaranByID(inputID, inpuData)
	if err != nil {
		response := helper.APIResponse("Failed to update guru", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to update guru", http.StatusOK, "success", mata_pelajaran.FormatMataPelajaranDetail(updateMataPelajaran))
	c.JSON(http.StatusOK, response)
}

func (h *mata_pelajaranHandler) DeleteByIDMataPelajaran(c *gin.Context) {
	var input mata_pelajaran.InputIDMataPelajaran

	err := c.ShouldBindUri(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Get mata pelajaran failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	isMPExist, err := h.mata_pelajaranService.GetOneMataPelajaran(input.IdMataPelajaran)
	if err != nil {
		response := helper.APIResponse("Get mata pelajaran failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if isMPExist.IdMataPelajaran == 0 {
		response := helper.APIResponse("Mata pelajaran Not Found", http.StatusNotFound, "error", nil)
		c.JSON(http.StatusNotFound, response)
		return
	}

	err_ := h.mata_pelajaranService.DeleteMataPelajaranById(input.IdMataPelajaran)
	if err_ != nil {
		response := helper.APIResponse("Get mata pelajaran failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success delete mata pelajaran", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}
