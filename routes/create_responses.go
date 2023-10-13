package routes

import "github.com/Gerardo02/Back-Fiber/models"

func CreateResponseAdmin(adminModel models.Administraciones, alumnoNombre string, alumnoMatricula string) Administraciones {
	return Administraciones{
		ID:              adminModel.ID,
		Adeudo:          adminModel.Adeudo,
		Estado:          adminModel.Estado,
		AlumnoNombre:    alumnoNombre,
		AlumnoMatricula: alumnoMatricula,
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

func CreateEspecialidadResponse(especialidadModel models.Especialidades) Especialidades {
	return Especialidades{
		ID:           especialidadModel.ID,
		Materia:      especialidadModel.Materia,
		Especialidad: especialidadModel.Especialidad,
	}
}

func CreateGruposActivosResponse(gruposActivosModel models.GruposActivos, especialidad Especialidades) GruposActivos {
	return GruposActivos{
		ID:              gruposActivosModel.ID,
		Nombre:          gruposActivosModel.Nombre,
		CantidadAlumnos: gruposActivosModel.CantidadAlumnos,
		Trimestre:       gruposActivosModel.Trimestre,
		ListaAsistencia: gruposActivosModel.ListaAsistencia,
		Especialidad:    especialidad,
	}
}

func CreateGruposConcluidosResponse(gruposConcluidosModel models.GruposConcluidos, especialidad Especialidades) GruposConcluidos {
	return GruposConcluidos{
		ID:              gruposConcluidosModel.ID,
		Nombre:          gruposConcluidosModel.Nombre,
		CantidadAlumnos: gruposConcluidosModel.CantidadAlumnos,
		Especialidad:    especialidad,
	}
}

func CreatePermisosResponse(permisosModel models.Permisos) Permisos {
	return Permisos{Permiso: permisosModel.Permiso}
}

func CreateRelacionResponse(relacionModel models.RelacionAlumnoGrupo) RelacionAlumnoGrupo {
	return RelacionAlumnoGrupo{
		ID:                   relacionModel.ID,
		AlumnoRefer:          relacionModel.AlumnoRefer,
		GruposActivosRefer:   relacionModel.GruposActivosRefer,
		GruposAprobadosRefer: relacionModel.GruposAprobadosRefer,
		EspecialidadesRefer:  relacionModel.EspecialidadRefer,
	}
}

func CreateUsuariosResponse(usuariosModel models.Usuarios, permiso Permisos) Usuarios {
	return Usuarios{
		ID:       usuariosModel.ID,
		Usuario:  usuariosModel.Usuario,
		Password: usuariosModel.Password,
		Permisos: permiso,
	}
}
