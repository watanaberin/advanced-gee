package main

import (
	"gee"
	"net/http"
)

func main() {
	r := gee.New()

	r.GET("/hello", func(c *gee.Context) {
		c.String(http.StatusOK, "hello %s, ur at %s\n", c.Query("name"), c.Path)
	})

	r.POST("/login", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})
	r.POST("/pre/:name", func(c *gee.Context) {
		c.String(200, "get /pre/%s\n", c.Params["name"])
	})
	r.Run(":9999")
}
