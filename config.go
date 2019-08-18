package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type config struct {
	DBDriver, DBName, DBUser, DBPassword, DBProtocol, DBAddress, DBCharset, Pin, Port, StoragePath string
}

func newFromFile(filePath string) (*config, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var decodedConfig config
	err = json.NewDecoder(file).Decode(&decodedConfig)
	if err != nil {
		return nil, err
	}
	return &decodedConfig, nil
}

func (configValue *config) dbDataSourceName() string {
	return fmt.Sprintf("%s:%s@%s(%s)/%s?charset=%s", configValue.DBUser, configValue.DBPassword, configValue.DBProtocol, configValue.DBAddress, configValue.DBName, configValue.DBCharset)
}