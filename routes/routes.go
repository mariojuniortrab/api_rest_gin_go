package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/guilhermeonrails/api-go-gin/controllers"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/alunos", controllers.ListAllStudents)
	r.GET("/:nome", controllers.Greetings)
	r.POST("/alunos", controllers.RegisterStudent)
	r.GET("/alunos/:id", controllers.DetailStudentByID)
	r.DELETE("/alunos/:id", controllers.RemoveStudent)
	r.PATCH("/alunos/:id", controllers.UpdateStudent)
	r.GET("/alunos/cpf/:cpf", controllers.DetailStudentByCPF)
	r.Run()
}
