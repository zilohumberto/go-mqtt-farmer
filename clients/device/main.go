package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	mqtttools "github.com/zilohumberto/go-mqtt-farmer/mqtt"
)

var (
	sensorF     = flag.String("iot_sensor", "iot/%s/sensor/+/", "A IoT generate a sensor information or execute a action")
	mqttURL     = os.Getenv("MQTT_URL") //MQTT_URL:mqtt://<user>:<pass>@<server>.cloudmqtt.com:<port>
	topicFormat = "device/%s/%s/%s/"    //device/uuid/category/action
)

// Publish a payload into topic with Client credentials
func Publish(client mqtt.Client, topic string, payload string) {
	fmt.Println("Publishing into ", topic)
	client.Publish(topic, 0, false, payload)
}

// HearIoT is Subscriber to sensor
func HearIoT(client mqtt.Client) {
	sensor := fmt.Sprintf(*sensorF, "weareinlive")
	fmt.Println("subscribing into ", sensor)
	client.Subscribe(sensor, 0, func(client mqtt.Client, msg mqtt.Message) {
		fmt.Printf("sensors [%s] %s\n", msg.Topic(), string(msg.Payload()))
	})
}

func main() {
	flag.Parse()
	fmt.Println("Starting device simulation!")

	uri := mqtttools.GetURL(mqttURL)
	opts := mqtttools.CreateClientOptions("device", uri)
	client := mqtttools.Connect(opts)

	//go HearIoT(client)

	time.Sleep(1)
	fmt.Println("Enter a key to produce messages")
	fmt.Scanf("%s\n")

	Publish(client, fmt.Sprintf(topicFormat, "weareinlive", "ondemand", "window"), "on")
	Publish(client, fmt.Sprintf(topicFormat, "weareinlive", "ondemand", "window"), "off")
	//go Publish(client, fmt.Sprintf(topicFormat, "weareinlive", "ondemand", "light"), "0")
	//go Publish(client, fmt.Sprintf(topicFormat, "weareinlive", "ondemand", "light"), "50")
	//go Publish(client, fmt.Sprintf(topicFormat, "weareinlive", "ondemand", "light"), "100")
	//go Publish(client, fmt.Sprintf(topicFormat, "weareinlive", "ondemand", "weather"), "15")
	//go Publish(client, fmt.Sprintf(topicFormat, "weareinlive", "ondemand", "weather"), "20")
	//go Publish(client, fmt.Sprintf(topicFormat, "weareinlive", "ondemand", "weather"), "30")
	//go Publish(client, fmt.Sprintf(topicFormat, "weareinlive", "ondemand", "water"), "on")
	//go Publish(client, fmt.Sprintf(topicFormat, "weareinlive", "ondemand", "water"), "off")

	Publish(client, fmt.Sprintf(topicFormat, "weareinlive", "query", "window"), "on")
	Publish(client, fmt.Sprintf(topicFormat, "weareinlive", "query", "window"), "off")
	//go Publish(client, fmt.Sprintf(topicFormat, "weareinlive", "query", "light"), "0")
	//go Publish(client, fmt.Sprintf(topicFormat, "weareinlive", "query", "light"), "50")
	//go Publish(client, fmt.Sprintf(topicFormat, "weareinlive", "query", "light"), "100")
	//go Publish(client, fmt.Sprintf(topicFormat, "weareinlive", "query", "weather"), "15")
	//go Publish(client, fmt.Sprintf(topicFormat, "weareinlive", "query", "weather"), "20")
	//go Publish(client, fmt.Sprintf(topicFormat, "weareinlive", "query", "weather"), "30")
	//go Publish(client, fmt.Sprintf(topicFormat, "weareinlive", "query", "water"), "on")
	//go Publish(client, fmt.Sprintf(topicFormat, "weareinlive", "query", "water"), "off")
	for {
		time.Sleep(1)
	}
}
