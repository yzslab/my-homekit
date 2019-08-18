package ac

import (
	"github.com/brutella/hc/accessory"
	"github.com/brutella/hc/service"
)

type IRThermostatAccessory struct {
	*accessory.Accessory

	Thermostat *service.Thermostat
}

// NewThermostat returns a Thermostat which implements model.Thermostat.
func NewIRThermostatAccessory(info accessory.Info, temp, min, max float64) *IRThermostatAccessory {
	acc := IRThermostatAccessory{}
	acc.Accessory = accessory.New(info, accessory.TypeThermostat)
	acc.Thermostat = service.NewThermostat()
	acc.Thermostat.CurrentTemperature.SetValue(temp)
	acc.Thermostat.CurrentTemperature.SetMinValue(0)
	acc.Thermostat.CurrentTemperature.SetMaxValue(100)
	acc.Thermostat.CurrentTemperature.SetStepValue(0.1)

	acc.Thermostat.TargetTemperature.SetValue(min)
	acc.Thermostat.TargetTemperature.SetMinValue(min)
	acc.Thermostat.TargetTemperature.SetMaxValue(max)
	acc.Thermostat.TargetTemperature.SetStepValue(1)

	acc.AddService(acc.Thermostat.Service)

	return &acc
}
