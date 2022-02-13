package datastream

import (
	"cmh-backend/model"
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func HandleEndpointData(client mqtt.Client, msg mqtt.Message) {

}

func addDataToTimeSeriesDB(machineId string, data map[string]interface{}) {
	if model.InsertDataToTimeSeries(machineId, data) != nil {
		if model.CreateTimeSeries(machineId) {
			addDataToTimeSeriesDB(machineId, data)
		} else {
			log.Println("Could not log data for " + machineId)
		}
	}
}
