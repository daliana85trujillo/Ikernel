package entity

import "time"

type Incidencia struct {
	Id          int    `gorm:"primaryKey;autoIncrement"`
	Type        string `gorm:"type:varchar(50)"`
	Description string `gorm:"type:text"`
	Fase        string `gorm:"type:varchar(50)"`
	Date        time.Time
	Duration    int
	//ActividadId int
	//Actividad   *Actividad `gorm:"ForeignKey:ActividadId;constraint:OnDelete:CASCADE;"`
	//EtapaId     int
	//Etapa       *Etapa `gorm:"ForeignKey:EtapaId;constraint:OnDelete:CASCADE;"`
}

func (i *Incidencia) GetId() int {
	return i.Id
}
