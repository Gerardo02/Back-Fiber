package routes

import "github.com/Gerardo02/Back-Fiber/models"

func CreateGetAdminResponse(adminModel models.Administraciones, nombre string, apellido string, matricula string, adeudo string, estado string) Administraciones {
	return Administraciones{
		ID:        adminModel.ID,
		Nombre:    nombre,
		Apellido:  apellido,
		Matricula: matricula,
		Adeudo:    adeudo,
		Estado:    estado,
	}
}

func CreateHistorialResponse(historial models.HistorialLogs, alumno AlumnosNombres) HistorialLogs {
	return HistorialLogs{
		ID:         historial.ID,
		Movimiento: historial.Movimiento,
		Monto:      historial.Monto,
		Hora:       historial.Hora,
		Fecha:      historial.Fecha,
		Alumno:     alumno,
	}
}

func CreateGetCicloEscolarResponse(cicloModel models.CicloEscolar) CicloEscolar {
	return CicloEscolar{
		ID:        cicloModel.ID,
		Nombre:    cicloModel.Nombre,
		Year:      cicloModel.Year,
		Trimestre: cicloModel.Trimestre,
		Activo:    cicloModel.Activo,
	}
}

func CreateAdminResponse(adminModel models.Administraciones, adeudo string, estado string) Administraciones {
	return Administraciones{
		ID:     adminModel.ID,
		Adeudo: adeudo,
		Estado: estado,
	}
}

func CreateGetAllAlumnosResponse(alumnoModel models.Alumnos, grupoActivo []GruposActivos, gruposAprobados []GruposConcluidos, especialidades []Especialidades) Alumnos {
	return Alumnos{
		ID:              alumnoModel.ID,
		Nombre:          alumnoModel.Nombre,
		Apellidos:       alumnoModel.Apellidos,
		Matricula:       alumnoModel.Matricula,
		Edad:            alumnoModel.Edad,
		FechaNacimiento: alumnoModel.FechaNacimiento,
		NombreTutor:     alumnoModel.NombreTutor,
		CelularTutor:    alumnoModel.CelularTutor,
		Curp:            alumnoModel.Curp,
		Localidad:       alumnoModel.Localidad,
		CodigoPostal:    alumnoModel.CodigoPostal,
		Direccion:       alumnoModel.Direccion,
		TelefonoFijo:    alumnoModel.TelefonoFijo,
		Celular:         alumnoModel.Celular,
		Correo:          alumnoModel.Correo,
		Especialidad:    especialidades,
		GrupoActivo:     grupoActivo,
		GruposAprobados: gruposAprobados,
	}
}

func CreateAlumnosResponse(alumnoModel models.Alumnos) Alumnos {
	return Alumnos{
		ID:              alumnoModel.ID,
		Nombre:          alumnoModel.Nombre,
		Apellidos:       alumnoModel.Apellidos,
		Matricula:       alumnoModel.Matricula,
		Edad:            alumnoModel.Edad,
		FechaNacimiento: alumnoModel.FechaNacimiento,
		NombreTutor:     alumnoModel.NombreTutor,
		CelularTutor:    alumnoModel.CelularTutor,
		Curp:            alumnoModel.Curp,
		Localidad:       alumnoModel.Localidad,
		CodigoPostal:    alumnoModel.CodigoPostal,
		Direccion:       alumnoModel.Direccion,
		TelefonoFijo:    alumnoModel.TelefonoFijo,
		Celular:         alumnoModel.Celular,
		Correo:          alumnoModel.Correo,
	}
}

func CreateAlumnoNombreResponse(alumnoModel models.Alumnos, especialidades []Especialidades) AlumnosNombres {
	return AlumnosNombres{
		ID:           alumnoModel.ID,
		Nombre:       alumnoModel.Nombre,
		Apellidos:    alumnoModel.Apellidos,
		Matricula:    alumnoModel.Matricula,
		Especialidad: especialidades,
	}
}

func CreateAlumnoNombreNombreResponse(alumnoModel models.Alumnos) AlumnosNombres {
	return AlumnosNombres{
		ID:        alumnoModel.ID,
		Nombre:    alumnoModel.Nombre,
		Apellidos: alumnoModel.Apellidos,
		Matricula: alumnoModel.Matricula,
	}
}

func CreateEspecialidadResponse(especialidadModel models.Especialidades) Especialidades {
	return Especialidades{
		ID:           especialidadModel.ID,
		Materia:      especialidadModel.Materia,
		Especialidad: especialidadModel.Especialidad,
	}
}

func CreateGruposActivosResponse(gruposActivosModel models.GruposActivos, especialidad Especialidades, listaAsistencia []string, cicloEscolar CicloEscolar) GruposActivos {
	return GruposActivos{
		ID:              gruposActivosModel.ID,
		Nombre:          gruposActivosModel.Nombre,
		NombreMaestro:   gruposActivosModel.NombreMaestro,
		Dia:             gruposActivosModel.Dia,
		Entrada:         gruposActivosModel.Entrada,
		Salida:          gruposActivosModel.Salida,
		CantidadAlumnos: gruposActivosModel.CantidadAlumnos,
		CicloEscolar:    cicloEscolar,
		Trimestre:       gruposActivosModel.Trimestre,
		ListaAsistencia: listaAsistencia,
		Especialidad:    especialidad,
	}
}

func CreateGruposActivosAlumnosResponse(gruposActivosModel models.GruposActivos, especialidad Especialidades) GruposActivos {
	return GruposActivos{
		ID:              gruposActivosModel.ID,
		Nombre:          gruposActivosModel.Nombre,
		NombreMaestro:   gruposActivosModel.NombreMaestro,
		Dia:             gruposActivosModel.Dia,
		Entrada:         gruposActivosModel.Entrada,
		Salida:          gruposActivosModel.Salida,
		CantidadAlumnos: gruposActivosModel.CantidadAlumnos,
		Trimestre:       gruposActivosModel.Trimestre,
		Especialidad:    especialidad,
	}
}

func CreateAlumnosGruposConcluidosResponse(gruposConcluidosModel models.GruposConcluidos, especialidad Especialidades) GruposConcluidos {
	return GruposConcluidos{
		ID:              gruposConcluidosModel.ID,
		Nombre:          gruposConcluidosModel.Nombre,
		CantidadAlumnos: gruposConcluidosModel.CantidadAlumnos,
		Especialidad:    especialidad,
	}
}

func CreateGruposConcluidosResponse(gruposConcluidosModel models.GruposConcluidos, especialidad Especialidades, cicloEscolar CicloEscolar) GruposConcluidos {
	return GruposConcluidos{
		ID:              gruposConcluidosModel.ID,
		Nombre:          gruposConcluidosModel.Nombre,
		CantidadAlumnos: gruposConcluidosModel.CantidadAlumnos,
		CicloEscolar:    cicloEscolar,
		Especialidad:    especialidad,
	}
}

func CreatePermisosResponse(permisosModel models.Permisos) Permisos {
	return Permisos{
		ID:             permisosModel.ID,
		Permiso:        permisosModel.Permiso,
		Administracion: permisosModel.Administracion,
		Administrador:  permisosModel.Administrador,
		Inscripcion:    permisosModel.Inscripcion,
		ControlEscolar: permisosModel.ControlEscolar,
	}
}

func CreateRelacionAlumnoGrupoResponse(relacionModel models.RelacionAlumnoGrupo) RelacionAlumnoGrupo {
	return RelacionAlumnoGrupo{
		ID:                   relacionModel.ID,
		AlumnoRefer:          relacionModel.AlumnoRefer,
		GruposActivosRefer:   relacionModel.GruposActivosRefer,
		GruposAprobadosRefer: relacionModel.GruposAprobadosRefer,
		EspecialidadesRefer:  relacionModel.EspecialidadRefer,
	}
}

func CreateRelacionGrupoListasResponse(relacionModel models.RelacionGrupoLista) RelacionGrupoLista {
	return RelacionGrupoLista{
		ID:                 relacionModel.ID,
		ListaAsistencia:    relacionModel.ListaAsistencia,
		GruposActivosRefer: relacionModel.GruposActivosRefer,
	}
}

func CreateUsuariosResponse(usuariosModel models.Usuarios, permiso Permisos) Usuarios {
	return Usuarios{
		ID:       usuariosModel.ID,
		Usuario:  usuariosModel.Usuario,
		Permisos: permiso,
	}
}

func CreateGetDocumentosResponse(documentsModel models.Documentos, nombre string, apellidos string, matricula string) Documentos {
	return Documentos{
		ID:                   documentsModel.ID,
		AlumnoRefer:          documentsModel.AlumnoRefer,
		Nombre:               nombre,
		Apellidos:            apellidos,
		Matricula:            matricula,
		ActaNacimiento:       documentsModel.ActaNacimiento,
		Curp:                 documentsModel.Curp,
		ComprobanteDomicilio: documentsModel.ComprobanteDomicilio,
		MayorQuince:          documentsModel.MayorQuince,
		Fotos:                documentsModel.Fotos,
	}
}

func CreateDocumentosResponse(documentsModel models.Documentos) Documentos {
	return Documentos{
		ID:                   documentsModel.ID,
		AlumnoRefer:          documentsModel.AlumnoRefer,
		ActaNacimiento:       documentsModel.ActaNacimiento,
		Curp:                 documentsModel.Curp,
		ComprobanteDomicilio: documentsModel.ComprobanteDomicilio,
		MayorQuince:          documentsModel.MayorQuince,
		Fotos:                documentsModel.Fotos,
	}
}
