package config

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/viper"
)

var Config config

type config struct {
	System database `mapstructure:"systemdb"`
	Tenant database `mapstructure:"tenantdb"`
}

type database struct {
	Datasource string `mapstructure:"datasource"`
}

func LoadPath(path string) error {
	fmt.Println("CONFIG LOADED")
	/////////////////////////////
	// INITIALIZED CONFIG FILE //
	/////////////////////////////
	viper.SetConfigName("dbconfig") // name of config file (without extension)
	viper.SetConfigType("yml")      // REQUIRED if the config file does not have the extension in the name

	viper.AddConfigPath(".")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		return err
	}

	// Allow reading from environment variables
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// Read the config file
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	// Replace placeholders with environment variable values
	datasource := viper.GetString("systemdb.datasource")
	datasource = os.ExpandEnv(datasource) // Replace ${VAR} with the actual environment variable value

	println(datasource)

	err = viper.Unmarshal(&Config)
	if err != nil {
		return err
	}

	Config.System.Datasource = datasource

	return nil
}

func Load() error {
	return LoadPath("config")
}
