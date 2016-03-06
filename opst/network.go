package opst

import (
	"fmt"

	"github.com/rackspace/gophercloud"
	"github.com/rackspace/gophercloud/openstack/networking/v2/networks"
	"github.com/rackspace/gophercloud/openstack/networking/v2/subnets"
	"github.com/rackspace/gophercloud/pagination"
)

func GetNetworks() ([]networks.Network, error) {
	client, err := GetNetworkClient()
	if err != nil {
		return nil, err
	}

	result, err := getNetworkList(client)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func getNetworkList(client *gophercloud.ServiceClient) ([]networks.Network, error) {
	opts := networks.ListOpts{}
	pager := networks.List(client, opts)

	var result []networks.Network

	// Define an anonymous function to be executed on each page's iteration
	err := pager.EachPage(func(page pagination.Page) (bool, error) {
		networkList, err := networks.ExtractNetworks(page)
		if err != nil {
			fmt.Println("error ExtractNetworks")
		}

		result = append(result, networkList...)
		return true, nil
	})

	return result, err
}

func GetSubnets() ([]subnets.Subnet, error) {
	client, err := GetNetworkClient()
	if err != nil {
		return nil, err
	}

	opts := subnets.ListOpts{}
	result, err := getSubnetList(client, &opts)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func GetSubnetsOfNetwork(network_id string) ([]subnets.Subnet, error) {
	client, err := GetNetworkClient()
	if err != nil {
		return nil, err
	}

	opts := subnets.ListOpts{}
	opts.NetworkID = network_id

	result, err := getSubnetList(client, &opts)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func getSubnetList(
	client *gophercloud.ServiceClient,
	opts *subnets.ListOpts) ([]subnets.Subnet, error) {

	pager := subnets.List(client, opts)

	var result []subnets.Subnet

	// Define an anonymous function to be executed on each page's iteration
	err := pager.EachPage(func(page pagination.Page) (bool, error) {
		subnetList, err := subnets.ExtractSubnets(page)
		if err != nil {
			fmt.Println("error ExtractSubnets")
		}

		result = append(result, subnetList...)
		return true, nil
	})

	return result, err
}
