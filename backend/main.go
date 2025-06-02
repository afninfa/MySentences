package main

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func pingImpl(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func main() {

	db := Unwrap(sql.Open("sqlite3", "./mydata.db"))
	defer db.Close()

	DatabasePerformanceOptimisatioins(db)

	router := gin.Default()

	router.GET("/ping", pingImpl)

	router.Run(":8080")
}
