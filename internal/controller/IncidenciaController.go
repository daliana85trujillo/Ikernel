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
	incidencias, err := c.incidenciaService.FindAll(ctx)
	if err != nil {
		HandleIncidenciaError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": incidencias})
}

func HandleIncidenciaError(ctx *gin.Context, err error) {
	panic("unimplemented")
}

func (c *IncidenciasController) GetById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	incidencia, err := c.incidenciaService.FindById(ctx, id)
	if err != nil {
		HandleIncidenciaError(ctx, err)
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
	if err := c.incidenciaService.Create(ctx, &incidencia); err != nil {
		HandleIncidenciaError(ctx, err)
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"data": incidencia})
}

func (c *IncidenciasController) Update(ctx *gin.Context) {
	var incidencia dto.IncidenciaDto
	if err := ctx.ShouldBindJSON(&incidencia); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.incidenciaService.Update(ctx, &incidencia); err != nil {
		HandleIncidenciaError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": incidencia})
}

func (c *IncidenciasController) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	if err := c.incidenciaService.Delete(ctx, id); err != nil {
		HandleIncidenciaError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Incidencia eliminada correctamente"})
}
