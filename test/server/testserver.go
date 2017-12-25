package main

import (
	//	"fmt"

	"log"

	"github.com/lihuacat/surgemq/service"
	"github.com/surge/glog"
)

func main() {
	glog.CopyStandardLogTo("INFO")
	server := service.Server{
		KeepAlive:      60,
		ConnectTimeout: 10,
		TimeoutRetries: 2,
	}

	err := server.ListenAndServe("tcp://:1883")
	if err != nil {
		log.Fatal(err)
	}
}
