package service

import (
	"Ikernel/internal/dao"
	"Ikernel/internal/model/dto"
	"Ikernel/internal/model/entity"
	"context"
	"strconv"
	"time"
)

type IncidenciaService struct {
	dao *dao.IncidenciaDao
}

func NewIncidenciaService(dao *dao.IncidenciaDao) *IncidenciaService {
	return &IncidenciaService{dao: dao}
}

// Listar todas las incidencias
func (s *IncidenciaService) FindAll(ctx context.Context) ([]dto.IncidenciaDto, error) {
	incidencias, err := s.dao.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var incidenciasDto []dto.IncidenciaDto
	for _, i := range incidencias {
		incidenciasDto = append(incidenciasDto, dto.IncidenciaDto{
			Id:          i.Id,
			Type:        i.Type,
			Description: i.Description,
			Fase:        i.Fase,
			Duration:    strconv.Itoa(i.Duration),
			Date:        i.Date.Format("2006-01-02"),
			// ActividadId: i.EtapaId,
			// Etapa:       i.EtapaId,
		})
	}
	return incidenciasDto, nil
}

// Buscar incidencia por ID
func (s *IncidenciaService) FindById(ctx context.Context, id int) (*dto.IncidenciaDto, error) {
	i, err := s.dao.FindById(ctx, id)
	if err != nil {
		return nil, err
	}
	return &dto.IncidenciaDto{
		Id: i.Id,

		Type:        i.Type,
		Date:        i.Date.Format("2006-01-02"),
		Duration:    strconv.Itoa(i.Duration),
		Description: i.Description,
		Fase:        i.Fase,
		// ActividadId: i.ActividadId,
		// Etapa:       i.EtapaId,
	}, nil
}

func (s *IncidenciaService) Create(ctx context.Context, incidenciaDto *dto.IncidenciaDto) error {
	parsedDate, err := time.Parse("2006-01-02", incidenciaDto.Date)
	if err != nil {
		return err
	}
	parsedDuration, err := strconv.Atoi(incidenciaDto.Duration)
	if err != nil {
		return err
	}
	incidencia := entity.Incidencia{
		Type:        incidenciaDto.Type,
		Date:        parsedDate,
		Duration:    parsedDuration,
		Description: incidenciaDto.Description,
		Fase:        incidenciaDto.Fase,
		// ActividadId: incidenciaDto.ActividadId,
		// EtapaId:     incidenciaDto.Etapa,
		//
	}
	return s.dao.Create(ctx, &incidencia)
}

// Actualizar incidencia existente
func (s *IncidenciaService) Update(ctx context.Context, incidenciaDto *dto.IncidenciaDto) error {
	parsedDate, err := time.Parse("2006-01-02", incidenciaDto.Date)
	if err != nil {
		return err
	}
	parsedDuration, err := strconv.Atoi(incidenciaDto.Duration)
	if err != nil {
		return err
	}
	incidencia := entity.Incidencia{
		Id:          incidenciaDto.Id,
		Type:        incidenciaDto.Type,
		Date:        parsedDate,
		Duration:    parsedDuration,
		Description: incidenciaDto.Description,
		Fase:        incidenciaDto.Fase,
		// ActividadId: incidenciaDto.ActividadId,
		// EtapaId:     incidenciaDto.Etapa,
	}
	return s.dao.Update(ctx, &incidencia)
}

// Eliminar incidencia
func (s *IncidenciaService) Delete(ctx context.Context, id int) error {
	return s.dao.Delete(ctx, id)
}
