package controller

import (
	"Ikernel/internal/model/dto"
	"Ikernel/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ActividadesController struct {
	actividadService *service.ActividadService
}

// Constructor
func NewActividadesController(actividadService *service.ActividadService) *ActividadesController {
	return &ActividadesController{actividadService: actividadService}
}

// -------------------------
// CRUD ACTIVIDADES
// -------------------------
func (c *ActividadesController) GetAll(ctx *gin.Context) {
	actividades, err := c.actividadService.FindAll(ctx)
	if err != nil {
		HandleActividadError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": actividades})
}

func (c *ActividadesController) GetById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	actividad, err := c.actividadService.GetById(ctx, id)
	if err != nil {
		HandleActividadError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": actividad})
}

func HandleActividadError(ctx *gin.Context, err error) {
	panic("unimplemented")
}

func (c *ActividadesController) Create(ctx *gin.Context) {
	var actividad dto.ActividadDto
	if err := ctx.ShouldBindJSON(&actividad); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.actividadService.Create(ctx, &actividad); err != nil {
		HandleActividadError(ctx, err)
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"data": actividad})
}

func (c *ActividadesController) Update(ctx *gin.Context) {
	var actividad dto.ActividadDto
	if err := ctx.ShouldBindJSON(&actividad); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.actividadService.Update(ctx, &actividad); err != nil {
		HandleActividadError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": actividad})
}

func (c *ActividadesController) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	if err := c.actividadService.Delete(ctx, id); err != nil {
		HandleActividadError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Actividad eliminada correctamente"})
}
