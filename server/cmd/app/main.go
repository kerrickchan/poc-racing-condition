package main

import (
	"fmt"
	"racing-condition/internal/database"
	"racing-condition/internal/model"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello",
		})
	})

	db := database.Conn()
	db.AutoMigrate(&model.Ticket{})
	counter := 0
	r.POST("/ticket", func(c *gin.Context) {
		counter++
		code := fmt.Sprintf("code%d", counter)
		insertTicket := &model.Ticket{Code: code, User: c.Request.FormValue("user")}
		db.Create(insertTicket)

		readTicket := &model.Ticket{}
		db.First(&readTicket, "code = ?", code)
		c.JSON(200, gin.H{
			"message": "success",
			"ticket":  readTicket,
		})
	})

	r.Run(":4000")
}
