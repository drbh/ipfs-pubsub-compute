# ipfs-pubsub-compute

## Overview

This is a tiny IPFS PubSub application built in Golang and is less than 130 lines of code. It allows you to write `Python 3.7` code in the browser send the code and the input to IPFS pubsub, which another listening device in the channel then picks up and executes using AWS's Lambda Docker container.

<p align="center">
  <img width="481" height="550" src="https://github.com/drbh/ipfs-pubsub-compute/blob/master/example/demo.gif">
</p>

 
## Setup

Please install the dependencies
```
docker - version 18.09.1
go     - go1.11.5 
ipfs   - version 0.4.18
```

Also you'll need to fetch the AWS Lambda Docker at, https://github.com/lambci/docker-lambda

To get the image just run:
```

docker pull lambci/lambda:python3.7 

```

## Running

First we need to make sure we can connect to a IPFS gateway that has PubSub enabled. The app is looking for a gateway at `http://localhost:5001` so please use `go-ipfs` to start a local node. check out https://github.com/ipfs/go-ipfs to get IPFS setup on your computer.

Now start the deamon with the pubsub flag enabled.
```
ipfs daemon --enable-pubsub-experiment
```

Now, we can start out app that uses IPFS's new gen protocols.

```
git clone https://github.com/drbh/ipfs-pubsub-compute.git
cd ipfs-pubsub-compute
```

Running precompiled binary (on OSX)
```
./ipfs-pubsub-compute 
```

Run `go` source files files, you'll need to `go get` a few libraries
```
go run execute.go server.go 
```


