package service

import (
	"Ikernel/internal/dao"
	"Ikernel/internal/model/dto"
	"Ikernel/internal/model/entity"
	"context"
	"time"
)

type ProyectoService struct {
	dao *dao.ProyectoDao
}

func NewProyectoService(dao *dao.ProyectoDao) *ProyectoService {
	return &ProyectoService{dao: dao}
}

// ======================= CRUD =======================

func (s *ProyectoService) Create(ctx context.Context, proyectoDto *dto.ProyectoDto) error {
	startDate, err := time.Parse("2006-01-02", proyectoDto.Start_date)
	if err != nil {
		return err
	}
	endDate, err := time.Parse("2006-01-02", proyectoDto.End_date)
	if err != nil {
		return err
	}
	proyecto := entity.Proyecto{
		Name:        proyectoDto.Name,
		Description: proyectoDto.Description,
		Objective:   proyectoDto.Objective,
		Status:      proyectoDto.Status,
		Start_date:  startDate,
		End_date:    endDate,
		// Usuario field should be set here if you want to associate users, e.g.:
		// Usuario: usuariosSlice, // You need to fetch []*entity.Usuario by IDs if needed
	}
	return s.dao.Create(ctx, &proyecto)
}

// Actualizar
func (s *ProyectoService) Update(ctx context.Context, proyectoDto *dto.ProyectoDto) error {
	startDate, err := time.Parse("2006-01-02", proyectoDto.Start_date)
	if err != nil {
		return err
	}
	endDate, err := time.Parse("2006-01-02", proyectoDto.End_date)
	if err != nil {
		return err
	}
	proyecto := entity.Proyecto{
		Id:          proyectoDto.Id,
		Name:        proyectoDto.Name,
		Description: proyectoDto.Description,
		Objective:   proyectoDto.Objective,
		Status:      proyectoDto.Status,
		Start_date:  startDate,
		End_date:    endDate,
		// Set user association here if needed, e.g.:
		// Usuario: usuariosSlice, // You need to fetch []*entity.Usuario by IDs if needed
	}
	return s.dao.Update(ctx, &proyecto)
}

// Eliminar
func (s *ProyectoService) Delete(ctx context.Context, id int) error {
	return s.dao.Delete(ctx, id)
}

// Buscar por ID
func (s *ProyectoService) FindById(ctx context.Context, id int) (*dto.ProyectoDto, error) {
	proyecto, err := s.dao.FindById(ctx, id)
	if err != nil {
		return nil, err
	}
	return &dto.ProyectoDto{
		Id:          proyecto.Id,
		Name:        proyecto.Name,
		Description: proyecto.Description,
		Objective:   proyecto.Objective,
		Status:      proyecto.Status,
		Start_date:  proyecto.Start_date.Format("2006-01-02"),
		End_date:    proyecto.End_date.Format("2006-01-02"),
	}, nil
}

// Listar todos
func (s *ProyectoService) FindAll(ctx context.Context) ([]dto.ProyectoDto, error) {
	proyectos, err := s.dao.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	var proyectosDto []dto.ProyectoDto
	for _, p := range proyectos {
		proyectosDto = append(proyectosDto, dto.ProyectoDto{
			Id:          p.Id,
			Name:        p.Name,
			Description: p.Description,
			Objective:   p.Objective,
			Status:      p.Status,
			Start_date:  p.Start_date.Format("2006-01-02"),
			End_date:    p.End_date.Format("2006-01-02"),
		})
	}
	return proyectosDto, nil
}

// ======================= FILTROS =======================

func (s *ProyectoService) FindByFiltro(ctx context.Context, filtro dto.FiltroProyectoDto) ([]dto.ProyectoDto, error) {
	proyectos, err := s.dao.FindByFiltro(ctx, filtro)
	if err != nil {
		return nil, err
	}

	var proyectosDto []dto.ProyectoDto
	for _, p := range proyectos {
		proyectosDto = append(proyectosDto, dto.ProyectoDto{
			Id:          p.Id,
			Name:        p.Name,
			Description: p.Description,
			Objective:   p.Objective,
			Status:      p.Status,
			Start_date:  p.Start_date.Format("2006-01-02"),
			End_date:    p.End_date.Format("2006-01-02"),
		})
	}
	return proyectosDto, nil
}
