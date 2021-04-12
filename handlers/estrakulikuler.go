package handlers

import (
	"fmt"
	"net/http"

	"github.com/dionarya23/sipgri-backend/estrakulikuler"
	"github.com/dionarya23/sipgri-backend/helper"
	"github.com/gin-gonic/gin"
)

type eskulHandler struct {
	eskulService estrakulikuler.Service
}

func NewEskulHandler(eskulService estrakulikuler.Service) *eskulHandler {
	return &eskulHandler{eskulService}
}

func (h *eskulHandler) Create(c *gin.Context) {
	var inputData estrakulikuler.InputNewEskull

	err := c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Create estrakulikuler failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newEskul, err := h.eskulService.Create(inputData)
	if err != nil {
		response := helper.APIResponse("Create estrakulikuler failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := estrakulikuler.FormatEskulDetail(newEskul)
	response := helper.APIResponse("Success create estrakulikuler", http.StatusCreated, "success", formatter)
	c.JSON(http.StatusCreated, response)
}

func (h *eskulHandler) GetAll(c *gin.Context) {
	listEskull, err := h.eskulService.GetAll()

	if err != nil {
		response := helper.APIResponse("Get list estrakulikuler failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := estrakulikuler.FormatEskulList(listEskull)
	response := helper.APIResponse("List estrakulikuler", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *eskulHandler) GetById(c *gin.Context) {
	var input estrakulikuler.InputIDEskul
	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Get list estrakulikuler failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	eskul, err := h.eskulService.GetByID(input)
	if err != nil {
		response := helper.APIResponse("Get list estrakulikuler failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if eskul.IDEstrakulikuler == 0 {
		response := helper.APIResponse("estrakulikuler not found", http.StatusNotFound, "error", nil)
		c.JSON(http.StatusNotFound, response)
		return
	}

	formatter := estrakulikuler.FormatEskulDetail(eskul)
	response := helper.APIResponse("estrakulikuler get kelas", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *eskulHandler) GetByNipGuru(c *gin.Context) {
	var input estrakulikuler.InputNipGuru
	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Get list estrakulikuler failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	eskul, err := h.eskulService.GetByNipGuru(input)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Error disini")
		response := helper.APIResponse("Get list estrakulikuler failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if eskul.IDEstrakulikuler == 0 {
		response := helper.APIResponse("estrakulikuler not found", http.StatusNotFound, "error", nil)
		c.JSON(http.StatusNotFound, response)
		return
	}

	formatter := estrakulikuler.FormatEskulDetail(eskul)
	response := helper.APIResponse("success get estrakulikuler", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *eskulHandler) UpdateById(c *gin.Context) {
	var inputID estrakulikuler.InputIDEskul
	var inputData estrakulikuler.InputNewEskull

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Update estrakulikuler failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err_ := c.ShouldBindJSON(&inputData)
	if err_ != nil {
		errors := helper.FormatValidationError(err_)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Update estrakulikuler failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	updatedEskul, err := h.eskulService.Update(inputID, inputData)
	if err != nil {
		response := helper.APIResponse("Update estrakulikuler failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := estrakulikuler.FormatEskulDetail(updatedEskul)
	response := helper.APIResponse("Success update estrakulikuler", http.StatusOK, "error", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *eskulHandler) DeleteById(c *gin.Context) {
	var inputID estrakulikuler.InputIDEskul

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Delete estrakulikuler failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err_ := h.eskulService.Delete(inputID)

	if err_ != nil {
		response := helper.APIResponse("Delete estrakulikuler failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Delete estrakulikuler success", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)

}
