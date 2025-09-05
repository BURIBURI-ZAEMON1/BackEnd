package main

import (
	"backend/conf/database"
	"backend/conf/router"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Init()
	r := gin.Default()
	router.Init(r)
	err := r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
