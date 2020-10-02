package configuration

// Configuration holder
type Configuration struct {
	DatabaseName string
	DatabaseURL  string
	HTTPPort     string
}

// Config object
var (
	Config Configuration
)

func initConfig() {
	Config = Configuration{
		DatabaseName: "alvarium-db",
		DatabaseURL:  "mongodb://localhost:27017",
		HTTPPort:     "9090",
	}
}
