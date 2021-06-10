package VE_Direct

import (
	"strconv"
	"strings"
	"time"
)

// shunt Data Type
// Tags will be used as name while publishing to MQTT or DB
type shunt struct {
	Time             time.Time `name:"LastUpdate"`
	ProductID        string    `name:"PID"`
	MainVoltage      float64   `name:"MainVoltage"`      // V
	AuxVoltage       float64   `name:"AuxVoltage"`       // V
	Current          float64   `name:"MainCurrent"`      // A
	Power            int       `name:"MainPower"`        // W
	ConsumedAmpHours float64   `name:"ConsumedAmpHours"` // Ah
	StateOfCharge    float64   `name:"SoC"`              // %
	TimeToGo         int       `name:"MinutesLeft"`      // Minutes
	Temperature      int       `name:"Temperature"`      // °C
	Alarm            string    `name:"AlarmReason"`      // empty if none, else reason
}

func readShunt(message string) (*shunt, error) {
	response := new(shunt)
	response.Time = time.Now()
	lines := strings.Split(message, "\n")

	for _, element := range lines {
		elements := strings.Split(element, "\t")
		if len(elements) == 2 {
			var err error = nil
			switch elements[0] {
			case "PID":
				response.ProductID = elements[1]

			case "V":
				response.MainVoltage, err = strconv.ParseFloat(elements[1], 64)
				if err != nil {
					return nil, err
				}
				response.MainVoltage /= 1000.0

			case "VS":
				response.AuxVoltage, err = strconv.ParseFloat(elements[1], 64)
				if err != nil {
					return nil, err
				}
				response.AuxVoltage /= 1000.0

			case "I":
				response.Current, err = strconv.ParseFloat(elements[1], 64)
				if err != nil {
					return nil, err
				}
				response.Current /= 1000.0

			case "T":
				response.Temperature, err = strconv.Atoi(elements[1])
				if err != nil {
					return nil, err
				}

			case "P":
				response.Power, err = strconv.Atoi(elements[1])
				if err != nil {
					return nil, err
				}

			case "CE":
				if elements[1] == "---" {
					response.ConsumedAmpHours = -127
				} else {
					response.ConsumedAmpHours, err = strconv.ParseFloat(elements[1], 64)
					if err != nil {
						return nil, err
					}
					response.ConsumedAmpHours *= -1
					response.ConsumedAmpHours /= 1000.0
				}

			case "SOC":
				if elements[1] == "---" {
					response.StateOfCharge = -127
				} else {
					response.StateOfCharge, err = strconv.ParseFloat(elements[1], 64)
					if err != nil {
						return nil, err
					}
					response.StateOfCharge /= 10.0
				}
			case "TTG":
				if elements[1] == "---" {
					response.TimeToGo = -127
				} else {
					response.TimeToGo, err = strconv.Atoi(elements[1])
					if err != nil {
						return nil, err
					}
				}
			case "AR":
				ar := 0
				ar, err = strconv.Atoi(elements[1])
				if err != nil {
					return nil, err
				}
				response.Alarm = convertAlarmReason(ar)
			}
		}
	}

	return response, nil
}
