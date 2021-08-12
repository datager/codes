package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func main() {
	i := 0
	go func() {
		r1 := gin.Default()
		r1.GET("/", func(c *gin.Context) {
			c.JSON(200, fmt.Sprintf("[%v] running, %v", i, time.Now().String()))
		})

		if err := r1.Run(":8891"); err != nil {
			log.Print("error", err)
		}
	}()

	for {
		i++
		<-time.After(1 * time.Second)
		log.Printf("[%v] running, %v", i, time.Now().String())
	}
}
