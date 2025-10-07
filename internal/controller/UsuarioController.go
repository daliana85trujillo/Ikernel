package controller

import (
	"Ikernel/internal/model/dto"
	"Ikernel/internal/model/entity"
	"Ikernel/internal/service"
	"strconv"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type UsuarioController struct {
	UsuarioService *service.UsuarioService
	RolService     *service.RolService
}

func NewUsuarioController(usuarioService *service.UsuarioService, rolService *service.RolService) *UsuarioController {
	return &UsuarioController{
		UsuarioService: usuarioService,
		RolService:     rolService,
	}
}

// -------------------------
// CRUD USUARIOS
// -------------------------
func (c *UsuarioController) GetAll(r *ghttp.Request) {
	usuarios, err := c.UsuarioService.FindAll(r.GetCtx())
	if err != nil {
		r.Response.WriteStatusExit(500, err.Error())
		return
	}
	r.Response.WriteJson(usuarios)
}

func (c *UsuarioController) GetById(r *ghttp.Request) {
	idStr := r.Get("id").String()
	id, err := strconv.Atoi(idStr)
	if err != nil {
		r.Response.WriteStatusExit(400, "ID inválido")
		return
	}
	usuario, err := c.UsuarioService.FindById(r.GetCtx(), id)
	if err != nil {
		r.Response.WriteStatusExit(500, err.Error())
		return
	}
	r.Response.WriteJson(usuario)
}

func (c *UsuarioController) Create(r *ghttp.Request) {
	var usuario dto.UsuarioDto

	// Parsear el cuerpo del request
	if err := r.Parse(&usuario); err != nil {
		r.Response.WriteJsonExit(g.Map{
			"status":  "error",
			"message": "Datos inválidos",
			"error":   err.Error(),
		})
		return
	}

	// Crear el usuario usando el servicio
	if err := c.UsuarioService.Create(r.GetCtx(), &entity.Usuario{}); err != nil {
		r.Response.WriteJsonExit(g.Map{
			"status":  "error",
			"message": "No se pudo crear el usuario",
			"error":   err.Error(),
		})
		return
	}

	// Respuesta exitosa
	r.Response.WriteJsonExit(g.Map{
		"status":  "success",
		"message": "Usuario creado correctamente",
		"data":    usuario,
	})
}

func (c *UsuarioController) Update(r *ghttp.Request) {
	var usuario dto.UsuarioDto
	if err := r.Parse(&usuario); err != nil {
		r.Response.WriteStatusExit(400, err.Error())
		return
	}
	if err := c.UsuarioService.Update(r.GetCtx(), &usuario); err != nil {
		r.Response.WriteStatusExit(500, err.Error())
		return
	}
	r.Response.WriteJson(usuario)
}

// -------------------------
// ASIGNAR / QUITAR ROL
// -------------------------
func (c *UsuarioController) AsignarRol(r *ghttp.Request) {
	usuarioId, err := strconv.Atoi(r.Get("id").String())
	if err != nil {
		r.Response.WriteStatusExit(400, "ID de usuario inválido")
		return
	}
	rolId, err := strconv.Atoi(r.Get("rolId").String())
	if err != nil {
		r.Response.WriteStatusExit(400, "ID de rol inválido")
		return
	}
	if err := c.RolService.AsignarRol(usuarioId, rolId); err != nil {
		r.Response.WriteStatusExit(500, err.Error())
		return
	}
	r.Response.WriteJson(map[string]string{"message": "Rol asignado correctamente"})
}
