// +build file_test

package general

import (
	"github.com/deka108/goplay/pkg/testutil"
	"github.com/spf13/viper"

	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"
)


func TestMain(m *testing.M) {
    log.Println("Do stuff BEFORE the tests!")
	testutil.LoadConfig()
    exitVal := m.Run()
    log.Println("Do stuff AFTER the tests!")

    os.Exit(exitVal)
}

func TestListDirectories(t *testing.T){
	files, err := ioutil.ReadDir("./")
    if err != nil {
        log.Fatal(err)
    }

    for _, f := range files {
		fmt.Println(f.Name())
	}
}

func TestGetConfig(t *testing.T){
	fmt.Printf("Project ID: %s, Bucket: %s\n", viper.GetString("gcs.PROJECT_ID"), viper.GetString("gcs.BUCKET"))
}