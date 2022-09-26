package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/mariojuniortrab/api-rest-gin-go/controllers"
	"github.com/mariojuniortrab/api-rest-gin-go/database"
	"github.com/mariojuniortrab/api-rest-gin-go/models"
	"github.com/stretchr/testify/assert"
)

var ID int

func RoutesTestsSetup() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	rotas := gin.Default()
	return rotas
}

func RegisterStudentMock() {
	student := models.Student{Name: "any_name", CPF: "12345678902", RG: "any_rg"}
	database.DB.Create(&student)
	ID = int(student.ID)
}

func RemoveStudentMock() {
	var student models.Student
	database.DB.Delete(&student, ID)
}

func TestStatusCodeGreetings(t *testing.T) {
	r := RoutesTestsSetup()
	r.GET("/:name", controllers.Greetings)

	req, _ := http.NewRequest("GET", "/mario", nil)
	resp := httptest.NewRecorder()

	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code, "Should be equals")

}
func TestResponseGreetings(t *testing.T) {
	r := RoutesTestsSetup()
	r.GET("/:name", controllers.Greetings)

	req, _ := http.NewRequest("GET", "/mario", nil)
	resp := httptest.NewRecorder()

	r.ServeHTTP(resp, req)

	body, _ := ioutil.ReadAll(resp.Body)

	assert.Equal(t, `{"API says:":"Hello mario, how are you doing?"}`, string(body))
}

func TestListAllStudents(t *testing.T) {
	database.DatabaseConnect()

	RegisterStudentMock()

	r := RoutesTestsSetup()
	r.GET("/students", controllers.ListAllStudents)
	req, _ := http.NewRequest("GET", "/students", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	defer RemoveStudentMock()

}

func TestDetailStudentByCPF(t *testing.T) {
	database.DatabaseConnect()
	RegisterStudentMock()
	r := RoutesTestsSetup()

	r.GET("/students/cpf/:cpf", controllers.DetailStudentByCPF)

	req, _ := http.NewRequest("GET", "/students/cpf/12345678902", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	defer RemoveStudentMock()
}

func TestDetailStudentById(t *testing.T) {
	database.DatabaseConnect()
	RegisterStudentMock()
	r := RoutesTestsSetup()
	r.GET("/students/:id", controllers.DetailStudentByID)
	path := "/students/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("GET", path, nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	var studentMock models.Student
	json.Unmarshal(resp.Body.Bytes(), &studentMock)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Equal(t, "any_name", studentMock.Name)
	assert.Equal(t, "12345678902", studentMock.CPF)
	assert.Equal(t, "any_rg", studentMock.RG)

	defer RemoveStudentMock()
}

func TestRemoveStudent(t *testing.T) {
	database.DatabaseConnect()
	RegisterStudentMock()
	r := RoutesTestsSetup()
	r.DELETE("/students/:id", controllers.RemoveStudent)
	path := "/students/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("DELETE", path, nil)

	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	body, _ := ioutil.ReadAll(resp.Body)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Equal(t, `{"data":"Student was removed"}`, string(body))

}

func TestUpdateStudent(t *testing.T) {
	database.DatabaseConnect()
	RegisterStudentMock()
	r := RoutesTestsSetup()
	r.PATCH("/students/:id", controllers.UpdateStudent)
	student := models.Student{Name: "any_name2", CPF: "13345678902", RG: "any_rg2"}
	jsonValue, _ := json.Marshal(student)

	path := "/students/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("PATCH", path, bytes.NewBuffer(jsonValue))
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	var updatedStudent models.Student
	json.Unmarshal(resp.Body.Bytes(), &updatedStudent)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Equal(t, "any_name2", updatedStudent.Name)
	assert.Equal(t, "13345678902", updatedStudent.CPF)
	assert.Equal(t, "any_rg2", updatedStudent.RG)
}
