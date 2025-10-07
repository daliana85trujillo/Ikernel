package controller

import (
	"Ikernel/internal/model/dto"
	"Ikernel/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type IncidenciasController struct {
	incidenciaService *service.IncidenciaService
}

// Constructor
func NewIncidenciasController(incidenciaService *service.IncidenciaService) *IncidenciasController {
	return &IncidenciasController{incidenciaService: incidenciaService}
}

// -------------------------
// CRUD INCIDENCIAS
// -------------------------

func (c *IncidenciasController) GetAll(ctx *gin.Context) {
	incidencias, err := c.incidenciaService.FindAll(ctx.Request.Context())
	if err != nil {
		handleIncidenciaError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": incidencias})
}

func (c *IncidenciasController) GetById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	incidencia, err := c.incidenciaService.FindById(ctx.Request.Context(), id)
	if err != nil {
		handleIncidenciaError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": incidencia})
}

func (c *IncidenciasController) Create(ctx *gin.Context) {
	var incidencia dto.IncidenciaDto

	if err := ctx.ShouldBindJSON(&incidencia); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.incidenciaService.Create(ctx.Request.Context(), &incidencia); err != nil {
		handleIncidenciaError(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Incidencia creada correctamente", "data": incidencia})
}

func (c *IncidenciasController) Update(ctx *gin.Context) {
	var incidencia dto.IncidenciaDto

	if err := ctx.ShouldBindJSON(&incidencia); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.incidenciaService.Update(ctx.Request.Context(), &incidencia); err != nil {
		handleIncidenciaError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Incidencia actualizada correctamente", "data": incidencia})
}

func (c *IncidenciasController) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	if err := c.incidenciaService.Delete(ctx.Request.Context(), id); err != nil {
		handleIncidenciaError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Incidencia eliminada correctamente"})
}

// -------------------------
// MANEJO DE ERRORES
// -------------------------
func handleIncidenciaError(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusInternalServerError, gin.H{
		"error": err.Error(),
	})
}
