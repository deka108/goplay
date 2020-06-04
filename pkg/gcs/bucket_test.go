// +build gcs_test bucket_test

package gcs

import (
	"github.com/deka108/goplay/pkg/testutil"
	"github.com/spf13/viper"

	"bytes"
	"fmt"
	"log"
	"os"
	"testing"
)

type TestConfig struct {
	bucket string
	projectId string
}

var testConfig TestConfig

func TestMain(m *testing.M) {
	log.Println("Do stuff BEFORE the tests!")
	testutil.LoadConfig()
	testConfig = TestConfig{
		bucket: viper.GetString("gcs.BUCKET"),
		projectId: viper.GetString("gcs.PROJECT_ID"),
	}
	exitVal := m.Run()
	log.Println("Do stuff AFTER the tests!")

	os.Exit(exitVal)
}

func TestGetBucketMetadata(t *testing.T){
	bucketName := testConfig.bucket
	buf := new(bytes.Buffer)

	if _, err := getBucketMetadata(buf, bucketName); err != nil {
		t.Errorf("getBucketMetadata: %#v", err)
	}

	fmt.Println(buf)
}