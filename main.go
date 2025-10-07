package main

import (
	router "Ikernel/internal/Router"
	"Ikernel/internal/cmd"
	"Ikernel/internal/controller"
	"Ikernel/internal/dao"
	"Ikernel/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

func main() {
	// Migraciones y conexión a la base de datos
	db := cmd.Migrate()

	// Crear instancia del servidor
	s := g.Server()
	s.SetPort(8199)

	// Inicialización de todos los servicios con los DAOs
	usuarioService := service.NewUsuarioService(dao.NewUsuarioDao(db))
	rolService := service.NewRolService(db, dao.NewRolesDao(db))
	actividadService := service.NewActivityService(dao.NewActividadDao(db))
	etapaService := service.NewEtapaService(dao.NewEtapaDao(db))
	incidenciaService := service.NewIncidenciaService(dao.NewIncidenciaDao(db))
	proyectoService := service.NewProyectoService(dao.NewProyectoDao(db))

	// Inicialización de todos los controladores
	userCtrl := controller.NewUsuarioController(usuarioService, rolService)
	rolCtrl := controller.NewRolesController(rolService)
	actividadCtrl := controller.NewActividadesController(actividadService)
	etapaCtrl := controller.NewEtapasController(etapaService)
	incidenciaCtrl := controller.NewIncidenciasController(incidenciaService)

	// ⚠️ Asegúrate de que tu constructor de ProyectoController reciba estos 4 servicios
	proyectoCtrl := controller.NewProyectoController(
		proyectoService,
		etapaService,
		usuarioService,
		actividadService,
	)

	// Registro de todas las rutas
	router.RegisterRoutes(
		s,
		userCtrl,
		rolCtrl,
		actividadCtrl,
		etapaCtrl,
		incidenciaCtrl,
		proyectoCtrl,
	)

	// Ejecutar el servidor
	s.Run()
}
