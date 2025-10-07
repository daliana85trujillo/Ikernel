package controller

import (
	"Ikernel/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type AuthController struct {
	service *service.UsuarioService
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (c *AuthController) Login(r *ghttp.Request) {
	var req *LoginRequest
	if err := r.Parse(&req); err != nil {
		r.Response.WriteStatus(400, "Datos inválidos")
		return
	}

	usuario, err := c.service.ValidarCredenciales(r.Context(), req.Email, req.Password)
	if err != nil {
		r.Response.WriteStatus(401, "Credenciales incorrectas")
		return
	}
	r.Response.WriteJson(g.Map{
		"message": "Login exitoso",
		"user":    usuario,
	})
}
