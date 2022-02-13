package datastream

import (
	"cmh-backend/cmhtypes"
	"cmh-backend/model"
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var SourceMap map[string]cmhtypes.Statistics
var MqttClient mqtt.Client

func Init() {
	LoadSourceTypesInMemory()
	MqttClient.Subscribe("/system", 2, HandleSystemCmds)
}

func LoadSourceTypesInMemory() {
	SourceMap = model.FetchSourceAndTypes()
	fmt.Println(SourceMap)
	for key := range SourceMap {
		SpawnStreamListener(key)
	}
	fmt.Println("SourceMap loading done")
}

func HandleSystemCmds(client mqtt.Client, msg mqtt.Message) {
	fmt.Println(string(msg.Payload()))
}

func ConnectToMqttBroker(host string, port int16) {
	connectionString := fmt.Sprintf("tcp://%s:%d", host, port)
	opts := mqtt.NewClientOptions().AddBroker(connectionString)
	opts = opts.SetOrderMatters(true)

	MqttClient = mqtt.NewClient(opts)
	if token := MqttClient.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	Init()
}

func DisconnectFromMqttBroker() {
	if MqttClient.IsConnected() {
		MqttClient.Disconnect(250)
	}
}
