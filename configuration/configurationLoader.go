package configuration

type ConfigurationLoader interface {
	Load() string
}

type configurationLoder struct {
}

func NewConfigLoader() *configurationLoder {
	return &configurationLoder{}
}

func (c configurationLoder) Load() string {
	return "this is to test the config json loader"
}
