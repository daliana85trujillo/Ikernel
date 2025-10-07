package dao

import "gorm.io/gorm"

type EtapaDao struct {
	// Aquí puedes agregar dependencias como la conexión a la base de datos, etc.
	db *gorm.DB
}

func NewEtapaDao(db *gorm.DB) *EtapaDao {
	return &EtapaDao{db: db}
}

// Aquí puedes agregar métodos para interactuar con la base de datos, por ejemplo:
func (e *EtapaDao) FindAll() ([]string, error) {
	// Lógica para obtener todas las etapas desde la base de datos
	return []string{"Etapa 1", "Etapa 2", "Etapa 3"}, nil
}

func (e *EtapaDao) Create(etapa string) error {
	// Lógica para crear una nueva etapa en la base de datos
	return nil
}

func (e *EtapaDao) FindById(id int) (string, error) {
	// Lógica para obtener una etapa por ID desde la base de datos
	return "Etapa " + string(rune(id)), nil
}

func (e *EtapaDao) Update(id int, etapa string) error {
	// Lógica para actualizar una etapa existente en la base de datos
	return nil
}
func (e *EtapaDao) Delete(id int) error {
	// Lógica para eliminar una etapa por ID desde la base de datos
	return nil
}
