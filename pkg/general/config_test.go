// +build config_test

package general

import (
	"github.com/spf13/viper"

	"fmt"
	"os"
	"testing"
)

func TestReadConfig(t *testing.T){
	fmt.Printf("config path %s\n", os.Getenv("CONFIG_FILE"))
	viper.SetConfigType("yaml") 
	viper.SetConfigFile(os.Getenv("CONFIG_FILE"))

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	fmt.Println(viper.AllSettings())
	fmt.Printf("Project ID: %s, Bucket: %s\n", viper.GetString("gcs.PROJECT_ID"), viper.GetString("gcs.GCS_BUCKET"))
}