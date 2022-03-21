package controllers

import (
	"fmt"
	"golang-mock/database"
	"golang-mock/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllStudents(c *gin.Context) {
	var students []models.Student
	database.DB.Raw("SELECT * FROM students WHERE deleted_at IS NULL").Scan(&students)

	c.HTML(http.StatusOK, "Index", students)
}

func NewStudent(c *gin.Context) {
	c.HTML(http.StatusOK, "New-student", nil)
}

func CreateStudent(c *gin.Context) {
	var student models.Student

	if err := c.ShouldBind(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error()})
		return
	}
	fmt.Println(student)
	database.DB.Exec("INSERT INTO students (created_at, updated_at, name, cpf, rg) VALUES (GETDATE(),GETDATE(),?,?,?)", student.Name, student.CPF, student.RG)

	c.Redirect(http.StatusMovedPermanently, "/pocdotnetcore/")
}

func EditStudent(c *gin.Context) {
	var student models.Student
	id := c.Query("id")
	database.DB.Raw("SELECT top 1 * FROM students WHERE id = ?", id).Scan(&student)

	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Student not found"})
		return
	}

	c.HTML(http.StatusOK, "Edit-student", student)
}

func GetStudentById(c *gin.Context) {
	var student []models.Student
	id := c.Params.ByName("id")
	database.DB.Raw("SELECT top 1 * FROM students WHERE id = ? AND deleted_at = NULL", id).Scan(&student)

	c.HTML(http.StatusOK, "Index", student)
}

func GetStudentByCPF(c *gin.Context) {
	var student []models.Student
	cpf := c.Param("cpf")
	database.DB.Raw("SELECT top 1 * FROM students WHERE cpf = ? AND deleted_at = NULL", cpf).Scan(&student)

	c.HTML(http.StatusOK, "Index", student)
}

func DeleteStudent(c *gin.Context) {
	id := c.Query("id")
	database.DB.Exec("DELETE FROM students WHERE id = ?", id)

	c.Redirect(http.StatusMovedPermanently, "/pocdotnetcore/")
}

func UpdateStudent(c *gin.Context) {
	var student models.Student

	if err := c.ShouldBind(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error()})
		return
	}

	database.DB.Exec("UPDATE students SET name = ?, cpf = ?, rg = ? WHERE id = ? ", student.Name, student.CPF, student.RG, student.ID)
	c.Redirect(http.StatusMovedPermanently, "/pocdotnetcore/")
}
