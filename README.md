# slug-sync-go


### Requirements when running:  

## Building slug-sync  
slug-sync program uses the gophercloud package from:  
https://github.com/rackspace/gophercloud

in order to install and build you will have to download it using "go get" then build for linux
```
go get github.com/rackspace/gophercloud
```

Build for linux from other operating system:  
```
GOOS=linux go build
```

## Test in docker  
There is a docker file that allows you to test in the container. prior to building and running the container you will have to 
build the binary for linux using the commands above in the build section

```
docker build -t slug-sync .
docker run -it slug-sync
```

From inside the container you can run slug sync as long as you add the config file
in the format listed bellow.



## Running slug-sync  
To run slug-sync use the following command:  
```bash
slug-sync --config /etc/slug-sync/slug-config.json

```

If the --config <filename> flag is left off then the program will try to open slug-config.json in the same directory


The config file should look like this:  
```json

{
  "region": "<Three letter region code>",
  "container_name": "<Name of container in cloud files>",
  "object_name": "<filename on cloudfiles to download>",
  "username": "<username for RS cloud>",
  "password": "<password>",
  "save_location": "<location to save including file name>",
  "extract_location": "<what folder to extract tar/bz2 file to>"
}

```



