package service

import (
	"Ikernel/internal/dao"
	"Ikernel/internal/model/dto"
	"context"

	"github.com/gin-gonic/gin"
)

type ActividadService struct {
	dao *dao.ActividadDao
}

func (s *ActividadService) GetAll(ctx *gin.Context) (any, any) {
	panic("unimplemented")
}

func NewActivityService(dao *dao.ActividadDao) *ActividadService {
	return &ActividadService{dao: dao}
}

func (s *ActividadService) FindAll(ctx context.Context) ([]dto.ActividadDto, error) {
	activities, err := s.dao.FindAll() // ✅ sin ctx
	if err != nil {
		return nil, err
	}
	var dtos []dto.ActividadDto
	for _, a := range activities {
		dtos = append(dtos, dto.ActividadDto{
			Id:          a.Id,
			Description: a.Description,
			Status:      a.Status,
			// Etapaid:     a.EtapaId,
		})
	}
	return dtos, nil
}

func (s *ActividadService) GetById(ctx context.Context, id int) (*dto.ActividadDto, error) {
	activity, err := s.dao.FindById(id)
	if err != nil {
		return nil, err
	}
	return &dto.ActividadDto{
		Id:          activity.Id,
		Description: activity.Description,
		Status:      activity.Status,
		// Etapaid:     activity.EtapaId,
	}, nil
}

func (s *ActividadService) Create(ctx context.Context, dtoActivity *dto.ActividadDto) error {
	actividad := dtoActivity.ToEntity()
	return s.dao.Create(actividad)
}

func (s *ActividadService) Update(ctx context.Context, dtoActivity *dto.ActividadDto) error {
	actividad := dtoActivity.ToEntity()
	return s.dao.Update(actividad)
}

func (s *ActividadService) Delete(ctx context.Context, id int) error {
	return s.dao.Delete(id)
}
