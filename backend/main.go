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

func createAccountImplGen(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{})
	}
}

func loginImplGen(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{})
	}
}

func main() {

	db, err := sql.Open("sqlite3", "./mydata.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	DatabasePerformanceOptimisatioins(db)

	router := gin.Default()

	router.GET("/ping", pingImpl)
	router.POST("/createaccount", createAccountImplGen(db))
	router.POST("/createaccount", loginImplGen(db))

	router.Run(":8080")
}
