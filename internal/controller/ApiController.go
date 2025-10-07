package controller

import (
	"Ikernel/internal/model/entity"

	"gorm.io/gorm"
)

type ApiController struct {
	db *gorm.DB
}

func NewApiController(db *gorm.DB) *ApiController {
	return &ApiController{db: db}
}

// ----- Usuarios -----
func (c *ApiController) ListarUsuarios() []entity.Usuario {
	var usuarios []entity.Usuario
	c.db.Find(&usuarios)
	return usuarios
}

func (c *ApiController) CrearUsuario(usuario entity.Usuario) error {
	return c.db.Create(&usuario).Error
}

// ----- Proyectos -----
func (c *ApiController) ListarProyectos() []entity.Proyecto {
	var proyectos []entity.Proyecto
	c.db.Find(&proyectos)
	return proyectos
}

func (c *ApiController) CrearProyecto(proyecto entity.Proyecto) error {
	return c.db.Create(&proyecto).Error
}

// ----- Actividades -----
func (c *ApiController) ListarActividades() []entity.Actividad {
	var actividades []entity.Actividad
	c.db.Find(&actividades)
	return actividades
}

// ----- Etapas -----
func (c *ApiController) ListarEtapas() []entity.Etapa {
	var etapas []entity.Etapa
	c.db.Find(&etapas)
	return etapas
}

// ----- Incidencias -----
func (c *ApiController) ListarIncidencias() []entity.Incidencia {
	var incidencias []entity.Incidencia
	c.db.Find(&incidencias)
	return incidencias
}

// ----- Roles -----
func (c *ApiController) ListarRoles() []entity.Rol {
	var roles []entity.Rol
	c.db.Find(&roles)
	return roles
}
