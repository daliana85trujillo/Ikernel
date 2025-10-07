package entity

type Actividad struct {
	Id          int    `gorm:"primaryKey;autoIncrement"`
	Description string `gorm:"type:varchar(255)"`
	Status      string `gorm:"type:varchar(50)"`
	// EtapaId     int
	// Etapa       *Etapa `gorm:"ForeignKey:EtapaId;constraint:OnDelete:CASCADE;"`
}

func (a *Actividad) GetId() int {
	return a.Id
}
