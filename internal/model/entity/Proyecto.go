package entity

import "time"

type Proyecto struct {
	Id          int       `gorm:"primaryKey;autoIncrement"`
	Name        string    `gorm:"type:varchar(100)"`
	Description string    `gorm:"type:varchar(255)"`
	Status      string    `gorm:"type:varchar(50)"`
	Start_date  time.Time `gorm:"type:date"`
	End_date    time.Time `gorm:"type:date"`
	Objective   string    `gorm:"type:varchar(100)"`
	//Usuario     []Usuario `gorm:"many2many:usuario_proyecto;uniqueIndex:idx_usuario_proyecto"`
}

func (p *Proyecto) GetId() int {
	return p.Id
}
