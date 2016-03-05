package opst

import (
	"fmt"

	"github.com/rackspace/gophercloud"
	"github.com/rackspace/gophercloud/openstack/compute/v2/flavors"
	"github.com/rackspace/gophercloud/pagination"
)

func GetFlavors() ([]flavors.Flavor, error) {
	client, err := GetComputeClient()
	if err != nil {
		return nil, err
	}

	result, err := getFlavorList(client)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func GetFlavorDetail(id string) (*flavors.Flavor, error) {
	client, err := GetComputeClient()
	if err != nil {
		return nil, err
	}

	result, err := getFlavorDetail(client, id)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func getFlavorDetail(client *gophercloud.ServiceClient, flavor_id string) (*flavors.Flavor, error) {
	flavor, err := flavors.Get(client, flavor_id).Extract()
	if err != nil {
		return nil, err
	}

	return flavor, nil
}

func getFlavorList(client *gophercloud.ServiceClient) ([]flavors.Flavor, error) {
	opts := flavors.ListOpts{ChangesSince: "2014-01-01T01:02:03Z", MinRAM: 4}

	// Retrieve a pager (i.e. a paginated collection)
	pager := flavors.ListDetail(client, opts)

	var result2 []flavors.Flavor

	// Define an anonymous function to be executed on each page's iteration
	err := pager.EachPage(func(page pagination.Page) (bool, error) {
		flavorList, err := flavors.ExtractFlavors(page)

		if err != nil {
			fmt.Println("error ExtractFlavors")
		}

		result2 = append(result2, flavorList...)
		/*
			for _, f := range flavorList {
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
