package main

import (
	"fmt"
	"github.com/rackspace/gophercloud"
	"github.com/rackspace/gophercloud/openstack"
	"github.com/rackspace/gophercloud/openstack/compute/v2/extensions/keypairs"
	"github.com/rackspace/gophercloud/openstack/compute/v2/servers"
	"log"
	"os"
)

func main() {
	opts, err := openstack.AuthOptionsFromEnv()
	if err != nil {
		log.Fatal("OpenStack Auth Error: " + err.Error())
	}
	provider, err := openstack.AuthenticatedClient(opts)
	client, err := openstack.NewComputeV2(provider, gophercloud.EndpointOpts{
		Region: os.Getenv("OS_REGION_NAME"),
	})

	var server_opts servers.CreateOptsBuilder
	server_opts = &servers.CreateOpts{
		Name:      "testserver",
		FlavorRef: "small",
		ImageRef:  "e45c00df-b189-414c-b82c-0917cd8a464d",
	}
	server_opts = &keypairs.CreateOptsExt{
		CreateOptsBuilder: server_opts,
		KeyName:           "nibz",
	}

	server, err := servers.Create(client, server_opts).Extract()
	if err != nil {
		fmt.Println("Unable to create server: %s", err)
	}
	fmt.Println("Server ID: %s", server.ID)
}
