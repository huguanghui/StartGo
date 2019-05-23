// main.go
// 用于创建一个mqtt的发布者

package main

import (
	"fmt"
	"os"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

var f MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	fmt.Println("Topic: ", msg.Topic())
	fmt.Println("Message: ", msg.Payload())
}

func main() {
	opts := MQTT.NewClientOptions().AddBroker("tcp://192.168.1.145:1883")
	opts.SetClientID("mqt-sub")

	c := MQTT.NewClient(opts)

	if token := c.Connect(); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	text := fmt.Sprintf("Hello World!")
	token := c.Publish("topic/led", 0, false, text)
	token.Wait()

	c.Disconnect(250)
}
