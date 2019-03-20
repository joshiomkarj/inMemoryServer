
# In Memory Server

This repository implements a simple in memory server with GET, POST, and PATCH requests.


## Purpose

This is a simple in memory server to quickly implement go based server using mux.

## Details

The server is seeded with one vm object which it returns when `/servers` endpoint is hit. By making post calls to `/servers` more vms can be added. The only field that needs to be sent in json is `vmname`.

Example request body:
```js
{
    "vmname" : "VM-Shinjuku-160-0022"
}
```


## Running

The project can be run locally by building the go executable manually or by simply running `run.sh`
The docker image exposes the server on port `8080`, the default host port is also the same. You can change the host port in `run.sh` if you need to use some other port than `8080`



## Use Cases

This is a simple in memory server that can be extended in future. It serves as a as a wuick prototyping template for building APIs and most of the resources and constructs, including the directory structure is fairly reusable.



