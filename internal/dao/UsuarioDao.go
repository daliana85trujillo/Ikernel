package dao

import (
	"Ikernel/internal/model/dto"
	"Ikernel/internal/model/entity"
	"context"

	"gorm.io/gorm"
)

type UsuarioDao struct {
	db *gorm.DB
}

func NewUsuarioDao(db *gorm.DB) *UsuarioDao {
	return &UsuarioDao{db: db}
}

// Crear usuario
func (u *UsuarioDao) Create(ctx context.Context, usuario *entity.Usuario) error {
	return u.db.WithContext(ctx).Create(usuario).Error
}

// Actualizar usuario
func (u *UsuarioDao) Update(ctx context.Context, usuario *entity.Usuario) error {
	return u.db.WithContext(ctx).Save(usuario).Error
}

// Eliminar usuario
func (u *UsuarioDao) Delete(ctx context.Context, id int) error {
	return u.db.WithContext(ctx).Delete(&entity.Usuario{}, id).Error
}

// Buscar por ID
func (u *UsuarioDao) FindById(ctx context.Context, id int) (*entity.Usuario, error) {
	var usuario entity.Usuario
	if err := u.db.WithContext(ctx).First(&usuario, id).Error; err != nil {
		return nil, err
	}
	return &usuario, nil
}

// Listar todos
func (u *UsuarioDao) FindAll(ctx context.Context) ([]entity.Usuario, error) {
	var usuarios []entity.Usuario
	if err := u.db.WithContext(ctx).Find(&usuarios).Error; err != nil {
		return nil, err
	}
	return usuarios, nil
}
func (u *UsuarioDao) FindWithFilters(filter dto.UsuarioFilter) ([]entity.Usuario, error) {
	query := u.db.Model(&entity.Usuario{})

	if filter.Nombre != "" {
		query = query.Where("nombre LIKE ?", "%"+filter.Nombre+"%")
	}

	if filter.Email != "" {
		query = query.Where("email LIKE ?", "%"+filter.Email+"%")
	}

	if filter.Estado != "" {
		query = query.Where("estado = ?", filter.Estado)
	}

	var usuarios []entity.Usuario
	if err := query.Find(&usuarios).Error; err != nil {
		return nil, err
	}
	return usuarios, nil

}
func (d *UsuarioDao) GetDB() *gorm.DB {
	return d.db
}

func (d *UsuarioDao) FindByEmail(email string) (*entity.Usuario, error) {
	var usuario entity.Usuario
	if err := d.db.Where("email = ?", email).First(&usuario).Error; err != nil {
		return nil, err
	}
	return &usuario, nil
}
func (d *UsuarioDao) UpdatePassword(id int, hashedPassword string) error {
	return d.db.Model(&entity.Usuario{}).Where("id_usuario = ?", id).Update("password", hashedPassword).Error
}
