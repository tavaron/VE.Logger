package VE_Direct

import (
	"strconv"
	"strings"
)

type history struct {
	DeepestDischarge          int `name:"DeepestDischarge"` // mAh
	LastDischarge             int `name:"LastDischarge"`    // mAh
	AverageDischarge          int `name:"AverageDischarge"` // mAh
	NumberOfCycles            int `name:"NumberOfCycles"`
	NumberOfFullDischarges    int `name:"NumberOfFullDischarges"`
	TotalAmpHoursDrawn        int `name:"TotalAmpHoursDrawn"`    // mAh
	TotalEnergyDischarged     int `name:"TotalEnergyDischarged"` // Wh
	TotalEnergyCharged        int `name:"TotalEnergyCharged"`    // Wh
	MainVoltageMinimum        int `name:"MainVoltageMinimum"`    // mV
	MainVoltageMaximum        int `name:"MainVoltageMaximum"`    // mV
	TimeSinceFullCharge       int `name:"TimeSinceFullCharge"`   // Seconds
	MainVoltageLowAlarmCount  int `name:"MainVoltageLowAlarmCount"`
	MainVoltageHighAlarmCount int `name:"MainVoltageHighAlarmCount"`
}

func ReadHistory(message string) (*history, error) {
	response := new(history)
	lines := strings.Split(message, "\n")

	for _, element := range lines {
		elements := strings.Split(element, "\t")
		if len(elements) == 2 {
			var err error = nil
			switch elements[0] {
			case "H1":
				response.DeepestDischarge, err = strconv.Atoi(elements[1])
				if err != nil {
					return nil, err
				}
				response.DeepestDischarge *= -1

			case "H2":
				response.LastDischarge, err = strconv.Atoi(elements[1])
				if err != nil {
					return nil, err
				}
				response.LastDischarge *= -1

			case "H3":
				response.AverageDischarge, err = strconv.Atoi(elements[1])
				if err != nil {
					return nil, err
				}
				response.AverageDischarge *= -1

			case "H4":
				response.NumberOfCycles, err = strconv.Atoi(elements[1])
				if err != nil {
					return nil, err
				}

			case "H5":
				response.NumberOfFullDischarges, err = strconv.Atoi(elements[1])
				if err != nil {
					return nil, err
				}

			case "H6":
				response.TotalAmpHoursDrawn, err = strconv.Atoi(elements[1])
				if err != nil {
					return nil, err
				}
				response.TotalAmpHoursDrawn *= -1

			case "H7":
				response.MainVoltageMinimum, err = strconv.Atoi(elements[1])
				if err != nil {
					return nil, err
				}

			case "H8":
				response.MainVoltageMaximum, err = strconv.Atoi(elements[1])
				if err != nil {
					return nil, err
				}

			case "H9":
				response.TimeSinceFullCharge, err = strconv.Atoi(elements[1])
				if err != nil {
					return nil, err
				}

			case "H11":
				response.MainVoltageLowAlarmCount, err = strconv.Atoi(elements[1])
				if err != nil {
					return nil, err
				}

			case "H12":
				response.MainVoltageHighAlarmCount, err = strconv.Atoi(elements[1])
				if err != nil {
					return nil, err
				}

			case "H17":
				response.TotalEnergyDischarged, err = strconv.Atoi(elements[1])
				if err != nil {
					return nil, err
				}
				response.TotalEnergyDischarged *= 10

			case "H18":
				response.TotalEnergyCharged, err = strconv.Atoi(elements[1])
				if err != nil {
					return nil, err
				}
				response.TotalEnergyCharged *= 10
			}
		}
	}
	return response, nil
}
