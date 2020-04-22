package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	mqtttools "github.com/zilohumberto/go-mqtt-farmer/mqtt"
)

var (
	uuid      = flag.String("uuid", "weareinlive", "A IoT uuid for the records")
	iotFormat = flag.String("iot_format", "iot/%s/sensor/%s/", "A IoT generate a sensor information or execute a action")
	mqttURL   = os.Getenv("MQTT_URL") //MQTT_URL:mqtt://<user>:<pass>@<server>.cloudmqtt.com:<port>
)

func splitTopic(topic string) []string {
	return strings.Split(topic, "/")
}

// Publish a payload into topic with Client credentials
func Publish(client mqtt.Client, action string, payload string) {
	topic := fmt.Sprintf(*iotFormat, *uuid, action)
	fmt.Println("Publishing into ", topic)
	client.Publish(topic, 0, false, payload)
}

// HearDevice is Subscriber to sensor
func HearDevice(client mqtt.Client) {
	topic := fmt.Sprintf("device/+/ondemand/+/")
	fmt.Println("subscribing into ", topic)
	client.Subscribe(topic, 0, func(client mqtt.Client, msg mqtt.Message) {
		fmt.Printf("HearDevice [%s] %s\n", msg.Topic(), string(msg.Payload()))
		action := splitTopic(msg.Topic())[3]
		go Publish(client, action, string(msg.Payload()))
	})
}

// HearSystem is subscribe
func HearSystem(client mqtt.Client) {
	topic := fmt.Sprintf("system/%s/ondemand/+/", *uuid)
	fmt.Println("subscribing into ", topic)
	client.Subscribe(topic, 0, func(client mqtt.Client, msg mqtt.Message) {
		fmt.Printf("HearSystem [%s] %s\n", msg.Topic(), string(msg.Payload()))
		action := splitTopic(msg.Topic())[3]
		go Publish(client, action, string(msg.Payload()))
	})
}

func main() {
	flag.Parse()
	fmt.Println("Starting device simulation!")

	uri := mqtttools.GetURL(mqttURL)
	opts := mqtttools.CreateClientOptions("iot", uri)
	client := mqtttools.Connect(opts)

	go HearDevice(client)
	go HearSystem(client)
	time.Sleep(1)
	fmt.Println("Enter a key to produce messages")
	fmt.Scanf("%s\n")
	for i := 0; i < 5; i++ {
		time.Sleep(1)
		//send status of all peripheral
		Publish(client, "temperature", "21")
		Publish(client, "humedity", "91")
		Publish(client, "wind", "11")
	}

	for {
		//keep alive!!
		time.Sleep(1)
	}
}
