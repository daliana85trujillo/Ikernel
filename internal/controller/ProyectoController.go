package controller

import (
	"Ikernel/internal/model/dto"
	"Ikernel/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProyectoController struct {
	proyectoService  *service.ProyectoService
	etapaService     *service.EtapaService
	usuarioService   *service.UsuarioService
	actividadService *service.ActividadService
}

// Constructor
func NewProyectoController(
	proyectoService *service.ProyectoService,
	etapaService *service.EtapaService,
	usuarioService *service.UsuarioService,
	actividadService *service.ActividadService,
) *ProyectoController {
	return &ProyectoController{
		proyectoService:  proyectoService,
		etapaService:     etapaService,
		usuarioService:   usuarioService,
		actividadService: actividadService,
	}
}

// -------------------------
// WRAPPER DE ERRORES SEGURO
// -------------------------
func handleProyectoError(ctx *gin.Context, err any) {
	if e, ok := err.(error); ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": e.Error()})
	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error desconocido"})
	}
}

// -------------------------
// CRUD PROYECTOS
// -------------------------
func (c *ProyectoController) GetAll(ctx *gin.Context) {
	proyectos, err := c.proyectoService.FindAll(ctx)
	if err != nil {
		handleProyectoError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": proyectos})
}

func (c *ProyectoController) GetById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	proyecto, err := c.proyectoService.FindById(ctx, id)
	if err != nil {
		handleProyectoError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": proyecto})
}

func (c *ProyectoController) Create(ctx *gin.Context) {
	var proyecto dto.ProyectoDto
	if err := ctx.ShouldBindJSON(&proyecto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.proyectoService.Create(ctx, &proyecto); err != nil {
		handleProyectoError(ctx, err)
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"data": proyecto})
}

func (c *ProyectoController) Update(ctx *gin.Context) {
	var proyecto dto.ProyectoDto
	if err := ctx.ShouldBindJSON(&proyecto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.proyectoService.Update(ctx, &proyecto); err != nil {
		handleProyectoError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": proyecto})
}

func (c *ProyectoController) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	if err := c.proyectoService.Delete(ctx, id); err != nil {
		handleProyectoError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Proyecto eliminado"})
}
