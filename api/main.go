package main

import (
	"api/oura"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getOuraSnapshot(c *gin.Context) {
	snapshot, err := oura.GetSnapshot()
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusBadRequest, snapshot)
	} else {
		c.IndentedJSON(http.StatusOK, snapshot)
	}
}

func main() {
	router := gin.Default()
	router.GET("/oura", getOuraSnapshot)
	router.Run(":4041")
}
