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
	jsonfile, e := os.Open(fileName)

	if e != nil {
		return configData, e
	}
	defer jsonfile.Close()

	byteValue, err := ioutil.ReadAll(jsonfile)
	if err != nil {
		return configData, err
	}
	json.Unmarshal(byteValue, &configData)
	return configData, nil
}
