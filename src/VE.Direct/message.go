package VE_Direct

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/influxdata/influxdb-client-go/api"
	"strings"
)

type Message interface {
	MQTT(client mqtt.Client, topic string) error
	Influx2(writeAPI api.WriteAPI)
	Print()
}

func ParseMessage(msg string) (Message, error) {
	var response Message = nil
	var err error = nil
	if strings.Contains(msg, "BMV") {
		response, err = readShunt(msg)
	} else if !strings.Contains(msg, "PID") {
		response, err = readHistory(msg)
	}
	return response, err
}
