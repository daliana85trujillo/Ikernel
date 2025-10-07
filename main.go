package main

import (
	router "Ikernel/internal/Router"
	"Ikernel/internal/cmd"
	"Ikernel/internal/controller"
	"Ikernel/internal/dao"
	"Ikernel/internal/service"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func main() {

	// Iniciar el servidor HTTP
	db := cmd.Migrate()
	_ = db
	s := g.Server()
	s.SetPort(8199)

	// Aquí puedes registrar tus controladores y rutas
	s = ghttp.GetServer()

	UsuarioService := service.NewUsuarioService(cmd.UsuarioDao(db))
	userCtrl := controller.NewUsuarioController(UsuarioService)

	router.RegisterRoutes(s, userCtrl)

	// Iniciar el servidor

	s.Run()
}
