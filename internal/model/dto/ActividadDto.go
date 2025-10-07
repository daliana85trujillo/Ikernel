package dto

import "Ikernel/internal/model/entity"

type ActividadDto struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
	Status      string `json:"etatus"`
	Etapaid     int    `json:"etapa_id"`
}

func (a *ActividadDto) ToEntity() *entity.Actividad {
	return &entity.Actividad{
		Id:          a.Id,
		Description: a.Description,
		Status:      a.Status,
		// EtapaId:     a.Etapaid,
	}
}
