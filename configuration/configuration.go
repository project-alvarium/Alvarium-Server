package configuration

import()

type Configuration struct{
	DatabaseName string
	DatabaseURL string
}

var (
	Config Configuration
)

func initConfig()  {
	Config = Configuration{
		DatabaseName : "alvarium-db",
		DatabaseURL : "mongodb://localhost:27017", 
	}
}