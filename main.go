package main

import (
	"net/http"
	"gwf"
)

func main() {
	r := gwf.New()
	r.GET("/", func(c *gwf.Context) {
		c.HTML(http.StatusOK, "<h1>Hello gwf</h1>")
	})
	r.GET("/hello", func(c *gwf.Context) {
		// except /hello?name=aasd
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.POST("/login", func(c *gwf.Context) {
		c.JSON(http.StatusOK, gwf.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	r.Run(":9999")
}