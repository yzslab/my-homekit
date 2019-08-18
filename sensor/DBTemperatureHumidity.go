package sensor

import "database/sql"

type DBCreator func () *sql.DB

type DBTemperatureHumidity struct {
	db *sql.DB
	selectQuery, temperatureColumn, relativeHumidityColumn string
}

func NewDBTemperatureHumidity(db *sql.DB, selectQuery, temperatureColumn, relativeHumidityColumn string) *DBTemperatureHumidity {
	return &DBTemperatureHumidity{
		db: db,
		selectQuery: selectQuery,
		temperatureColumn: temperatureColumn,
		relativeHumidityColumn: relativeHumidityColumn,
	}
}

func (dbTemperatureHumidity *DBTemperatureHumidity) Read() (temperature float64, relativeHumidity float64, err error) {
	err = dbTemperatureHumidity.db.QueryRow(dbTemperatureHumidity.selectQuery).Scan(&temperature, &relativeHumidity)
	return
}