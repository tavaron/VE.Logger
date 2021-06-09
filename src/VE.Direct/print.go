package VE_Direct

import (
	"fmt"
)

func (data shunt) Print() {
	fmt.Printf("\nMessage from %02d:%02d:%02d.%03d\n", data.Time.Hour(), data.Time.Minute(), data.Time.Second(), data.Time.Nanosecond()/1000000)
	fmt.Printf("ProductID\t%s\n", data.ProductID)
	fmt.Printf("MainVoltage\t%d mV\n", data.MainVoltage)
	if data.AuxVoltage > 0 {
		fmt.Printf("AuxVoltage\t%d mV\n", data.AuxVoltage)
	}
	fmt.Printf("Current\t\t%d mAh\n", data.Current)
	fmt.Printf("Power\t\t%d W\n", data.Power)
	fmt.Printf("Consumed\t%d mAh\n", data.ConsumedAmpHours)
	fmt.Printf("SoC\t\t%5.1f %%\n", float64(data.StateOfCharge)/10.0)
	fmt.Printf("Time left\t%d minutes\n", data.TimeToGo)
	fmt.Printf("Temperature\t%d °C\n", data.Temperature)
	if data.Alarm != "" {
		fmt.Printf("Alarm: %s\n", data.Alarm)
	}
}

func (data *history) Print() {
	fmt.Println("\nHistorical data")
	fmt.Printf("Last Discharge\t%d mAh\n", data.LastDischarge)
	fmt.Printf("Deepest Discharge\t%d mAh\n", data.DeepestDischarge)
	fmt.Printf("Average Discharge\t%d mAh\n", data.AverageDischarge)
	fmt.Printf("Number of Cycles\t%d\n", data.NumberOfCycles)
	fmt.Printf("Number of full Discharges\t%d\n", data.NumberOfFullDischarges)
	fmt.Printf("Total mAh drawn\t%d mAh\n", data.TotalAmpHoursDrawn)
	fmt.Printf("Total Energy Discharged\t%.2f kWh\n", float64(data.TotalEnergyDischarged)/100)
	fmt.Printf("Total Energy Charged\t%.2f kWh\n", float64(data.TotalEnergyCharged)/100)
	fmt.Printf("Main Voltage Minimum\t%d mV\n", data.MainVoltageMinimum)
	fmt.Printf("Main Voltage Maximum\t%d mV\n", data.MainVoltageMaximum)
	fmt.Printf("Time since full Charge\t%d seconds\n", data.TimeSinceFullCharge)
	fmt.Printf("Number of Low Main Voltage Alarms\t%d\n", data.MainVoltageLowAlarmCount)
	fmt.Printf("Number of Low Main Voltage Alarms\t%d\n", data.MainVoltageHighAlarmCount)
}
