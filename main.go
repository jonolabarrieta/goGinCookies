package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.LoadHTMLFiles("index.html")
	router.Use(CORSMiddleware())
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	router.GET("/setToken", func(c *gin.Context) {

		token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"

		cookie, err := c.Cookie("token")

		if err != nil {
			cookie = "NotSet"
			c.SetCookie("token", token, 3600, "/", "localhost", false, true)
		}

		fmt.Printf("Cookie value: %s \n", cookie)

		c.String(http.StatusOK, "cookie seteada")
	})

	router.GET("/getToken", func(c *gin.Context) {

		cookie, err := c.Cookie("token")

		if err != nil {
			cookie = "NotSet"
		}

		fmt.Printf("Cookie value: %s \n", cookie)

		c.String(http.StatusOK, cookie)
	})

	router.Run()
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
