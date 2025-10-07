package dao

import (
	"Ikernel/internal/model/dto"
	"Ikernel/internal/model/entity"
	"context"

	"gorm.io/gorm"
)

type RolesDao struct {
	db *gorm.DB
}

func NewRolesDao(db *gorm.DB) *RolesDao {
	return &RolesDao{db: db}
}

// Listar todos
func (r *RolesDao) FindAll(ctx context.Context) ([]dto.RolesDto, error) {
	var roles []entity.Rol
	if err := r.db.WithContext(ctx).Find(&roles).Error; err != nil {
		return nil, err
	}

	// Convertir entity → dto aquí mismo
	var rolesDto []dto.RolesDto
	for _, rol := range roles {
		rolesDto = append(rolesDto, dto.RolesDto{
			Id_rol: rol.Id_rol,
			Name:   rol.Name_rol,
		})
	}
	return rolesDto, nil
}

// Crear
func (r *RolesDao) Create(ctx context.Context, rolDto *dto.RolesDto) error {
	rol := entity.Rol{
		Id_rol:   rolDto.Id_rol,
		Name_rol: rolDto.Name,
	}
	return r.db.WithContext(ctx).Create(&rol).Error
}

// Buscar por ID
func (r *RolesDao) FindById(ctx context.Context, id int) (*dto.RolesDto, error) {
	var rol entity.Rol
	if err := r.db.WithContext(ctx).First(&rol, id).Error; err != nil {
		return nil, err
	}

	rolDto := &dto.RolesDto{
		Id_rol: rol.Id_rol,
		Name:   rol.Name_rol,
	}
	return rolDto, nil
}

// Actualizar
func (r *RolesDao) Update(ctx context.Context, rolDto *dto.RolesDto) error {
	rol := entity.Rol{
		Id_rol:   rolDto.Id_rol,
		Name_rol: rolDto.Name,
	}
	return r.db.WithContext(ctx).Save(&rol).Error
}

// Eliminar
func (r *RolesDao) Delete(ctx context.Context, id int) error {
	return r.db.WithContext(ctx).Delete(&entity.Rol{}, id).Error
}
