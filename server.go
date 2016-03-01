package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/rackspace/gophercloud"
	"github.com/rackspace/gophercloud/openstack"
	//"github.com/rackspace/gophercloud/openstack/utils"
	"github.com/rackspace/gophercloud/openstack/compute/v2/flavors"
	//	"github.com/rackspace/gophercloud/openstack/compute/v2/images"
	"github.com/rackspace/gophercloud/pagination"
)

func main() {
	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")

	router.GET("/flavors", func(c *gin.Context) {
		result, err := getFlavors()
		if err != nil {
			fmt.Println("error: " + err.Error())
		}

		c.HTML(http.StatusOK, "flavors.tmpl", gin.H{
			"flavors": result,
		})
	})

	router.Run(":3001")
}

func getFlavors() ([]string, error) {

	client, err := computeClient()
	if err != nil {
		return nil, err
	}

	result, err := flavorList(client)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func computeClient() (*gophercloud.ServiceClient, error) {
	opts, _ := openstack.AuthOptionsFromEnv()

	provider, err := openstack.AuthenticatedClient(opts)
	if err != nil {
		return nil, err
	}

	client, err := openstack.NewComputeV2(provider, gophercloud.EndpointOpts{
		Region: "RegionOne",
	})
	if err != nil {
		return nil, err
	}

	return client, nil
}

func flavorList(client *gophercloud.ServiceClient) ([]string, error) {
	opts := flavors.ListOpts{ChangesSince: "2014-01-01T01:02:03Z", MinRAM: 4}

	// Retrieve a pager (i.e. a paginated collection)
	pager := flavors.ListDetail(client, opts)

	var result []string

	// Define an anonymous function to be executed on each page's iteration
	err := pager.EachPage(func(page pagination.Page) (bool, error) {
		flavorList, err := flavors.ExtractFlavors(page)

		if err != nil {
			fmt.Println("error ExtractFlavors")
		}

		for _, f := range flavorList {
			result = append(result, f.Name)
		}

		return true, nil
	})

	if err != nil {
		return nil, err
	}

	return result, nil
}
