package main

import (
	"time"
	//	"log"
	//	"fmt"

	"crypto/tls"
	"log"

	"github.com/lihuacat/surgemq/service"
	"github.com/surge/glog"
	"github.com/surgemq/message"
)

func main() {
	// Instantiates a new Client
	glog.CopyStandardLogTo("INFO")
	c := &service.Client{}

	// Creates a new MQTT CONNECT message and sets the proper parameters
	msg := message.NewConnectMessage()
	msg.SetWillQos(1)
	msg.SetVersion(4)
	msg.SetCleanSession(true)
	msg.SetClientId([]byte("surgemq"))
	msg.SetKeepAlive(10)
	msg.SetWillTopic([]byte("will"))
	msg.SetWillMessage([]byte("send me home"))
	msg.SetUsername([]byte("surgemq"))
	msg.SetPassword([]byte("verysecret"))

	conf := &tls.Config{
		InsecureSkipVerify: true,
	}
	// Connects to the remote server at 127.0.0.1 port 1883
	err := c.TLSConnect("tcp://192.168.31.122:8883", msg, conf)
	if err != nil {
		glog.Errorln(err)
		return
	}
	// Creates a new SUBSCRIBE message to subscribe to topic "abc"
	submsg := message.NewSubscribeMessage()
	err = submsg.AddTopic([]byte("abc"), 0)
	if err != nil {
		log.Println(err)
		return
	}

	// Subscribes to the topic by sending the message. The first nil in the function
	// call is a OnCompleteFunc that should handle the SUBACK message from the server.
	// Nil means we are ignoring the SUBACK messages. The second nil should be a
	// OnPublishFunc that handles any messages send to the client because of this
	// subscription. Nil means we are ignoring any PUBLISH messages for this topic.
	err = c.Subscribe(submsg, nil, onPublish)
	if err != nil {
		log.Println(err)
		return
	}

	// Creates a new PUBLISH message with the appropriate contents for publishing
	pubmsg := message.NewPublishMessage()
	pubmsg.SetTopic([]byte("abc"))
	//	pubmsg.SetPayload(make([]byte, 1024))
	pubmsg.SetPayload([]byte("hahaha"))
	pubmsg.SetQoS(0)

	// Publishes to the server by sending the message
	for {
		err = c.Publish(pubmsg, nil)
		if err != nil {
			log.Println(err)
			return
		}
		time.Sleep(1 * time.Second)
	}

	// Disconnects from the server
	//	c.Disconnect()
	select {}
}

func onPublish(msg *message.PublishMessage) error {
	log.Println(string(msg.Payload()))
	return nil
}
