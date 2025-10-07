package dto

type FiltroProyectoDto struct {
	Name       string `json:"name,omitempty"`
	Status     string `json:"status,omitempty"`
	Start_date string `json:"Start_date,omitempty"`
	End_date   string `json:"End_date,omitempty"`
	LiderID    int    `json:"lider_id,omitempty"`
}
