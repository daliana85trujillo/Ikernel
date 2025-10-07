package router

import (
	"Ikernel/internal/controller"

	"github.com/gogf/gf/v2/net/ghttp"
)

// RegisterRoutes registra todas las rutas de la API
func RegisterRoutes(
	s *ghttp.Server,
	userCtrl *controller.UsuarioController,
	rolCtrl *controller.RolesController,
	actividadCtrl *controller.ActividadesController,
	etapaCtrl *controller.EtapasController,
	incidenciaCtrl *controller.IncidenciasController,
	proyectoCtrl *controller.ProyectoController,
) {

	// Rutas de Usuario
	s.Group("/usuario", func(group *ghttp.RouterGroup) {
		group.POST("/create", userCtrl.Create)
		group.GET("/findAll", userCtrl.GetAll)
		group.GET("/{id}", userCtrl.GetById)
		group.PUT("/update", userCtrl.Update)
	})

	// Rutas de Rol
	s.Group("/rol", func(group *ghttp.RouterGroup) {
		group.POST("/create", rolCtrl.Create)
		group.GET("/findAll", rolCtrl.GetAll)
		group.GET("/{id}", rolCtrl.GetById)
		group.PUT("/update", rolCtrl.Update)
	})

	// Rutas de Actividad
	s.Group("/actividad", func(group *ghttp.RouterGroup) {
		group.POST("/create", actividadCtrl.Create)
		group.GET("/findAll", actividadCtrl.GetAll)
		group.GET("/{id}", actividadCtrl.GetById)
		group.PUT("/update", actividadCtrl.Update)
	})

	// Rutas de Etapa
	s.Group("/etapa", func(group *ghttp.RouterGroup) {
		group.POST("/create", etapaCtrl.Create)
		group.GET("/findAll", etapaCtrl.GetAll)
		group.GET("/{id}", etapaCtrl.GetById)
		group.PUT("/update", etapaCtrl.Update)
	})

	// Rutas de Incidencia
	s.Group("/incidencia", func(group *ghttp.RouterGroup) {
		group.POST("/create", incidenciaCtrl.Create)
		group.GET("/findAll", incidenciaCtrl.GetAll)
		group.GET("/{id}", incidenciaCtrl.GetById)
		group.PUT("/update", incidenciaCtrl.Update)
	})

	// Rutas de Proyecto
	s.Group("/proyecto", func(group *ghttp.RouterGroup) {
		group.POST("/create", proyectoCtrl.Create)
		group.GET("/findAll", proyectoCtrl.GetAll)
		group.GET("/{id}", proyectoCtrl.GetById)
		group.PUT("/update", proyectoCtrl.Update)
	})
}
