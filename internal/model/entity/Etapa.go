package entity

import "time"

type Etapa struct {
	Id          int    `gorm:"primaryKey;autoIncrement"`
	Name        string `gorm:"type:varchar(100)"`
	Description string `gorm:"type:varchar(255)"`
	Start_date  time.Time
	End_date    time.Time
	State       string `gorm:"type:varchar(50)"`
	//ProyectoId  int
	//Proyecto    *Proyecto `gorm:"ForeignKey:ProyectoId;constraint:OnDelete:CASCADE;"`
	//Actividad   []Actividad
	//Incidencia  []Incidencia
}

func (e *Etapa) GetId() int {
	return e.Id
}
