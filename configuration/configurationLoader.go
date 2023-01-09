package configuration

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type ConfigurationLoader interface {
	Load(fileName string, filePath string) (ConfigData, error)
}

type configurationLoder struct {
}

func NewConfigLoader() *configurationLoder {
	return &configurationLoder{}
}

func (c configurationLoder) Load(fileName string, filePath string) (ConfigData, error) {
	var configData ConfigData
	jsonfile, jsonOpenError := os.Open(fileName)

	if jsonOpenError != nil {
		return configData, jsonOpenError
	}
	defer jsonfile.Close()

	byteValue, jsonReadError := ioutil.ReadAll(jsonfile)
	if jsonReadError != nil {
		return configData, jsonReadError
	}
	json.Unmarshal(byteValue, &configData)
	return configData, nil
}
