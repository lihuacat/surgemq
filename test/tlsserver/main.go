package main

import (
	"crypto/tls"
	//	"fmt"

	"log"

	"github.com/lihuacat/surgemq/service"
	"github.com/surge/glog"
)

func main() {
	glog.CopyStandardLogTo("INFO")
	cer, err := tls.LoadX509KeyPair("cert.pem", "key.pem")
	if err != nil {
		log.Println(err)
		return
	}
	config := &tls.Config{Certificates: []tls.Certificate{cer}}
	server := service.Server{
		KeepAlive:      60,
		ConnectTimeout: 10,
		TimeoutRetries: 2,
	}

	err = server.TLSListenAndServe("tcp://:8883", config)
	if err != nil {
		log.Fatal(err)
		return
	}
}
