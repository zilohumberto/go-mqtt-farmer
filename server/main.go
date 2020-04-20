package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	mqtttools "github.com/zilohumberto/go-mqtt-farmer/mqtt"
	"github.com/zilohumberto/go-mqtt-farmer/server/handler"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var (
	req     = flag.String("client_req", "device/+/ondemand/+/", "A client request to change any action in the IoT, server save the información requested")
	query   = flag.String("client_query", "device/+/query/+/", "A client ask to change any action in the IoT, server verify the information (decide if send or nor to IoT)")
	sensor  = flag.String("iot_sensor", "iot/+/sensor/+/", "A IoT generate a sensor information or execute a action")
	mqttURL = os.Getenv("MQTT_URL") //MQTT_URL:mqtt://<user>:<pass>@<server>.cloudmqtt.com:<port>
)

func splitTopic(topic string) []string {
	return strings.Split(topic, "/")
}

func SaveClientReq(client mqtt.Client) {
	fmt.Println("subscribing into ", *req)
	client.Subscribe(*req, 0, func(client mqtt.Client, msg mqtt.Message) {
		topics := splitTopic(msg.Topic())
		uuidDevice := topics[1]
		action := topics[3]
		fmt.Printf("req [%s] %s\n", msg.Topic(), string(msg.Payload()))
		handler.SaveInformation(uuidDevice, action, msg.Topic(), msg.Payload())
	})
}

func QueryClient(client mqtt.Client) {
	fmt.Println("subscribing into ", *query)
	client.Subscribe(*query, 0, func(client mqtt.Client, msg mqtt.Message) {
		// simulate we spent 1 seg to calculate the result
		fmt.Printf("query [%s] %s\n", msg.Topic(), string(msg.Payload()))
		//produce a message to iot!
		topics := splitTopic(msg.Topic())
		uuidDevice := topics[1]
		action := topics[3]
		handler.SaveInformation(uuidDevice, action, msg.Topic(), msg.Payload())
		produceMessage(client, uuidDevice, action, string(msg.Payload()))
	})
}

func produceMessage(client mqtt.Client, uuid string, action string, payload string) {
	topic := fmt.Sprintf("system/%s/ondemand/%s", uuid, action)
	client.Publish(topic, 0, false, payload)
	handler.SaveInformation(uuid, action, topic, []byte("+"+payload))
}

func SaveSensorInfo(client mqtt.Client) {
	fmt.Println("subscribing into ", *sensor)
	client.Subscribe(*sensor, 0, func(client mqtt.Client, msg mqtt.Message) {
		fmt.Printf("sensons [%s] %s\n", msg.Topic(), string(msg.Payload()))
		topics := splitTopic(msg.Topic())
		uuidDevice := topics[1]
		action := topics[3]
		go handler.SaveInformation(uuidDevice, action, msg.Topic(), msg.Payload())
	})
}

func main() {
	flag.Parse()
	uri := mqtttools.GetURL(mqttURL)
	opts := mqtttools.CreateClientOptions("server", uri)
	client := mqtttools.Connect(opts)

	go SaveClientReq(client)
	go QueryClient(client)
	go SaveSensorInfo(client)
	for {
		//keep alive!
		time.Sleep(1)
	}
}
