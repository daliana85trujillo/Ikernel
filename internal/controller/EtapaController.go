package controller

import (
	"Ikernel/internal/model/dto"
	"Ikernel/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type EtapasController struct {
	etapaService *service.EtapaService
}

// Constructor
func NewEtapasController(etapaService *service.EtapaService) *EtapasController {
	return &EtapasController{etapaService: etapaService}
}

// -------------------------
// CRUD ETAPAS
// -------------------------
func (c *EtapasController) GetAll(ctx *gin.Context) {
	etapas, err := c.etapaService.FindAll(ctx)
	if err != nil {
		HandleEtapaError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": etapas})
}

func HandleEtapaError(ctx *gin.Context, err error) {
	panic("unimplemented")
}

func (c *EtapasController) GetById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	etapa, err := c.etapaService.GetById(ctx, id)
	if err != nil {
		HandleEtapaError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": etapa})
}

func (c *EtapasController) Create(ctx *gin.Context) {
	var etapa dto.EtapaDto
	if err := ctx.ShouldBindJSON(&etapa); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.etapaService.Create(ctx, &etapa); err != nil {
		HandleEtapaError(ctx, err)
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"data": etapa})
}

func (c *EtapasController) Update(ctx *gin.Context) {
	var etapa dto.EtapaDto
	if err := ctx.ShouldBindJSON(&etapa); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.etapaService.Update(ctx, &etapa); err != nil {
		HandleEtapaError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": etapa})
}

func (c *EtapasController) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	if err := c.etapaService.Delete(ctx, id); err != nil {
		HandleEtapaError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Etapa eliminada correctamente"})
}
