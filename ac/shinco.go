package ac

import (
	"fmt"
	"github.com/brutella/hc/characteristic"
)

func ShincoCodeFactory(irThermostat *IRThermostat, state int, temperature float64) *string {
	var code string
	if state == characteristic.CurrentHeatingCoolingStateOff {
		code = "off"
	} else if state != characteristic.CurrentHeatingCoolingStateHeat {
		code = fmt.Sprintf("%.0f", irThermostat.irThermostatAccessory.Thermostat.TargetTemperature.GetValue()) + "c"
	}
	return &code
}