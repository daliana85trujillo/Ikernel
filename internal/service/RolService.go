package service

import (
	"Ikernel/internal/model/dto"
	"Ikernel/internal/model/entity"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RolService struct {
	db *gorm.DB
}

func NewRolService(db *gorm.DB) *RolService {
	return &RolService{db: db}
}

func (s *RolService) Delete(ctx *gin.Context, id int) any {
	panic("unimplemented")
}

func (s *RolService) Update(ctx *gin.Context, dto *dto.RolesDto) any {
	panic("unimplemented")
}

func (s *RolService) Create(ctx *gin.Context, dto *dto.RolesDto) any {
	panic("unimplemented")
}

func (s *RolService) FindById(ctx *gin.Context, id int) (any, error) {
	panic("unimplemented")
}

func (s *RolService) GetAllRoles() (any, any) {
	panic("unimplemented")
}

func (s *RolService) FindAll(ctx *gin.Context) (any, any) {
	panic("unimplemented")
}

// Asignar un rol a un usuario
func (s *RolService) AsignarRol(usuarioId int, rolId int) error {
	usuarioRol := entity.Usuario{
		Id_Usuario: usuarioId,
		// Replace 'RolId' with the correct field name as defined in entity.Usuario, for example:
		// Roles: []int{rolId},
	}

	// Evita duplicados
	return s.db.FirstOrCreate(&usuarioRol, usuarioRol).Error
}

// Quitar un rol a un usuario
func (s *RolService) QuitarRol(usuarioId int, rolId int) error {
	return s.db.Delete(&entity.Usuario{}, "usuario_id = ? AND rol_id = ?", usuarioId, rolId).Error
}

// Consultar roles de un usuario
func (s *RolService) GetRolesDeUsuario(usuarioId int) ([]entity.Rol, error) {
	var roles []entity.Rol
	err := s.db.Table("rols").
		Select("rols.*").
		Joins("JOIN usuario_rols ur ON ur.rol_id = rols.id").
		Where("ur.usuario_id = ?", usuarioId).
		Scan(&roles).Error
	return roles, err
}
