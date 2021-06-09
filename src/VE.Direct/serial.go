package VE_Direct

import (
	"bufio"
	"github.com/tarm/serial"
	"log"
	"strings"
)

func ReadSerial(config *serial.Config, messageChan chan Message) error {

	s, err := serial.OpenPort(config)
	if err != nil {
		return err
	}

	Scanner := bufio.NewScanner(s)

	var message string
	for Scanner.Scan() {
		if len(message) > 0 {
			message += "\n"
		}
		text := Scanner.Text()
		if strings.Contains(text, "Checksum") {
			//TODO implement checksum verification
			data, err := ParseMessage(message)
			if err != nil {
				log.Println("Malformed message received via serial ", config.Name, "\t", err.Error())
			} else {
				messageChan <- data
			}
			message = ""
		} else {
			message += text
		}
	}

	//TODO serial error handling

	return nil
}
