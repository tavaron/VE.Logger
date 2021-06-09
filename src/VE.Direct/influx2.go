package VE_Direct

import (
	influxdb2 "github.com/influxdata/influxdb-client-go"
	"github.com/influxdata/influxdb-client-go/api"
	"reflect"
	"time"
)

func (data *shunt) Influx2(writeAPI api.WriteAPI) {
	dataMap := make(map[string]interface{})
	t := reflect.TypeOf(*data)
	v := reflect.ValueOf(*data)
	for i := 0; i < t.NumField(); i++ {
		dataMap[t.Field(i).Tag.Get("name")] = v.Field(i).Interface()
	}

	p := influxdb2.NewPoint("shunt", map[string]string{}, dataMap, data.Time)

	writeAPI.WritePoint(p)
}

func (data *history) Influx2(writeAPI api.WriteAPI) {
	dataMap := make(map[string]interface{})
	t := reflect.TypeOf(*data)
	v := reflect.ValueOf(*data)
	for i := 0; i < t.NumField(); i++ {
		dataMap[t.Field(i).Tag.Get("name")] = v.Field(i).Interface()
	}

	p := influxdb2.NewPoint("history", map[string]string{}, dataMap, time.Now())

	writeAPI.WritePoint(p)
}
