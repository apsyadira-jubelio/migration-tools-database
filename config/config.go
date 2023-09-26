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

	err = viper.Unmarshal(&Config)
	if err != nil {
		return err
	}

	// Replace placeholders with environment variable values
	systemdbDataSource := viper.GetString("systemdb.datasource")
	tenantdbDataSource := viper.GetString("tenantdb.datasource")
	systemdbDataSource = os.ExpandEnv(systemdbDataSource) // Replace ${VAR} with the actual environment variable value
	tenantdbDataSource = os.ExpandEnv(tenantdbDataSource) // Replace ${VAR} with the actual environment variable value
	Config.System.Datasource = systemdbDataSource
	Config.Tenant.Datasource = tenantdbDataSource

	return nil
}

func Load() error {
	return LoadPath("config")
}
