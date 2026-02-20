package config

import "os"

func GetPort() string {
	port, ok := os.LookupEnv("PORT_OVERRIDE")
	if !ok {
		return ":80"
	}
	return ":" + port
}
