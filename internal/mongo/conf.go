package mongo

var configuration Configuration

// Configuration is the configuration object configures MongoDB connection
type Configuration struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

func GetConfiguration() Configuration {
	return configuration
}

func SetConfiguration(c Configuration) {
	configuration = c
}
