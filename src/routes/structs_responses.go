package routes

type Administraciones struct {
	ID        int    `json:"id" gorm:"primaryKey"`
	Nombre    string `json:"nombre"`
	Apellido  string `json:"apellido"`
	Matricula string `json:"matricula"`
	Adeudo    string `json:"adeudo"`
	Estado    string `json:"estado"`
}

type CicloEscolar struct {
	ID        int    `json:"id" gorm:"primaryKey"`
	Nombre    string `json:"nombre"`
	Year      string `json:"year"`
	Trimestre int    `json:"trimestre"`
	Activo    bool   `json:"activo"`
}

type HistorialLogs struct {
	ID         int            `json:"id" gorm:"primaryKey"`
	Movimiento string         `json:"movimiento"`
	Monto      int            `json:"monto"`
	Hora       string         `json:"hora"`
	Fecha      string         `json:"fecha"`
	Alumno     AlumnosNombres `json:"alumno"`
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

type AlumnosNombres struct {
	ID           int              `json:"id" gorm:"primaryKey"`
	Nombre       string           `json:"nombre"`
	Apellidos    string           `json:"apellidos"`
	Matricula    string           `json:"matricula"`
	Especialidad []Especialidades `json:"especialidad"`
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
	Horario         []Horarios     `json:"horario"`
	CantidadAlumnos int            `json:"cantidad_de_alumnos"`
	Trimestre       int            `json:"trimestre"`
	ListaAsistencia []string       `json:"lista_asistencia"`
	CicloEscolar    CicloEscolar   `json:"ciclo_escolar"`
	Especialidad    Especialidades `json:"especialidad"`
}

type Horarios struct {
	ID      int    `json:"id" gorm:"primaryKey"`
	Dia     string `json:"dia"`
	Entrada string `json:"entrada"`
	Salida  string `json:"salida"`
	DiaData int    `json:"diaData"`
}

type GruposConcluidos struct {
	ID              int            `json:"id" gorm:"primaryKey"`
	Estado          string         `json:"estado"`
	Nombre          string         `json:"nombre"`
	CantidadAlumnos int            `json:"cantidad_de_alumnos"`
	Especialidad    Especialidades `json:"especialidad"`
	CicloEscolar    CicloEscolar   `json:"ciclo_escolar"`
}

type Permisos struct {
	ID             int    `json:"id" gorm:"primaryKey"`
	Permiso        string `json:"permiso"`
	Administracion bool   `json:"administracion"`
	ControlEscolar bool   `json:"control_escolar"`
	Administrador  bool   `json:"administrador"`
	Inscripcion    bool   `json:"inscripcion"`
}

type RelacionAlumnoGrupo struct {
	ID                   int    `json:"id"`
	Estado               string `json:"estado"`
	AlumnoRefer          int    `json:"alumno_id"`
	GruposActivosRefer   int    `json:"grupo_activo_id"`
	GruposAprobadosRefer int    `json:"grupo_aprobado_id"`
	EspecialidadesRefer  int    `json:"especialidad_id"`
}

type Usuarios struct {
	ID       int      `json:"id" gorm:"primaryKey"`
	Usuario  string   `json:"usuario"`
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

const (
	NoAplica = iota
	Proceso
	Desercion
	Acreditacion
	NoAcreditacion
)
