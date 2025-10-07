package controller

import (
	"Ikernel/internal/model/dto"
	"Ikernel/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RolesController struct {
	roleService *service.RolService
}

func NewRolesController(roleService *service.RolService) *RolesController {
	return &RolesController{roleService: roleService}
}

// Obtener todos los roles
func (c *RolesController) GetAll(ctx *gin.Context) {
	roles, err := c.roleService.FindAll(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Rol no obtenido"})
		return
	}
	ctx.JSON(http.StatusOK, roles)
}

// Obtener un rol por ID
func (c *RolesController) GetById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	role, err := c.roleService.FindById(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, role)
}

// Crear un nuevo rol
func (c *RolesController) Create(ctx *gin.Context) {
	var role dto.RolesDto
	if err := ctx.ShouldBindJSON(&role); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.roleService.Create(ctx, &role); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "No se puede crear un rol"})
		return
	}
	ctx.JSON(http.StatusCreated, role)
}

// Actualizar un rol existente
func (c *RolesController) Update(ctx *gin.Context) {
	var role dto.RolesDto
	if err := ctx.ShouldBindJSON(&role); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.roleService.Update(ctx, &role); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "El rol no se pudo actualizar"})
		return
	}
	ctx.JSON(http.StatusOK, role)
}

// Eliminar un rol
func (c *RolesController) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	if err := c.roleService.Delete(ctx, id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "El rol no se pudo eliminar"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Rol eliminado correctamente"})
}
