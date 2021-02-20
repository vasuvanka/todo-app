package config

import (
	"os"
)

//Config - Application Configuration
type Config struct {
	DatabaseURI                  string
	Port                         string
	FEPath string
	Env string
}

//New - will return new config object
func New() *Config {
	return &Config{}
}

// Init - init config
func (c *Config) Init() {
	c.DatabaseURI = buildEnvVar("DB_URI","mongodb://localhost:27017/todos")
	c.Port = buildEnvVar("PORT", "8080")
	c.Env = buildEnvVar("GOENV","dev")
}

//getEnvVar - returns env variable value else empty string
func getEnvVar(name string) string {
	value, ok := os.LookupEnv(name)
	if !ok {
		return ""
	}
	return value
}

// buildEnvVar - build env varibales
func buildEnvVar(name, value string) string {
	envValue := getEnvVar(name)
	if envValue == "" {
		envValue = value
	}
	return envValue
}