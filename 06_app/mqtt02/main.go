// main.go
// 创建客户端1,订阅主题为/topic/led的信息

package main

import (
	"fmt"
	"os"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

var f MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	fmt.Println("Topic: ", msg.Topic())
	fmt.Println("Msg: ", string(msg.Payload()))
	ctx := string(msg.Payload())
	if ctx == "close" {
		exit = true
	}
}

var exit = false

func main() {
	opts := MQTT.NewClientOptions().AddBroker("tcp://192.168.1.145:1883")
	opts.SetClientID("client-01")
	opts.SetDefaultPublishHandler(f)

	// 创建client01客户端
	c := MQTT.NewClient(opts)

	// 建立连接
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error)
	}

	if token := c.Subscribe("led", 0, nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	for {
		fmt.Println("sleep 3s....")
		if exit {
			break
		}
		time.Sleep(3 * time.Second)
	}

	c.Disconnect(250)
}
