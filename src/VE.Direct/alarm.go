package VE_Direct

func convertAlarmReason(ar int) string {
	var response string
	//TODO add check for several combined alarm reasons
	switch ar {
	case 0:
		response = ""
	case 1:
		response = "Low Voltage"
	case 2:
		response = "High Voltage"
	case 4:
		response = "Low State-Of-Charge"
	case 8:
		response = "Low Auxiliary Voltage"
	case 16:
		response = "High Auxiliary Voltage"
	case 32:
		response = "Low Temperature"
	case 64:
		response = "High Temperature"
	case 128:
		response = "Mid-Point Voltage"
	}

	return response
}
