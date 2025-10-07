package dto

type RolesDto struct {
	Id_rol      int    `json:"id"`
	Name        string `json:"name"`
	Usuarios_Id []int  `json:"usuarios_id"`
}
