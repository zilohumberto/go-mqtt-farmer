package mqtttools

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func GetURL(mqttURL string) *url.URL {
	//MQTT_URL:mqtt://<user>:<pass>@<server>.cloudmqtt.com:<port>/<topic>
	uri, err := url.Parse(os.Getenv("MQTT_URL"))
	if err != nil {
		log.Fatal(err)
	}
	return uri
}

func GetTopic(url *url.URL) string {
	topic := url.Path[1:len(url.Path)]
	if topic == "" {
		topic = "test"
	}
	return topic
}

func Connect(opts *mqtt.ClientOptions) mqtt.Client {
	client := mqtt.NewClient(opts)
	token := client.Connect()
	for !token.WaitTimeout(1 * time.Second) {
	}
	if err := token.Error(); err != nil {
		log.Fatal(err)
	}
	return client
}

func CreateClientOptions(clientID string, uri *url.URL) *mqtt.ClientOptions {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s", uri.Host))
	opts.SetUsername(uri.User.Username())
	password, _ := uri.User.Password()
	opts.SetPassword(password)
	opts.SetClientID(clientID)
	return opts
}
