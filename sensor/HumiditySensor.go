package sensor

import (
	"github.com/brutella/hc/accessory"
	"github.com/brutella/hc/service"
)

type HumiditySensor struct {
	*accessory.Accessory

	HumiditySensor *service.HumiditySensor
}

// NewTemperatureSensor returns a Thermometer which implements model.Thermometer.
func NewHumiditySensor(info accessory.Info, relativeHumidity, min, max, steps float64) *HumiditySensor {
	acc := HumiditySensor{}
	acc.Accessory = accessory.New(info, accessory.TypeSensor)
	acc.HumiditySensor = service.NewHumiditySensor()
	acc.HumiditySensor.CurrentRelativeHumidity.SetValue(relativeHumidity)
	acc.HumiditySensor.CurrentRelativeHumidity.SetMinValue(min)
	acc.HumiditySensor.CurrentRelativeHumidity.SetMaxValue(max)
	acc.HumiditySensor.CurrentRelativeHumidity.SetStepValue(steps)

	acc.AddService(acc.HumiditySensor.Service)

	return &acc
}
