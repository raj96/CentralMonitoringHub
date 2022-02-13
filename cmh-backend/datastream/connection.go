package datastream

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func SpawnStreamListener(endpoint string) {
	if SourceMap[endpoint] == nil {
		return
	}
	fmt.Println("Subscribing to " + endpoint)
	MqttClient.Subscribe(endpoint, 2, func(c mqtt.Client, m mqtt.Message) {
		var jsonData map[string]interface{}
		var dataForDB = make(map[string]interface{})

		err := json.Unmarshal(m.Payload(), &jsonData)
		if err != nil {
			log.Println("Not a valid JSON, dropping data")
			return
		}

		if jsonData["machineId"] == nil || !isString(jsonData["machineId"]) {
			log.Println("No Machined ID or MachineID format error, dropping data")
			return
		}

		for key := range SourceMap[endpoint] {
			if jsonData[key] == nil {
				log.Println("Dropped message, due to missing field: " + key)
				return
			}
			dataForDB[key] = jsonData[key]
		}

		for key, val := range dataForDB {
			if SourceMap[endpoint][key] == "number" && !isNumeric(val) {
				log.Println("Dropped message, due to type mismatch of field: " + key)
				return
			} else if SourceMap[endpoint][key] == "state" && !isState(val) {
				log.Println("Dropped message, due to type mismatch of field: " + key)
				return
			}
		}

		dataForDB["timestamp"] = time.Now().UTC()
		//TODO: userId identification
		dataForDB["metadata"] = map[string]string{"machineId": jsonData["machineId"].(string), "userId": "test"}
		// if jsonData["timestamp"] == nil {
		// 	dataForDB["timestamp"] = time.Now()
		// } else {
		// 	dataForDB["timestamp"] = jsonData["timestamp"]
		// }

		go addDataToTimeSeriesDB(jsonData["machineId"].(string), dataForDB)
	})
}

func DestroyStreamListener(endpoint string) {
	MqttClient.Unsubscribe(endpoint)
}

func isNumeric(s interface{}) bool {
	typeof := reflect.TypeOf(s).Kind()
	return (typeof == reflect.Int || typeof == reflect.Float64)

}

func isState(s interface{}) bool {
	ss := fmt.Sprintf("%v", s)
	return (ss == "0" || ss == "1")
}

func isString(s interface{}) bool {
	return reflect.TypeOf(s).Kind() == reflect.String
}
