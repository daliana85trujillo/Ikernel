package service

import (
	"Ikernel/internal/dao"
	"Ikernel/internal/model/dto"
	"Ikernel/internal/model/entity"
	"context"

	"gorm.io/gorm"
)

type RolService struct {
	db  *gorm.DB
	dao *dao.RolesDao
}

func NewRolService(db *gorm.DB, dao *dao.RolesDao) *RolService {
	return &RolService{db: db, dao: dao}
}

// Crear un rol
func (s *RolService) Create(ctx context.Context, rolDto *dto.RolesDto) error {
	return s.dao.Create(ctx, rolDto)
}

// Obtener todos los roles
func (s *RolService) FindAll(ctx context.Context) ([]dto.RolesDto, error) {
	return s.dao.FindAll(ctx)
}

// Buscar rol por ID
func (s *RolService) FindById(ctx context.Context, id int) (*dto.RolesDto, error) {
	return s.dao.FindById(ctx, id)
}

// Actualizar rol
func (s *RolService) Update(ctx context.Context, rolDto *dto.RolesDto) error {
	return s.dao.Update(ctx, rolDto)
}

// Eliminar rol
func (s *RolService) Delete(ctx context.Context, id int) error {
	return s.dao.Delete(ctx, id)
}

// Asignar un rol a un usuario
func (s *RolService) AsignarRol(usuarioId int, rolId int) error {
	asignacion := map[string]any{
		"usuario_id": usuarioId,
		"rol_id":     rolId,
	}
	return s.db.Table("usuario_rols").FirstOrCreate(asignacion, asignacion).Error
}

// Quitar un rol a un usuario
func (s *RolService) QuitarRol(usuarioId int, rolId int) error {
	return s.db.Table("usuario_rols").
		Where("usuario_id = ? AND rol_id = ?", usuarioId, rolId).
		Delete(nil).Error
}

// Consultar roles de un usuario
func (s *RolService) GetRolesDeUsuario(usuarioId int) ([]entity.Rol, error) {
	var roles []entity.Rol
	err := s.db.Table("rols").
		Select("rols.*").
		Joins("JOIN usuario_rols ur ON ur.rol_id = rols.id_rol").
		Where("ur.usuario_id = ?", usuarioId).
		Scan(&roles).Error
	return roles, err
}
