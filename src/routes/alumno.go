package routes

// func GetAllAlumnos(c *fiber.Ctx) error {
// 	alumnos := []models.Alumnos{}
// 	relaciones := []models.RelacionAlumnoGrupo{}

// 	var grupoActivo models.GruposActivos
// 	var grupoAprobado models.GruposConcluidos

// 	//var relacion models.RelacionAlumnoGrupo

// 	responseAlumnos := []Alumnos{}
// 	responseGruposActivos := []GruposActivos{}
// 	responseGruposAprobados := []GruposConcluidos{}

// 	database.Database.Db.Find(&alumnos)

// 	for _, alumno := range alumnos {
// 		/*
// 			if err := findAlumnoRelacion(alumno.ID, &relacion); err != nil {
// 				return c.Status(400).JSON(err.Error())
// 			}*/
// 		database.Database.Db.Find(&relaciones, models.RelacionAlumnoGrupo{AlumnoRefer: 1})
// 		for _, relacion := range relaciones {

// 			/*
// 				if err := findGrupoActivo(relacion.GruposActivosRefer, &grupoActivo); err != nil {
// 					return c.Status(400).JSON(err.Error())
// 				}*/

// 			database.Database.Db.Find(&grupoActivo, "id = ?", relacion.GruposActivosRefer)
// 			database.Database.Db.Find(&grupoAprobado, "id = ?", relacion.GruposAprobadosRefer)

// 			//log.Print(relacion.GruposActivosRefer)

// 			responseGrupoActivo := CreateGruposActivosResponse(grupoActivo)
// 			responseGruposActivos = append(responseGruposActivos, responseGrupoActivo)

// 			responseGrupoAprobado := CreateGruposConcluidosResponse(grupoAprobado)
// 			responseGruposAprobados = append(responseGruposAprobados, responseGrupoAprobado)
// 			//log.Print(responseGruposActivos)
// 		}
// 		// database.Database.Db.Find(&gruposActivos)
// 		// database.Database.Db.Find(&gruposAprobados)
// 		responseAlumno := CreateGetAllAlumnosResponse(alumno, responseGruposActivos, responseGruposAprobados)
// 		responseAlumnos = append(responseAlumnos, responseAlumno)

// 	}
// 	return c.Status(200).JSON(responseAlumnos)

// }
