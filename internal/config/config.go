package config

import (
	"log"
	"os/user"
	"path"

	viper "github.com/spf13/viper"
)

// GetString Wrapper for standard Viper call
func GetString(key string) string {
	return viper.GetString(key)
}

// GetHome Get current user's home directory
func GetHome() (string, error) {
	u, err := user.Current()
	return u.HomeDir, err
}

// ReadConfig Read configuration
func ReadConfig() {
	cpath, err := GetHome()
	if err != nil {
		log.Printf("Error getting home directory for current user: %s", err.Error())
	}

	viper.SetDefault("Workdir", path.Join(cpath, "testing-service"))
	viper.BindEnv("Workdir", "TESTING_SERVICE_WORKDIR")
}
