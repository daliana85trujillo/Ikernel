package service

import (
	"Ikernel/internal/dao"
	"Ikernel/internal/model/dto"
	"Ikernel/internal/model/entity"
	"context"
	"errors"
	"time"
)

type UsuarioService struct {
	dao dao.Idao[entity.Usuario, int]
}

func NewUsuarioService(dao dao.Idao[entity.Usuario, int]) *UsuarioService {
	return &UsuarioService{dao: dao}
}

func (s *UsuarioService) ValidarCredenciales(ctx context.Context, email string, password string) (*entity.Usuario, error) {
	usuarios, err := s.dao.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	for _, u := range usuarios {
		if u.Email == email && u.Password == password {
			return &u, nil
		}
	}

	return nil, errors.New("credenciales inválidas")
}

func (s *UsuarioService) FindAll(ctx context.Context) ([]entity.Usuario, error) {
	return s.dao.FindAll(ctx)

}

func (s *UsuarioService) FindById(ctx context.Context, id int) (*entity.Usuario, error) {
	return s.dao.FindById(ctx, id)
}

func (s *UsuarioService) Create(ctx context.Context, usuario *entity.Usuario) error {
	return s.dao.Create(ctx, usuario)
}

func (s *UsuarioService) Update(ctx context.Context, usuario *dto.UsuarioDto) error {
	birthDate, err := time.Parse("2006-01-02", usuario.Birth_date)
	if err != nil {
		return err
	}
	usuario.Birth_date = birthDate.Format("2006-01-02")
	entidad := entity.Usuario{
		Id_Usuario:     usuario.Id_Usuario,
		Name:           usuario.Name,
		Last_name:      usuario.Last_name,
		Birth_date:     birthDate,
		Identification: usuario.Identification,
		Address:        usuario.Address,
		Profession:     usuario.Profession,
		Specialty:      usuario.Specialty,
		Photo:          usuario.Photo,
		Email:          usuario.Email,
	}
	return s.dao.Update(ctx, &entidad)
}
