# go-mqtt-farmer
Example how to use Mqtt with a IoT. Server and Device in golang. Farmer example

[Requirements](https://docs.google.com/document/d/1Pv7xQTRpJOWQDgW-81lXTJhmt2Jr_bB-aaBoB-gFyl4/edit?usp=sharing)


### Setup 
    go get ./... 
    docker-compose up -d
    export MQTT_URL=mqtt://ANY_USER:ANY_PASS@127.0.0.1:1884
    go run service/main.go