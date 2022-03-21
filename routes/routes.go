package routes

import (
	studentController "golang-mock/controllers"
	studentControllerApi "golang-mock/controllers/api"
	"os"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {

	r := gin.Default()

	r.LoadHTMLGlob("templates/*.html")

	urlPrefix := r.Group("/pocdotnetcore")
	{
		urlPrefix.GET("/", studentController.GetAllStudents)
		urlPrefix.GET("/new-student", studentController.NewStudent)
		urlPrefix.POST("/create-student", studentController.CreateStudent)
		urlPrefix.GET("/:id", studentController.GetStudentById)
		urlPrefix.GET("/cpf/:cpf", studentController.GetStudentByCPF)
		urlPrefix.GET("/edit-student", studentController.EditStudent)
		urlPrefix.POST("/update-student", studentController.UpdateStudent)
		urlPrefix.GET("/delete-student", studentController.DeleteStudent)

		api := urlPrefix.Group("/api")
		{
			api.GET("/:name", studentControllerApi.Greeting)
			api.GET("/students", studentControllerApi.GetAllStudents)
			api.GET("/students/:id", studentControllerApi.GetStudentById)
			api.GET("/students/cpf/:cpf", studentControllerApi.GetStudentByCPF)
			api.POST("/students", studentControllerApi.CreateStudent)
			api.PATCH("/students/:id", studentControllerApi.UpdateStudent)
			api.DELETE("/students/:id", studentControllerApi.DeleteStudent)
		}
	}

	port := "8080"
	if os.Getenv("ASPNETCORE_PORT") != "" { // get enviroment variable that set by ACNM
		port = os.Getenv("ASPNETCORE_PORT")
	}
	r.Run(":" + port)
}
