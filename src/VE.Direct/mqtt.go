package VE_Direct

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"reflect"
	"strconv"
	"time"
)

func (data *shunt) MQTT(client mqtt.Client, topic string) error {

	t := reflect.TypeOf(*data)
	v := reflect.ValueOf(*data)
	for i := 0; i < t.NumField(); i++ {
		pubTopic := topic + "/" + t.Field(i).Tag.Get("name")
		var value string

		switch val := v.Field(i).Interface().(type) {
		case int:
			value = strconv.Itoa(val)
		case float64:
			value = fmt.Sprintf("%f", val)
		case string:
			value = val
		case time.Time:
			value = val.String()
		}

		client.Publish(pubTopic, 0, false, value)
	}
	return nil
}

func (data *history) MQTT(client mqtt.Client, topic string) error {

	t := reflect.TypeOf(*data)
	v := reflect.ValueOf(*data)
	for i := 0; i < t.NumField(); i++ {
		pubTopic := topic + "/" + t.Field(i).Tag.Get("name")
		client.Publish(pubTopic, 0, false, strconv.Itoa(int(v.Field(i).Int())))
	}
	return nil
}
