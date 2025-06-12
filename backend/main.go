package main

import (
	"database/sql"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func pingImpl(c *gin.Context) {
	session := sessions.Default(c)
	if user := session.Get("user"); user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorised",
		})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func createAccountImplGen(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data SignupInput
		// JSON fails to parse
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid input: " + err.Error(),
			})
			return
		}
		// User fails to insert
		err := CallInsertUserQuery(
			db,
			data.Email,
			data.Password,
			data.TargetLanguage,
		)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		// Success
		c.JSON(http.StatusOK, gin.H{
			"message": "Account created successfully!",
		})
	}
}

func loginImplGen(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data LoginInput
		// JSON fails to parse
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid input: " + err.Error(),
			})
			return
		}
		// User fails to login
		err := CallPasswordCheckQuery(
			db,
			data.Email,
			data.Password,
		)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			return
		}
		// Success (create session)
		session := sessions.Default(c)
		session.Set("user", data.Email)
		session.Save()
		c.JSON(http.StatusOK, gin.H{
			"message": "Logged in successfully!",
		})
	}
}

func main() {

	fileName := "./mydata.db"
	db, err := sql.Open("sqlite3", fileName)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	DatabasePerformanceOptimisatioins(db)
	_, err = db.Exec(CreateTableQuery)
	if err != nil {
		panic(err)
	}

	router := gin.Default()

	store := cookie.NewStore([]byte("super-secret-key"))
	router.Use(sessions.Sessions("MySession", store))

	router.GET("/ping", pingImpl)
	router.POST("/createAccount", createAccountImplGen(db))
	router.POST("/login", loginImplGen(db))

	router.Run(":8080")
}
