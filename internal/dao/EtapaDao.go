package dao

import (
	"Ikernel/internal/model/entity"
	"context"

	"gorm.io/gorm"
)

type EtapaDao struct {
	db *gorm.DB
}

func NewEtapaDao(db *gorm.DB) *EtapaDao {
	return &EtapaDao{db: db}
}

// Obtener todas las etapas
func (e *EtapaDao) FindAll(ctx context.Context) ([]entity.Etapa, error) {
	var etapas []entity.Etapa
	if err := e.db.WithContext(ctx).Find(&etapas).Error; err != nil {
		return nil, err
	}
	return etapas, nil
}

// Crear nueva etapa
func (e *EtapaDao) Create(ctx context.Context, etapa *entity.Etapa) error {
	return e.db.WithContext(ctx).Create(etapa).Error
}

// Buscar por ID
func (e *EtapaDao) FindById(ctx context.Context, id int) (*entity.Etapa, error) {
	var etapa entity.Etapa
	if err := e.db.WithContext(ctx).First(&etapa, id).Error; err != nil {
		return nil, err
	}
	return &etapa, nil
}

// Actualizar una etapa
func (e *EtapaDao) Update(ctx context.Context, etapa *entity.Etapa) error {
	return e.db.WithContext(ctx).Save(etapa).Error
}

// Eliminar una etapa
func (e *EtapaDao) Delete(ctx context.Context, etapa *entity.Etapa) error {
	return e.db.WithContext(ctx).Delete(etapa).Error
}
