package main

import (
	VED "./VE.Direct"
	"flag"
	"github.com/tarm/serial"
	"log"
)

func main() {

	messageChan := make(chan VED.Message)

	printActive := flag.Bool("print", false, "read from serial input")

	serialActive := flag.Bool("serial", false, "print messages to stdout")
	serialDevicePath := flag.String("serial-device", "/dev/ttyUSB0", "path to VE.Direct serial device")
	serialTimeOut := flag.Int("serial-timeout", 2, "timeout for serial connection in seconds")
	serialBaudRate := flag.Int("baud", 19200, "baud rate for serial connection")

	mqttActive := flag.Bool("mqtt", false, "publish data to MQTT broker")
	mqttBroker := flag.String("mqtt-broker", "", "domain of the mqtt broker")
	mqttPort := flag.Int("mqtt-port", 1883, "port of the mqtt broker")
	mqttTLS := flag.Bool("mqtt-tls", false, "use TLS for connection to broker")
	mqttTopic := flag.String("mqtt-topic", "VE.Direct_Logger", "topic under which the data will be published")
	mqttUsername := flag.String("mqtt-user", "", "username for authentication on the broker")
	mqttPassword := flag.String("mqtt-pass", "", "password for authentication on the broker")
	mqttTimeout := flag.Int("mqtt-timeout", 5, "timeout for connection to broker")
	mqttRetry := flag.Int("mqtt-retry", 0, "retry interval for MQTT connection. disabled by default")

	influx2Active := flag.Bool("influx2", false, "publish data to InfluxDB 2")
	influx2Host := flag.String("influx2-host", "localhost", "host where InfluxDB 2 is running")
	influx2Port := flag.Int("influx2-port", 8086, "port of InfluxDB 2 on host")
	influx2Token := flag.String("influx2-token", "", "token to use for authentication")
	influx2Organisation := flag.String("influx2-org", "", "organisation for data")
	influx2Bucket := flag.String("influx2-bucket", "", "bucket for data")

	flag.Parse()

	if *serialActive {
		serialConfig := &serial.Config{
			Name:        *serialDevicePath,
			Baud:        *serialBaudRate,
			ReadTimeout: getDuration(*serialTimeOut),
			Size:        8,
			Parity:      0,
			StopBits:    1,
		}

		go func() {
			err := VED.ReadSerial(serialConfig, messageChan)
			if err != nil {
				log.Panicln("serial connection failed: ", err.Error())
			}
		}()
	} else {
		log.Println("serial interface disabled")
	}

	if *mqttActive {
		mqttState.topic = *mqttTopic
		go connectMQTT(*mqttBroker, *mqttPort, *mqttTLS, *mqttUsername, *mqttPassword, getDuration(*mqttTimeout), getDuration(*mqttRetry))
	} else {
		log.Println("mqtt interface disabled")
	}

	if *influx2Active && *influx2Bucket == "" {
		*influx2Active = false
		log.Fatalln("no bucket for InfluxDB 2 defined")

	} else if *influx2Active {
		go connectInflux2(*influx2Host, *influx2Port, *influx2Token, *influx2Organisation, *influx2Bucket)
	}

	for {
		msg := <-messageChan
		if *printActive {
			go msg.Print()
		}

		if *mqttActive && mqttState.connected && mqttState.client != nil {
			go func() {
				err := msg.MQTT(mqttState.client, mqttState.topic)
				if err != nil {
					log.Println("MQTT publishing failed: ", err.Error())
				}
			}()
		}

		if *influx2Active && influx2State.client != nil {
			go func() {
				msg.Influx2(influx2State.writeAPI)

			}()
		}
	}
}
