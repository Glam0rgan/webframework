package main

import (
	"net/http"

	"gwf"
)

func main() {
	r := gwf.New()
	r.GET("/index", func(c *gwf.Context) {
		c.HTML(http.StatusOK, "<h1>Index Page</h1>")
	})

	v1 := r.Group("/v1")
	{
		v1.GET("/", func(c *gwf.Context) {
			c.HTML(http.StatusOK, "<h1>Hello Gwf</h1>")
		})

		v1.GET("/hello", func(c *gwf.Context) {
			// expect /hello?name=gwf
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
		})
	}

	v2 := r.Group("/v2")
	{
		v2.GET("/hello/:name", func(c *gwf.Context) {
			// expect /hello/gwf
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
		v2.POST("/login", func(c *gwf.Context) {
			c.JSON(http.StatusOK, gwf.H{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})
		})
	}

	r.Run(":9999")
}
