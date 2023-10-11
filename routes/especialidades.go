package routes

type Especialidades struct {
	ID           uint   `json:"id" gorm:"primaryKey"`
	Materia      string `json:"materia"`
	Especialidad string `json:"especialidad"`
}
