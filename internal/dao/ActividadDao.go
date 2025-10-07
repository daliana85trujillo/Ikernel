package dao

import (
	"Ikernel/internal/model/entity"

	"gorm.io/gorm"
)

type ActividadDao struct {
	db *gorm.DB
}

func NewActividadDao(db *gorm.DB) *ActividadDao {
	return &ActividadDao{db: db}
}

func (a *ActividadDao) FindAll() ([]entity.Actividad, error) {
	var actividades []entity.Actividad
	result := a.db.Find(&actividades)
	return actividades, result.Error
}

func (a *ActividadDao) Create(actividad *entity.Actividad) error {
	return a.db.Create(actividad).Error
}

func (a *ActividadDao) FindById(id int) (*entity.Actividad, error) {
	var actividad entity.Actividad
	result := a.db.First(&actividad, id)
	return &actividad, result.Error
}

func (a *ActividadDao) Update(actividad *entity.Actividad) error {
	return a.db.Save(actividad).Error
}

func (a *ActividadDao) Delete(id int) error {
	return a.db.Delete(&entity.Actividad{}, id).Error
}
