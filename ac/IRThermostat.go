package ac

import (
	"github.com/brutella/hc/accessory"
	"github.com/brutella/hc/log"
	"os/exec"
)

type IRThermostat struct {
	remote string
	irThermostatAccessory *IRThermostatAccessory
	codeFactory CodeFactory
}

type CodeFactory func(irThermostat *IRThermostat, state int, temperature float64) *string

func NewIRThermostat(remote string, info accessory.Info, codeFactory CodeFactory, temp, min, max float64) *IRThermostat {
	irThermostat := &IRThermostat{
		remote: remote,
		irThermostatAccessory: NewIRThermostatAccessory(info, temp, min, max),
		codeFactory: codeFactory,
	}

	irThermostat.irThermostatAccessory.Thermostat.TargetHeatingCoolingState.OnValueRemoteUpdate(func(i int) {
		irThermostat.irThermostatAccessory.Thermostat.CurrentHeatingCoolingState.SetValue(i)
		irThermostat.irsendByState()
		log.Debug.Printf("%d", i)
	})

	irThermostat.irThermostatAccessory.Thermostat.TargetTemperature.OnValueRemoteUpdate(func(f float64) {
		irThermostat.irsendByState()
		log.Debug.Printf("%.0f", f)
	})

	return irThermostat
}

func (irThermostat *IRThermostat) Accessory() *accessory.Accessory {
	return irThermostat.irThermostatAccessory.Accessory
}

func (irThermostat *IRThermostat) SetCurrentTemperature(temperature float64)  {
	irThermostat.irThermostatAccessory.Thermostat.CurrentTemperature.SetValue(temperature)
}

func (irThermostat *IRThermostat) irsendByState() {
	code := irThermostat.codeFactory(irThermostat, irThermostat.irThermostatAccessory.Thermostat.CurrentHeatingCoolingState.GetValue(), irThermostat.irThermostatAccessory.Thermostat.TargetTemperature.GetValue())
	if code == nil {
		return
	}
	irThermostat.irsend(*code)
}

func (irThermostat *IRThermostat) irsend(code string) {
	log.Debug.Printf("code: %s", code)
	cmd := exec.Command("irsend", "SEND_ONCE", "shinco", code)
	err := cmd.Start()
	if err != nil {
		log.Info.Println(err)
	}
}
