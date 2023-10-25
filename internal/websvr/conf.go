package websvr

var configuration *Configuration

// Configuration is the configuration file struct definition for the Gin webserver in Concourse
type Configuration struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

// SetConfiguration initialise the configuration object with provided values
func SetConfiguration(c Configuration) {
	configuration = &c
}

// GetConfiguration returns the configuration object
func GetConfiguration() Configuration {
	return *configuration
}
