package opst

import (
	"github.com/rackspace/gophercloud"
	"github.com/rackspace/gophercloud/openstack/compute/v2/images"
	"github.com/rackspace/gophercloud/pagination"
)

func GetImages() ([]images.Image, error) {
	client, err := GetComputeClient()
	if err != nil {
		return nil, err
	}

	result, err := getImageList(client)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func GetImageDetail(id string) (*images.Image, error) {
	client, err := GetComputeClient()
	if err != nil {
		return nil, err
	}

	result, err := getImageDetail(client, id)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func getImageDetail(client *gophercloud.ServiceClient, image_id string) (*images.Image, error) {
	image, err := images.Get(client, image_id).Extract()
	if err != nil {
		return nil, err
	}

	return image, nil
}

func getImageList(client *gophercloud.ServiceClient) ([]images.Image, error) {
	//opts := images.ListOpts{ChangesSince: "2014-01-01T01:02:03Z"}
	opts := images.ListOpts{}

	pager := images.ListDetail(client, opts)

	var result []images.Image

	// Define an anonymous function to be executed on each page's iteration
	err := pager.EachPage(func(page pagination.Page) (bool, error) {
		imageList, err := images.ExtractImages(page)

		if err != nil {
			return false, err
		}

		result = append(result, imageList...)
		return true, nil
	})

	if err != nil {
		return nil, err
	}

	return result, nil
}
