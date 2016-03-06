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
	router.StaticFS("/web", http.Dir("./www"))
	router.StaticFS("/css", http.Dir("./www/css"))
	router.StaticFS("/fonts", http.Dir("./www/fonts"))
	router.StaticFS("/js", http.Dir("./www/js"))

	router.HTMLRender = createMyRender()

	router.GET("/index", handle_index)
	router.GET("/charts", handle_charts)
	router.GET("/compute", handle_compute)
	router.GET("/network", handle_network)
	router.GET("/subnets_of_network", handle_subnets_of_network)

	remote_route := router.Group("/remote")
	{
		remote_route.GET("/flavors", handle_remote_flavors)
		remote_route.GET("/images", handle_remote_images)
		remote_route.GET("/servers", handle_remote_servers)
		remote_route.GET("/networks", handle_remote_networks)
		remote_route.GET("/subnets", handle_remote_subnets)
		remote_route.GET("/subnets_of_network", handle_remote_subnets_of_network)
	}

	router.Run(":3001")
}

func createMyRender() multitemplate.Render {
	r := multitemplate.New()
	r.AddFromFiles("index", "www/base.tmpl", "www/menu.tmpl", "www/index.tmpl")
	r.AddFromFiles("charts", "www/base.tmpl", "www/menu.tmpl", "www/charts.tmpl")
	r.AddFromFiles("compute", "www/base.tmpl", "www/menu.tmpl", "www/compute.tmpl")
	r.AddFromFiles("network", "www/base.tmpl", "www/menu.tmpl", "www/network.tmpl")
	r.AddFromFiles("subnets_of_network", "www/subnets_of_network.tmpl")

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
func handle_remote_networks(c *gin.Context) {
	result, err := opst.GetNetworks()
	if err != nil {
		fmt.Println("error 2: " + err.Error())
	}
	c.JSON(http.StatusOK, result)
}
func handle_remote_subnets(c *gin.Context) {
	result, err := opst.GetSubnets()
	if err != nil {
		fmt.Println("error 2: " + err.Error())
	}
	c.JSON(http.StatusOK, result)
}
func handle_remote_subnets_of_network(c *gin.Context) {
	network_id := "02773db5-37e4-448c-9a12-de9c52eddee2"
	result, err := opst.GetSubnetsOfNetwork(network_id)
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
		"title":   "Dashboard",
	}
	c.HTML(http.StatusOK, "index", obj)
}

func handle_compute(c *gin.Context) {
	obj := gin.H{
		"title": "Compute",
	}
	c.HTML(http.StatusOK, "compute", obj)
}

func handle_network(c *gin.Context) {
	obj := gin.H{
		"title": "Network",
	}
	c.HTML(http.StatusOK, "network", obj)
}

func handle_subnets_of_network(c *gin.Context) {
	obj := gin.H{}
	c.HTML(http.StatusOK, "subnets_of_network", obj)
}

func handle_charts(c *gin.Context) {
	obj := gin.H{"title": "charts"}
	c.HTML(http.StatusOK, "charts", obj)
}
