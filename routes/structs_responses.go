package routes

type Administraciones struct {
	ID              uint   `json:"id" gorm:"primaryKey"`
	Adeudo          bool   `json:"adeudo"`
	Estado          int    `json:"estado"`
	AlumnoNombre    string `json:"alumno_nombre"`
	AlumnoMatricula string `json:"alumno_matricula"`
}

type Alumnos struct {
	ID              uint               `json:"id" gorm:"primaryKey"`
	Nombre          string             `json:"nombre"`
	Apellidos       string             `json:"apellidos"`
	Matricula       string             `json:"matricula"`
	FechaNacimiento string             `json:"fecha_nacimiento"`
	Edad            uint               `json:"edad"`
	NombreTutor     string             `json:"nombre_tutor"`
	CelularTutor    string             `json:"celular_tutor"`
	Curp            string             `json:"curp"`
	Localidad       string             `json:"localidad"`
	CodigoPostal    string             `json:"codigo_postal"`
	Direccion       string             `json:"direccion"`
	TelefonoFijo    string             `json:"telefono_fijo"`
	Celular         string             `json:"celular"`
	Correo          string             `json:"correo"`
	Especialidad    []Especialidades   `json:"especialidad"`     //arreglo
	GrupoActivo     []GruposActivos    `json:"grupos_activos"`   //arreglo
	GruposAprobados []GruposConcluidos `json:"grupos_aprobados"` //arreglo
}

type Especialidades struct {
	ID           uint   `json:"id" gorm:"primaryKey"`
	Materia      string `json:"materia"`
	Especialidad string `json:"especialidad"`
}

type GruposActivos struct {
	ID              uint           `json:"id" gorm:"primaryKey"`
	Nombre          string         `json:"nombre"`
	CantidadAlumnos int            `json:"cantidad_de_alumnos"`
	Trimestre       int            `json:"trimestre"`
	ListaAsistencia string         `json:"lista_asistencia"`
	Especialidad    Especialidades `json:"especialidad"`
}

type GruposConcluidos struct {
	ID              uint           `json:"id" gorm:"primaryKey"`
	Nombre          string         `json:"nombre"`
	CantidadAlumnos int            `json:"cantidad_de_alumnos"`
	Especialidad    Especialidades `json:"especialidad"`
}

type Permisos struct {
	Permiso string `json:"permiso"`
}

type RelacionAlumnoGrupo struct {
	ID                   uint `json:"id"`
	AlumnoRefer          int  `json:"alumno_id"`
	GruposActivosRefer   int  `json:"grupo_activo_id"`
	GruposAprobadosRefer int  `json:"grupo_aprobado_id"`
	EspecialidadesRefer  int  `json:"especialidad_id"`
}

type Usuarios struct {
	ID       uint     `json:"id" gorm:"primaryKey"`
	Usuario  string   `json:"usuario"`
	Password string   `json:"password"`
	Permisos Permisos `json:"permisos"`
}
