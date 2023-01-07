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
	data := s.ConfigurationLoader.Load()
	s.Suite.Equal("this is to test the config json loader", data)
}
