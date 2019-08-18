package sensor

import (
	"github.com/brutella/hc/accessory"
	"github.com/brutella/hc/service"
)

type Thermometer struct {
	*accessory.Accessory

	TempSensor *service.TemperatureSensor
}

// NewTemperatureSensor returns a Thermometer which implements model.Thermometer.
func NewTemperatureSensor(info accessory.Info, temp, min, max, steps float64) *Thermometer {
	acc := Thermometer{}
	acc.Accessory = accessory.New(info, accessory.TypeSensor)
	acc.TempSensor = service.NewTemperatureSensor()
	acc.TempSensor.CurrentTemperature.SetValue(temp)
	acc.TempSensor.CurrentTemperature.SetMinValue(min)
	acc.TempSensor.CurrentTemperature.SetMaxValue(max)
	acc.TempSensor.CurrentTemperature.SetStepValue(steps)

	acc.AddService(acc.TempSensor.Service)

	return &acc
}
