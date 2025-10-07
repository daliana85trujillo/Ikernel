package entity

import (
	"time"
)

type Usuario struct {
	Id_Usuario     int       `gorm:"primaryKey autoincrement"`
	Name           string    `gorm:"type:varchar(45)"`
	Last_name      string    `gorm:"type:varchar(45)"`
	Birth_date     time.Time `gorm:"type:date"`
	Identification string    `gorm:"type:varchar(20);uniqueIndex"`
	Address        string    `gorm:"type:varchar(35)"`
	Profession     string    `gorm:"type:varchar(30)"`
	Specialty      string    `gorm:"type:varchar(100)"`
	Email          string    `gorm:"type:varchar(100);uniqueIndex"`
	Status         string    `gorm:"type:varchar(50)"`
	Photo          []byte    `gorm:"type:bytea"`
	Password       string    `gorm:"type:varchar(255)"`
	//Rol            []Rol      `gorm:"many2many:usuario_rol;uniqueIndex:idx_usuario_rol"`
	//Proyecto       []Proyecto `gorm:"many2many:usuario_proyecto;uniqueIndex:idx_usuario_proyecto"`
}

func (u *Usuario) GetId() int {
	return u.Id_Usuario
}
