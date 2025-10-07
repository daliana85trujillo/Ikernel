package cmd

import (
	"Ikernel/internal/model/entity"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Migrate() *gorm.DB {
	dsn := "host=localhost user=postgres password=Dt851313 dbname=Ikernel port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// List of entities to migrate
	entities := []interface{}{
		&entity.Usuario{},
		&entity.Proyecto{},
		&entity.Etapa{},
		&entity.Actividad{},
		&entity.Incidencia{},
		&entity.Rol{},
	}

	// AutoMigrate the entities
	if err := db.AutoMigrate(entities...); err != nil {
		panic(err)
	}
	return db
}
