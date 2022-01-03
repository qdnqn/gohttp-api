package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Client struct {
	ClientId string `json:"clientId,omitempty"`
	IP       string `json:"ip,omitempty"`
}

type ClientsOnboard struct {
	ClientsOnboard int64 `json:"clientsOnboard"`
}

var clients = []Client{}

func getClientById(c *gin.Context) {
	ID := c.Param("id")
	for _, i := range clients {
		if i.ClientId == ID {
			c.IndentedJSON(http.StatusOK, i)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found!"})
}

func hasClients(c *gin.Context) {
	num := ClientsOnboard{
		ClientsOnboard: int64(len(clients)),
	}

	c.IndentedJSON(http.StatusOK, num)
}

func addClient(c *gin.Context) {
	var newUser Client

	if err := c.BindJSON(&newUser); err != nil {
		fmt.Println(err)
		return
	}

	clients = append(clients, newUser)

	c.IndentedJSON(http.StatusOK, clients)
}

func main() {
	router := gin.Default()

	router.GET("/client/:id", getClientById)
	router.GET("/hasClients", hasClients)
	router.POST("/addClient", addClient)

	router.Run(":8080")
}
