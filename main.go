package main

import (
	"MyHomeKit/ac"
	"MyHomeKit/sensor"
	"database/sql"
	"fmt"
	"github.com/brutella/hc"
	"github.com/brutella/hc/accessory"
	"github.com/brutella/hc/log"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"time"
)

func main() {
	log.Debug.Enable()

	if len(os.Args) != 2 {
		log.Info.Fatalf("Usage: %s CONFIGURATION_FILE_PATH", os.Args[0])
	}

	// Parse JSON configuration from file
	configValue, err := newFromFile(os.Args[1])
	if err != nil {
		log.Info.Fatalf("%s", err)
	}

	db, err := createDB(configValue)
	if err != nil {
		log.Info.Fatalf("%s", err)
	}
	defer db.Close()

	sht31Sensor, err := createDBSensor(db, "sht31", "SHT31")
	if err != nil {
		log.Info.Fatalf("%s", err)
	}
	dht22Sensor, err := createDBSensor(db, "dht22", "DHT22")
	if err != nil {
		log.Info.Fatalf("%s", err)
	}
	shincoIRThermostat := ac.NewIRThermostat("shinco", accessory.Info{
		Name: "Shinco",
		SerialNumber: "thermostat.shinco",
	}, ac.ShincoCodeFactory, sht31Sensor.TemperatureSensor.TempSensor.CurrentTemperature.GetValue(), 23, 25)

	config := hc.Config{Pin: configValue.Pin, Port: configValue.Port, StoragePath: configValue.StoragePath}
	t, err := hc.NewIPTransport(config,
		sht31Sensor.TemperatureSensor.Accessory,
		sht31Sensor.HumiditySensor.Accessory,
		dht22Sensor.TemperatureSensor.Accessory,
		dht22Sensor.HumiditySensor.Accessory,
		shincoIRThermostat.Accessory())

	if err != nil {
		log.Info.Fatalf("%s", err)
	}

	go func() {
		for {
			sht31Sensor.UpdateAccessoryData()
			dht22Sensor.UpdateAccessoryData()
			shincoIRThermostat.SetCurrentTemperature(sht31Sensor.TemperatureSensor.TempSensor.CurrentTemperature.GetValue())
			time.Sleep(5 * time.Second)
		}
	}()

	hc.OnTermination(func() {
		<-t.Stop()
	})

	t.Start()
}

func createDB(configValue *config) (db *sql.DB, err error) {
	db, err = sql.Open(configValue.DBDriver, configValue.dbDataSourceName())
	return
}

func createDBSensor(db *sql.DB, table string, name string) (*sensor.TemperatureHumidity, error) {
	dbSensorReader := sensor.NewDBTemperatureHumidity(db,
		fmt.Sprintf("SELECT temperature, relative_humidity FROM `%s` ORDER BY ID DESC LIMIT 1", table),
		"temperature",
		"relative_humidity")
	temperature, relativeHumidity, err := dbSensorReader.Read()

	if err != nil {
		return nil, err
	}

	temperatureSensorAccessory := sensor.NewTemperatureSensor(accessory.Info{
		Name: name,
		SerialNumber: fmt.Sprintf("sensor.%st", table),
	}, temperature, 0, 100, 0.1)
	humiditySensorAccessory := sensor.NewHumiditySensor(accessory.Info{
		Name: name,
		SerialNumber: fmt.Sprintf("sensor.%sh", table),
	}, relativeHumidity, 0, 100, 0.1)

	return &sensor.TemperatureHumidity{
		TemperatureHumidityReader: dbSensorReader,
		TemperatureSensor:         temperatureSensorAccessory,
		HumiditySensor:            humiditySensorAccessory,
	}, nil
}