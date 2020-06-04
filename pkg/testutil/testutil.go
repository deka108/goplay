package testutil

import (
	"fmt"
	"github.com/deka108/goplay/pkg/env"
	"github.com/spf13/viper"
)

func LoadConfig(){
	viper.SetConfigType("yaml") 
	viper.SetConfigFile(env.GetEnv("CONFIG_FILE", true))

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	//fmt.Printf("Project ID: %s, Bucket: %s\n", viper.GetString("gcs.PROJECT_ID"), viper.GetString("gcs.BUCKET"))
}