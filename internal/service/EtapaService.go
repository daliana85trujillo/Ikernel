package service

import (
	"Ikernel/internal/dao"
	"Ikernel/internal/model/dto"
	"Ikernel/internal/model/entity"
	"context"
	"time"
)

type EtapaService struct {
	dao dao.Idao[entity.Etapa, int]
}

func NewEtapaService(dao dao.Idao[entity.Etapa, int]) *EtapaService {
	return &EtapaService{dao: dao}
}

// Listar todas las etapas
func (s *EtapaService) FindAll(ctx context.Context) ([]dto.EtapaDto, error) {
	etapas, err := s.dao.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var etapasDto []dto.EtapaDto
	for _, e := range etapas {
		etapasDto = append(etapasDto, dto.EtapaDto{
			Id:          e.Id,
			Name:        e.Name,
			Description: e.Description,
			Start_date:  e.Start_date.Format("2006-01-02"),
			End_date:    e.End_date.Format("2006-01-02"),
			State:       e.State,
		})
	}
	return etapasDto, nil
}

// Obtener etapa por ID
func (s *EtapaService) GetById(ctx context.Context, id int) (*dto.EtapaDto, error) {
	e, err := s.dao.FindById(ctx, id)
	if err != nil {
		return nil, err
	}
	return &dto.EtapaDto{
		Id:          e.Id,
		Name:        e.Name,
		Description: e.Description,
		Start_date:  e.Start_date.Format("2006-01-02"),
		End_date:    e.End_date.Format("2006-01-02"),
		State:       e.State,
	}, nil
}

// Crear etapa
func (s *EtapaService) Create(ctx context.Context, etapa *dto.EtapaDto) error {
	startDate, err := time.Parse("2006-01-02", etapa.Start_date)
	if err != nil {
		return err
	}
	endDate, err := time.Parse("2006-01-02", etapa.End_date)
	if err != nil {
		return err
	}

	entidad := entity.Etapa{
		Name:        etapa.Name,
		Description: etapa.Description,
		Start_date:  startDate,
		End_date:    endDate,
		State:       etapa.State,
	}
	return s.dao.Create(ctx, &entidad)
}

// Actualizar etapa
func (s *EtapaService) Update(ctx context.Context, etapa *dto.EtapaDto) error {
	startDate, err := time.Parse("2006-01-02", etapa.Start_date)
	if err != nil {
		return err
	}
	endDate, err := time.Parse("2006-01-02", etapa.End_date)
	if err != nil {
		return err
	}

	entidad := entity.Etapa{
		Id:          etapa.Id,
		Name:        etapa.Name,
		Description: etapa.Description,
		Start_date:  startDate,
		End_date:    endDate,
		State:       etapa.State,
	}
	return s.dao.Update(ctx, &entidad)
}

// Eliminar etapa
func (s *EtapaService) Delete(ctx context.Context, id int) error {
	return s.dao.Delete(ctx, &entity.Etapa{Id: id})
}
