package router

import (
	"Ikernel/internal/controller"

	"github.com/gogf/gf/v2/net/ghttp"
)

func RegisterRoutes(s *ghttp.Server, userCtrl *controller.UsuarioController) {
	s.Group("/usuario", func(group *ghttp.RouterGroup) {
		// Usuarios
		//group.POST("/login", userCtrl.Login)
		// group.POST("/loginAdmin", userCtrl.LoginAdmin)
		group.POST("/create", userCtrl.Create)
		group.GET("/findAll", userCtrl.GetById)
		group.GET("/usuario/{id}", userCtrl.GetAll)
		group.PUT("/update", userCtrl.Update)
		//group.POST("/filter", userCtrl.Filter)
	})

	s.Group("/Rol", func(group *ghttp.RouterGroup) {
		// Roles
		//group.POST("/login", userCtrl.Login)
		group.POST("/create", userCtrl.Create)
		group.GET("/GetAll", userCtrl.GetAll)
		group.GET("/Rol/{id}", userCtrl.GetById)
		group.PUT("/update", userCtrl.Update)
		//group.POST("/filter", rolCtrl.Filter)
	})

	// Agrega más grupos y rutas según sea necesario
	s.Group("/Actividad", func(group *ghttp.RouterGroup) {
		// Actividades
		group.POST("/create", userCtrl.Create)
		group.GET("/findAll", userCtrl.GetById)
		group.GET("/Actividad/{id}", userCtrl.GetById)
		group.PUT("/update", userCtrl.Update)
		//group.POST("/filter", userCtrl.Filter)
	})

	s.Group("/Etapa", func(group *ghttp.RouterGroup) {
		// Etapas
		group.POST("/create", userCtrl.Create)
		group.GET("/findAll", userCtrl.GetAll)
		group.GET("/Etapa/{id}", userCtrl.GetById)
		group.PUT("/update", userCtrl.Update)
		//group.POST("/filter", userCtrl.Filter)
	})

	s.Group("/Incidencia", func(group *ghttp.RouterGroup) {
		// Incidencias
		group.POST("/create", userCtrl.Create)
		group.GET("/findAll", userCtrl.GetAll)
		group.GET("/Incidencia/{id}", userCtrl.GetById)
		group.PUT("/update", userCtrl.Update)
		//group.POST("/filter", userCtrl.Filter)
	})

	s.Group("/Proyecto", func(group *ghttp.RouterGroup) {
		// Proyectos
		group.POST("/create", userCtrl.Create)
		group.GET("/findAll", userCtrl.GetAll)
		group.GET("/Proyecto/{id}", userCtrl.GetById)
		group.PUT("/update", userCtrl.Update)
		//group.POST("/filter", userCtrl.Filter)
	})

}
