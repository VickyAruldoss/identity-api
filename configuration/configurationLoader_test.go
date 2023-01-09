package configuration

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type ConfigurationLoaderTestSuite struct {
	suite.Suite
	ConfigurationLoader
}

func TestConfigurationLoaderTestSuite(t *testing.T) {
	suite.Run(t, new(ConfigurationLoaderTestSuite))
}

func (suite *ConfigurationLoaderTestSuite) SetupTest() {
	suite.ConfigurationLoader = NewConfigLoader()
}

func (s *ConfigurationLoaderTestSuite) TestShouldCallLoadTheConfigurationFromConfigJSON() {
	expectedConfig := ConfigData{
		Environment: "DEV",
		Port:        "3000",
	}
	data, _ := s.ConfigurationLoader.Load("config.json", "")
	s.Suite.Equal(expectedConfig, data)
}

func (s *ConfigurationLoaderTestSuite) TestShouldThrowErrorWhenTheFileDoesnotExists() {
	_, err := s.ConfigurationLoader.Load("test.json", "")
	s.Suite.NotNil(err)
}
