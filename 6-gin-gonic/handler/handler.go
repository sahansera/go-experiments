package handler

import (
	"fmt"
	"gintest/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// curl http://localhost:8080/user-api/user
func GetUsers(c *gin.Context) {
	var user []model.User
	err := model.GetAllUsers(&user)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

// curl -X POST http://localhost:8080/user-api/user \
// -H "Content-Type: application/json" \
// -d '{
//   "name": "Test"
// }
// '
func CreateUser(c *gin.Context) {
	var user model.User
	c.BindJSON(&user)
	err := model.CreateUser(&user)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}
