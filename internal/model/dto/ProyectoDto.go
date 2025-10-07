package dto

type ProyectoDto struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      string `json:"status"`
	Start_date  string `json:"start_date"`
	End_date    string `json:"end_date"`
	Objective   string `json:"objetive"`
	UsuarioIds  []int  `json:"usuario_ids"`
}
