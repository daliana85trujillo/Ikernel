package dao

import (
	"Ikernel/internal/model/dto"
	"Ikernel/internal/model/entity"
	"context"

	"gorm.io/gorm"
)

type ProyectoDao struct {
	db *gorm.DB
}

func NewProyectoDao(db *gorm.DB) *ProyectoDao {
	return &ProyectoDao{db: db}
}

// Listar todos
func (p *ProyectoDao) FindAll(ctx context.Context) ([]entity.Proyecto, error) {
	var proyectos []entity.Proyecto
	if err := p.db.WithContext(ctx).Find(&proyectos).Error; err != nil {
		return nil, err
	}
	return proyectos, nil
}

// Crear
func (p *ProyectoDao) Create(ctx context.Context, proyecto *entity.Proyecto) error {
	return p.db.WithContext(ctx).Create(proyecto).Error
}

// Buscar por ID
func (p *ProyectoDao) FindById(ctx context.Context, id int) (*entity.Proyecto, error) {
	var proyecto entity.Proyecto
	if err := p.db.WithContext(ctx).First(&proyecto, id).Error; err != nil {
		return nil, err
	}
	return &proyecto, nil
}

// Actualizar
func (p *ProyectoDao) Update(ctx context.Context, proyecto *entity.Proyecto) error {
	return p.db.WithContext(ctx).Save(proyecto).Error
}

// Eliminar
func (p *ProyectoDao) Delete(ctx context.Context, id int) error {
	return p.db.WithContext(ctx).Delete(&entity.Proyecto{}, id).Error
}

func (p *ProyectoDao) FindByFiltro(ctx context.Context, filtro dto.FiltroProyectoDto) ([]entity.Proyecto, error) {
	query := p.db.WithContext(ctx).Model(&entity.Proyecto{})

	if filtro.Name != "" {
		query = query.Where("nombre LIKE ?", "%"+filtro.Name+"%")
	}
	if filtro.Status != "" {
		query = query.Where("estado = ?", filtro.Status)
	}
	if filtro.Start_date != "" {
		query = query.Where("fecha_inicio >= ?", filtro.Start_date)
	}
	if filtro.End_date != "" {
		query = query.Where("fecha_fin <= ?", filtro.End_date)
	}
	if filtro.LiderID != 0 {
		query = query.Where("lider_id = ?", filtro.LiderID)
	}

	var proyectos []entity.Proyecto
	if err := query.Find(&proyectos).Error; err != nil {
		return nil, err
	}
	return proyectos, nil
}
