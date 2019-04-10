package main

import (
	"github.com/rackspace/gophercloud/rackspace"
	"github.com/rackspace/gophercloud/openstack/objectstorage/v1/objects"
	"fmt"
	"os"
	"github.com/rackspace/gophercloud"
	"os/exec"
	"io/ioutil"
	"flag"
)

func init() {
}


func main() {
	var configFile = flag.String("config", "slug-config.json", "Config file location" )
	flag.Parse()
	if *configFile == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	conf := LoadConfig(*configFile)

	var auth_object = gophercloud.AuthOptions{
		Username: conf.Username,
		Password: conf.Password,
	}

	// Auth to cloud
	var Provider, Auth_err = rackspace.AuthenticatedClient(auth_object)
	if Auth_err != nil {
		fmt.Println("An error occured in authentication")
		fmt.Println(Auth_err.Error())
		os.Exit(1)
	}

	// Instance of object storage
	serviceClient, err := rackspace.NewObjectStorageV1(Provider, gophercloud.EndpointOpts{
		Region: conf.Region,
	})
	if err != nil {
		fmt.Println("Cannot create object storage")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// Download object from cloud
	result := objects.Download(serviceClient, conf.ContainerName, conf.ObjectName, nil)

	content, err := result.ExtractContent()
	if err != nil {
		fmt.Println("Error occured while extracting download from cloud")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println("Retrieved the following file: ")
	fmt.Println(result)

	error := ioutil.WriteFile(conf.SaveLocation, []byte(content), 0644)
	if error != nil {
		fmt.Println("Unable to store downloaded file")
		fmt.Println(error.Error())
		os.Exit(1)
	}

	// Remove old extract location and contents then create a new directory
	os.RemoveAll(conf.ExtractLocation)
	os.MkdirAll(conf.ExtractLocation, 0744)

	// Using call to shell command is currently required since there is an added label in the slug
	// That has version from git in the tarball. We Should look at removing this.


	cmdName := "tar"
	cmdArgs := []string{"-xvjf", conf.SaveLocation, "-C", conf.ExtractLocation}

	cmd := exec.Command(cmdName, cmdArgs...)
	cmderr := cmd.Run()
	if cmderr != nil {
		fmt.Printf("Error untarring file %s to location %s", conf.SaveLocation, conf.ExtractLocation)
		fmt.Println(err.Error())
		os.Exit(1)
	} else {
		fmt.Printf("Completed untarring %s to location %s", conf.SaveLocation, conf.ExtractLocation)
	}


}



