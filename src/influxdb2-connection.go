package main

import (
	"fmt"
	influxdb2 "github.com/influxdata/influxdb-client-go"
	"github.com/influxdata/influxdb-client-go/api"
)

var influx2State = struct {
	connected bool
	server    string
	client    influxdb2.Client
	writeAPI  api.WriteAPI
}{
	connected: false,
	server:    "",
	client:    nil,
	writeAPI:  nil,
}

func connectInflux2(host string, port int, token string, org string, bucket string) {
	influx2State.server = fmt.Sprintf("http://%s:%d", host, port)
	influx2State.client = influxdb2.NewClient(influx2State.server, token)
	influx2State.writeAPI = influx2State.client.WriteAPI(org, bucket)
}
