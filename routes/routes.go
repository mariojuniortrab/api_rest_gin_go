package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mariojuniortrab/api-rest-gin-go/controllers"
)

func HandleRequests() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./assets")

	r.GET("/students", controllers.ListAllStudents)
	r.GET("/:name", controllers.Greetings)
	r.POST("/students", controllers.RegisterStudent)
	r.GET("/students/:id", controllers.DetailStudentByID)
	r.DELETE("/students/:id", controllers.RemoveStudent)
	r.PATCH("/students/:id", controllers.UpdateStudent)
	r.GET("/students/cpf/:cpf", controllers.DetailStudentByCPF)
	r.GET("/index", controllers.ShowIndexPage)
	r.NoRoute(controllers.ShowPageNotFound)
	r.Run()
}
