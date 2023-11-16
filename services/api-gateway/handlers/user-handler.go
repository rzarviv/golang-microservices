package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

var (
	client           *resty.Client = nil
	logger           string        = ""
	endpointBaseAddr string        = "http://localhost:"
)

func InitUsersHandler(router *gin.Engine, usersApiPort string) error {
	_, err := strconv.ParseInt(usersApiPort, 10, strconv.IntSize)
	if err != nil {
		fmt.Println("cannot read users' api port from .env")
		return err
	}

	endpointBaseAddr = endpointBaseAddr + usersApiPort + "/"
	client = resty.New().SetCloseConnection(true).SetBaseURL(endpointBaseAddr)

	addRoutes(router)

	return nil
}

func addRoutes(router *gin.Engine) {
	users := router.Group("/user") // create user route groups /user/...
	{
		users.POST("/", handleUserCreationRequest)   // add create user route
		users.DELETE("/", handleUserDeletionRequest) // add create user route
	}
}

func handleUserCreationRequest(c *gin.Context) {
	msg, err := createUser(c.Request)
	if err != nil {
		fmt.Println("cannot create user")
		c.JSON(http.StatusBadRequest, gin.H{"message": "cannot create user"})
	}
	c.JSON(http.StatusOK, gin.H{"message": msg})
}

func handleUserDeletionRequest(c *gin.Context) {
	msg, err :=deleteUser(c.Request)
	if err != nil {
		fmt.Println("cannot delete user")
		c.JSON(http.StatusBadRequest, gin.H{"message": "cannot delete user"})
	}
	c.JSON(http.StatusOK, gin.H{"message": msg})
}

func createUser(req *http.Request) (string, error) {
	resp, err := client.R().
		SetHeader("Content-Type", req.Header.Get("Content-Type")).
		SetBody(req.Body).
		Post(endpointBaseAddr)

	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	if resp != nil {
		//TODO: Parse response and return custom message to client
	}
	return "user created successfully", nil
}

func deleteUser(req *http.Request) (string, error) {
	resp, err := client.R().
		SetHeader("Content-Type", req.Header.Get("Content-Type")).
		SetBody(req.Body).
		Delete(endpointBaseAddr)

	if err != nil {
		fmt.Println(err.Error())
		return "cannot delete user" ,err
	}

	if resp != nil {
		//TODO: Parse response and return custom message to client
	}

	return "user deleted successfully", nil
}
