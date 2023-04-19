package users

import (
	"bookstore_users-api/domain/users"
	"bookstore_users-api/services"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

func GetUser(c *gin.Context) {

}

func CreateUser(c *gin.Context) {
	var user users.User
	bytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		// TODO: Handle error
		return
	}
	if err := json.Unmarshal(bytes, &user); err != nil {
		// TODO: Handle error
		return
	}
	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		// TODO: Handle error
		return
	}
	c.JSON(http.StatusCreated, result)
}
