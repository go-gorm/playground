package v3

// Canton es una tabla de cantones
type Canton struct {
	ID        int    `json:"id" gorm:"primary_key"`
	Provincia int    `json:"provincia" gorm:"not null"`
	Canton    int    `json:"canton" gorm:"not null"`
	Nombre    string `json:"nombre" gorm:"not null"`
}

func (c Canton) TableName() string {
	return "cantones"
}

// Distrito es una tabla de distritos, su relación con la tabla de Canton es
// implícita por el ID de Canton
type Distrito struct {
	ID        int    `json:"id" gorm:"primary_key"`
	Provincia int    `json:"provincia" gorm:"not null"`
	Canton    int    `json:"canton" gorm:"not null"`
	Distrito  int    `json:"distrito" gorm:"not null"`
	Nombre    string `json:"nombre" gorm:"not null"`
}

// Barrio es una tabla de barrios, su relación con la tabla de Barrio es
// implícita por el ID de distrito
type Barrio struct {
	ID        int    `json:"id" gorm:"primary_key"`
	Provincia int    `json:"provincia" gorm:"not null"`
	Canton    int    `json:"canton" gorm:"not null"`
	Distrito  int    `json:"distrito" gorm:"not null"`
	Barrio    int    `json:"barrio" gorm:"not null"`
	Nombre    string `json:"nombre" gorm:"not null"`
}
