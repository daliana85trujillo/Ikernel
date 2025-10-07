package entity

type Rol struct {
	Id_rol   int    `gorm:"primaryKey;autoIncrement"`
	Name_rol string `gorm:"unique;type:varchar(50)"`
	//Usuario  []Usuario `gorm:"many2many:usuario_rol;uniqueIndex:idx_usuario_rol"`
}
