package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type MyHomeKitConfig struct {
	DBDriver, DBName, DBUser, DBPassword, DBProtocol, DBAddress, DBCharset, Pin, Port, StoragePath string
}

func newFromFile(filePath string) (*MyHomeKitConfig, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var decodedConfig MyHomeKitConfig
	err = json.NewDecoder(file).Decode(&decodedConfig)
	if err != nil {
		return nil, err
	}
	return &decodedConfig, nil
}

func (configValue *MyHomeKitConfig) dbDataSourceName() string {
	return fmt.Sprintf("%s:%s@%s(%s)/%s?charset=%s", configValue.DBUser, configValue.DBPassword, configValue.DBProtocol, configValue.DBAddress, configValue.DBName, configValue.DBCharset)
}
