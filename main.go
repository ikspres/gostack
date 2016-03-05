package main

import (
	"fmt"
	"github.com/gin-gonic/contrib/renders/multitemplate"
	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/ikspres/gostack/opst"
)

func main() {
	router := gin.Default()
	router.StaticFS("/css", http.Dir("./lumino/css"))
	router.StaticFS("/fonts", http.Dir("./lumino/fonts"))
	router.StaticFS("/js", http.Dir("./lumino/js"))

	router.HTMLRender = createMyRender()

	router.GET("/index", handle_index)
	router.GET("/charts", handle_charts)
	router.GET("/compute", handle_compute)

	remote_route := router.Group("/remote")
	{
		remote_route.GET("/flavors", handle_remote_flavors)
		remote_route.GET("/images", handle_remote_images)
		remote_route.GET("/servers", handle_remote_servers)
	}

	router.Run(":3001")
}

func createMyRender() multitemplate.Render {
	r := multitemplate.New()
	r.AddFromFiles("index", "lumino/base.tmpl", "lumino/menu.tmpl", "lumino/index.tmpl")
	r.AddFromFiles("charts", "lumino/base.tmpl", "lumino/menu.tmpl", "lumino/charts.tmpl")
	r.AddFromFiles("compute", "lumino/base.tmpl", "lumino/menu.tmpl", "lumino/compute.tmpl")

	return r
}

func handle_remote_flavors(c *gin.Context) {
	result, err := opst.GetFlavors()
	if err != nil {
		fmt.Println("error 2: " + err.Error())
	}
	c.JSON(http.StatusOK, result)
}
func handle_remote_images(c *gin.Context) {
	result, err := opst.GetImages()
	if err != nil {
		fmt.Println("error 2: " + err.Error())
	}
	c.JSON(http.StatusOK, result)
}
func handle_remote_servers(c *gin.Context) {
	result, err := opst.GetServers()
	if err != nil {
		fmt.Println("error 2: " + err.Error())
	}
	c.JSON(http.StatusOK, result)
}

func handle_index(c *gin.Context) {
	flavors, err := opst.GetFlavors()
	if err != nil {
		fmt.Println("error 2: " + err.Error())
	}

	images, err := opst.GetImages()
	if err != nil {
		fmt.Println("error 2: " + err.Error())
	}

	servers, err := opst.GetServers()
	if err != nil {
		fmt.Println("error 2: " + err.Error())
	}

	obj := gin.H{
		"flavors": flavors,
		"images":  images,
		"servers": servers,
		"title":   "INDEX",
	}
	c.HTML(http.StatusOK, "index", obj)
}

func handle_compute(c *gin.Context) {
	obj := gin.H{
		"title": "Compute",
	}
	c.HTML(http.StatusOK, "compute", obj)
}

func handle_charts(c *gin.Context) {
	obj := gin.H{"title": "charts"}
	c.HTML(http.StatusOK, "charts", obj)
}
