package routes

import (
	"errors"
	"time"

	"github.com/Gerardo02/Back-Fiber/database"
	"github.com/Gerardo02/Back-Fiber/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Administracion struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"fecha_creacion"`
	UpdatedAt time.Time      `json:"fecha_actualizacion"`
	DeletedAt gorm.DeletedAt `json:"fecha_borrado"`
	Alumno    Alumno         `json:"alumno"`
	Debe      bool           `json:"debe"`
}

func CreateResponseAdmin(adminModel models.Administracion, alumno Alumno) Administracion {
	return Administracion{
		ID:        adminModel.ID,
		CreatedAt: adminModel.CreatedAt,
		UpdatedAt: adminModel.UpdatedAt,
		DeletedAt: adminModel.DeletedAt,
		Alumno:    alumno,
		Debe:      adminModel.Debe,
	}
}

func findAdmin(id int, admin *models.Administracion) error {
	database.Database.Db.Find(&admin, "id = ?", id)

	if admin.ID == 0 {
		return errors.New("Alumno does not exist")
	}

	return nil
}

func CreateAdmin(c *fiber.Ctx) error {
	var admin models.Administracion

	if err := c.BodyParser(&admin); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var alumno models.Alumno

	if err := findAlumno(admin.AlumnoRefer, &alumno); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var control models.ControlEscolar

	if err := findControl(alumno.ControlEscolarRefer, &control); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&admin)

	responseControl := CreateResponseControl(control)
	responseAlumno := CreateResponseAlumno(alumno, responseControl)
	responseAdmin := CreateResponseAdmin(admin, responseAlumno)

	return c.Status(200).JSON(responseAdmin)
}

func GetAllAdministracion(c *fiber.Ctx) error {
	admins := []models.Administracion{}

	database.Database.Db.Find(&admins)
	responseAdmins := []Administracion{}

	for _, admin := range admins {
		var alumno models.Alumno
		var control models.ControlEscolar

		if err := findAlumno(admin.AlumnoRefer, &alumno); err != nil {
			return c.Status(400).JSON(err.Error())
		}

		if err := findControl(alumno.ControlEscolarRefer, &control); err != nil {
			return c.Status(400).JSON(err.Error())
		}

		responseControl := CreateResponseControl(control)
		responseAlumno := CreateResponseAlumno(alumno, responseControl)
		responseAdmin := CreateResponseAdmin(admin, responseAlumno)
		responseAdmins = append(responseAdmins, responseAdmin)
	}

	return c.Status(200).JSON(responseAdmins)

}

func GetAdministracion(c *fiber.Ctx) error {

	id, err := c.ParamsInt("id")
	var admin models.Administracion

	if err != nil {
		return c.SendString("id not an int")
	}

	if err := findAdmin(id, &admin); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var alumno models.Alumno

	if err := findAlumno(admin.AlumnoRefer, &alumno); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var control models.ControlEscolar

	if err := findControl(alumno.ControlEscolarRefer, &control); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	responseControl := CreateResponseControl(control)
	responseAlumno := CreateResponseAlumno(alumno, responseControl)
	responseAdmin := CreateResponseAdmin(admin, responseAlumno)

	return c.Status(200).JSON(responseAdmin)

}

func UpadateAdmin(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var admin models.Administracion

	if err != nil {
		return c.SendString("id not an int")
	}

	if err := findAdmin(id, &admin); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	type UpdatedAdmin struct {
		Alumno int  `json:"alumno_id"`
		Debe   bool `json:"debe"`
	}

	var updateData UpdatedAdmin

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	admin.AlumnoRefer = updateData.Alumno
	admin.Debe = updateData.Debe
	database.Database.Db.Save(&admin)

	var alumno models.Alumno

	if err := findAlumno(admin.AlumnoRefer, &alumno); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var control models.ControlEscolar

	if err := findControl(alumno.ControlEscolarRefer, &control); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	responseControl := CreateResponseControl(control)
	responseAlumno := CreateResponseAlumno(alumno, responseControl)
	responseAdmin := CreateResponseAdmin(admin, responseAlumno)

	return c.Status(200).JSON(responseAdmin)

}
