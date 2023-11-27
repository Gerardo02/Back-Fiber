package routes

type Administraciones struct {
	ID        int    `json:"id" gorm:"primaryKey"`
	Nombre    string `json:"nombre"`
	Apellido  string `json:"apellido"`
	Matricula string `json:"matricula"`
	Adeudo    string `json:"adeudo"`
	Estado    string `json:"estado"`
}

type Alumnos struct {
	ID              int                `json:"id" gorm:"primaryKey"`
	Nombre          string             `json:"nombre"`
	Apellidos       string             `json:"apellidos"`
	Matricula       string             `json:"matricula"`
	FechaNacimiento string             `json:"fecha_nacimiento"`
	Edad            int                `json:"edad"`
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
	ID           int    `json:"id" gorm:"primaryKey"`
	Materia      string `json:"materia"`
	Especialidad string `json:"especialidad"`
}

type GruposActivos struct {
	ID              int            `json:"id" gorm:"primaryKey"`
	Nombre          string         `json:"nombre"`
	NombreMaestro   string         `json:"nombre_maestro"`
	CantidadAlumnos int            `json:"cantidad_de_alumnos"`
	Trimestre       int            `json:"trimestre"`
	ListaAsistencia []string       `json:"lista_asistencia"`
	Especialidad    Especialidades `json:"especialidad"`
}

type GruposConcluidos struct {
	ID              int            `json:"id" gorm:"primaryKey"`
	Nombre          string         `json:"nombre"`
	CantidadAlumnos int            `json:"cantidad_de_alumnos"`
	Especialidad    Especialidades `json:"especialidad"`
}

type Permisos struct {
	ID      int    `json:"id" gorm:"primaryKey"`
	Permiso string `json:"permiso"`
}

type RelacionAlumnoGrupo struct {
	ID                   int `json:"id"`
	AlumnoRefer          int `json:"alumno_id"`
	GruposActivosRefer   int `json:"grupo_activo_id"`
	GruposAprobadosRefer int `json:"grupo_aprobado_id"`
	EspecialidadesRefer  int `json:"especialidad_id"`
}

type Usuarios struct {
	ID       int      `json:"id" gorm:"primaryKey"`
	Usuario  string   `json:"usuario"`
	Password string   `json:"password"`
	Permisos Permisos `json:"permisos"`
}

type RelacionGrupoLista struct {
	ID                 int    `json:"id" gorm:"primaryKey"`
	ListaAsistencia    string `json:"lista_asistencia"`
	GruposActivosRefer int    `json:"grupo_id"`
}

type Documentos struct {
	ID                   int    `json:"id" gorm:"primaryKey"`
	AlumnoRefer          int    `json:"alumno_id"`
	Nombre               string `json:"nombre"`
	Apellidos            string `json:"apellido"`
	Matricula            string `json:"matricula"`
	ActaNacimiento       bool   `json:"acta_de_nacimiento"`
	Curp                 bool   `json:"curp"`
	ComprobanteDomicilio bool   `json:"comprobante_de_domicilio"`
	MayorQuince          bool   `json:"mayor_quince"`
	Fotos                bool   `json:"fotos"`
}

const (
	Pendiente = iota
	EnProceso
	Listo
)
