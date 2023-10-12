package database

import (
	"log"

	"github.com/Gerardo02/Back-Fiber/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectDb() {
	db, err := gorm.Open(sqlite.Open("api.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connecto to the database! \n", err.Error())
	}

	log.Println("Connected to the database successfully")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running migrations")

	db.AutoMigrate(
		&models.Administraciones{},
		&models.Alumnos{},
		&models.Especialidades{},
		&models.GruposActivos{},
		&models.GruposConcluidos{},
		&models.Pagos{},
		&models.Permisos{},
		&models.Usuarios{},
		&models.RelacionAlumnoGrupo{},
	)

	Database = DbInstance{Db: db}

}
