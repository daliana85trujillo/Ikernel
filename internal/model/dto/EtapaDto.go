package dto

type EtapaDto struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Start_date  string `json:"start_date"`
	End_date    string `json:"end_date"`
	State       string `json:"state"`
	ProyectoId  int    `json:"proyecto_id"`
}
