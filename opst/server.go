package opst

import (
	"fmt"

	"github.com/rackspace/gophercloud"
	"github.com/rackspace/gophercloud/openstack/compute/v2/servers"
	"github.com/rackspace/gophercloud/pagination"
)

func GetServers() ([]servers.Server, error) {
	client, err := GetComputeClient()
	if err != nil {
		return nil, err
	}

	result, err := getServerList(client)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func GetServerDetail(id string) (*servers.Server, error) {
	client, err := GetComputeClient()
	if err != nil {
		return nil, err
	}

	result, err := getServerDetail(client, id)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func getServerDetail(client *gophercloud.ServiceClient, server_id string) (*servers.Server, error) {
	server, err := servers.Get(client, server_id).Extract()
	if err != nil {
		return nil, err
	}

	return server, nil
}

func getServerList(client *gophercloud.ServiceClient) ([]servers.Server, error) {
	opts := servers.ListOpts{}

	// Retrieve a pager (i.e. a paginated collection)
	pager := servers.List(client, opts)

	var result2 []servers.Server

	// Define an anonymous function to be executed on each page's iteration
	err := pager.EachPage(func(page pagination.Page) (bool, error) {
		serverList, err := servers.ExtractServers(page)

		if err != nil {
			fmt.Println("error ExtractServers")
		}

		result2 = append(result2, serverList...)
		/*
			for _, f := range serverList {
				result = append(result, f.Name+" "+f.ID)
			}
		*/

		return true, nil
	})

	if err != nil {
		return nil, err
	}

	return result2, nil
}
