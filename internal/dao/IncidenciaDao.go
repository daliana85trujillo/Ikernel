package dao

import (
	"Ikernel/internal/model/entity"
	"context"

	"gorm.io/gorm"
)

type IncidenciaDao struct {
	db *gorm.DB
}

func NewIncidenciaDao(db *gorm.DB) *IncidenciaDao {
	return &IncidenciaDao{db: db}
}

// Listar todas
func (i *IncidenciaDao) FindAll(ctx context.Context) ([]entity.Incidencia, error) {
	var incidencias []entity.Incidencia
	if err := i.db.WithContext(ctx).Find(&incidencias).Error; err != nil {
		return nil, err
	}
	return incidencias, nil
}

// Buscar por ID
func (i *IncidenciaDao) FindById(ctx context.Context, id int) (*entity.Incidencia, error) {
	var incidencia entity.Incidencia
	if err := i.db.WithContext(ctx).First(&incidencia, id).Error; err != nil {
		return nil, err
	}
	return &incidencia, nil
}

// Crear
func (i *IncidenciaDao) Create(ctx context.Context, incidencia *entity.Incidencia) error {
	return i.db.WithContext(ctx).Create(incidencia).Error
}

// Actualizar
func (i *IncidenciaDao) Update(ctx context.Context, incidencia *entity.Incidencia) error {
	return i.db.WithContext(ctx).Save(incidencia).Error
}

// Eliminar
func (i *IncidenciaDao) Delete(ctx context.Context, id int) error {
	return i.db.WithContext(ctx).Delete(&entity.Incidencia{}, id).Error
}
