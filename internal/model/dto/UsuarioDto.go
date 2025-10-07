package dto

type UsuarioDto struct {
	Id_Usuario     int           `json:"id"`
	Name           string        `json:"name"`
	Last_name      string        `json:"last_name"`
	Birth_date     string        `json:"birth_date "`
	Identification string        `json:"identification "`
	Address        string        `json:"address"`
	Profession     string        `json:"profession "`
	Specialty      string        `json:"specialty"`
	Email          string        `json:"email"`
	Status         string        `json:"status"`
	Photo          []byte        `json:"photo"`
	Password       string        `json:"password"`
	Proyectos      []ProyectoDto `json:"proyectos"`
	Roles          []RolesDto    `json:"roles"`
	RolesId        []int
	ProyectosId    []int
}

func (u *UsuarioDto) ToEntity() any {
	panic("unimplemented")
}
