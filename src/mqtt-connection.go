package main

import (
	"crypto/tls"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
	"net/url"
	"time"
)

var mqttState = struct {
	broker    string
	connected bool
	lastError error
	client    mqtt.Client
	topic     string
}{
	broker:    "",
	connected: false,
	lastError: nil,
	client:    nil,
	topic:     "",
}

var mqttMessageHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	log.Printf("MQTT received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}

var mqttConnectionAttemptHandler mqtt.ConnectionAttemptHandler = func(broker *url.URL, tlsCfg *tls.Config) *tls.Config {
	log.Println(time.Now().String(), "\tMQTT connecting to ", broker.String())
	return tlsCfg
}

var mqttOnConnectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	mqttState.connected = true
	log.Println(time.Now().String(), "\tMQTT connected to ", mqttState.broker)
}

var mqttConnectionLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	mqttState.connected = false
	mqttState.lastError = err
	log.Printf("MQTT connection lost: %v\n", err)
}

var mqttReconnectHandler mqtt.ReconnectHandler = func(client mqtt.Client, options *mqtt.ClientOptions) {
	mqttState.connected = false
	log.Printf("MQTT reconnecting to %s", options.Servers[0])
}

func connectMQTT(broker string, port int, useTLS bool, username string, password string, timeout time.Duration, retryInterval time.Duration) {
	log.Println("Attempting MQTT connection")
	opts := mqtt.NewClientOptions()
	mqttState.broker = fmt.Sprintf("tcp://%s:%d", broker, port)
	opts.AddBroker(mqttState.broker)
	opts.SetClientID("ve.direct_logger")
	opts.SetUsername(username)
	opts.SetPassword(password)
	opts.SetDefaultPublishHandler(mqttMessageHandler)
	opts.SetConnectTimeout(timeout)
	opts.SetWriteTimeout(timeout)
	opts.SetPingTimeout(timeout)
	opts.OnConnectAttempt = mqttConnectionAttemptHandler
	opts.OnConnect = mqttOnConnectHandler
	opts.OnConnectionLost = mqttConnectionLostHandler
	if useTLS {
		opts.SetTLSConfig(&tls.Config{
			InsecureSkipVerify: true,
		})
	}
	if retryInterval > 0 {
		opts.AutoReconnect = true
		opts.ConnectRetry = true
		opts.ConnectRetryInterval = retryInterval
		opts.OnReconnecting = mqttReconnectHandler
	}

	mqttState.client = mqtt.NewClient(opts)

	if token := mqttState.client.Connect(); token.Wait() && token.Error() != nil {
		mqttState.connected = false
		mqttState.lastError = token.Error()
		log.Fatalln("MQTT connection failed: ", token.Error())
	}
}
