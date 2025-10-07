package dao

import (
	"context"
)

type Idao[T any, ID comparable] interface {
	FindAll(ctx context.Context) ([]T, error)
	FindById(ctx context.Context, id ID) (*T, error)
	Create(ctx context.Context, entity *T) error
	Update(ctx context.Context, entity *T) error
	Delete(ctx context.Context, entity *T) error
	// New method to find with filters

	//FindWithFilters(filter dto.UsuarioFilter) ([]T, error)
}
