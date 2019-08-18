package sensor

import "github.com/brutella/hc/log"

type TemperatureHumidityReader interface {
	Read() (temperature float64, relativeHumidity float64, err error)
}

type TemperatureHumidity struct {
	TemperatureHumidityReader TemperatureHumidityReader
	TemperatureSensor         *Thermometer
	HumiditySensor            *HumiditySensor
}

func (temperatureHumidity *TemperatureHumidity) UpdateAccessoryData() (temperature float64, relativeHumidity float64, err error) {
	temperature, relativeHumidity, err = temperatureHumidity.TemperatureHumidityReader.Read()
	if err != nil {
		log.Info.Println(err)
		return
	}
	log.Debug.Printf("%f°C  %f\n", temperature, relativeHumidity)
	temperatureHumidity.TemperatureSensor.TempSensor.CurrentTemperature.SetValue(temperature)
	temperatureHumidity.HumiditySensor.HumiditySensor.CurrentRelativeHumidity.SetValue(relativeHumidity)
	return
}