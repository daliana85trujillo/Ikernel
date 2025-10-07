package dto

type IncidenciaDto struct {
	Id          int    `json:"id"`
	Type        string `json:"type"`
	Description string `json:"description"`
	Fase        string `json:"fase"`
	Date        string `json:"date"`
	Duration    string `json:"duration"`
	ActividadId int    `json:"actividad_id"`
	Etapa       int    `json:"etapa_id"`
}
