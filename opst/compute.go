package opst

import (
	"github.com/rackspace/gophercloud"
	"github.com/rackspace/gophercloud/openstack"
)

func GetComputeClient() (*gophercloud.ServiceClient, error) {
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
