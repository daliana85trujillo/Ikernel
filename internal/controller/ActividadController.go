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
	actividades, err := c.actividadService.FindAll(ctx.Request.Context())
	if err != nil {
		handleActividadError(ctx, err)
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

	actividad, err := c.actividadService.GetById(ctx.Request.Context(), id)
	if err != nil {
		handleActividadError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": actividad})
}

func (c *ActividadesController) Create(ctx *gin.Context) {
	var actividad dto.ActividadDto

	if err := ctx.ShouldBindJSON(&actividad); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.actividadService.Create(ctx.Request.Context(), &actividad); err != nil {
		handleActividadError(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Actividad creada correctamente",
		"data":    actividad,
	})
}

func (c *ActividadesController) Update(ctx *gin.Context) {
	var actividad dto.ActividadDto

	if err := ctx.ShouldBindJSON(&actividad); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.actividadService.Update(ctx.Request.Context(), &actividad); err != nil {
		handleActividadError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Actividad actualizada correctamente",
		"data":    actividad,
	})
}

func (c *ActividadesController) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	if err := c.actividadService.Delete(ctx.Request.Context(), id); err != nil {
		handleActividadError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Actividad eliminada correctamente"})
}

// -------------------------
// MANEJO DE ERRORES
// -------------------------
func handleActividadError(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusInternalServerError, gin.H{
		"error": err.Error(),
	})
}
