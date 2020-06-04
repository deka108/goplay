// +build gcs_test bucket_test

package gcs

import (
	"cloud.google.com/go/storage"
	"github.com/deka108/goplay/pkg/testutil"
	"github.com/spf13/viper"
	"golang.org/x/oauth2/google"

	"bytes"
	"context"
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

func TestGetCredentials_DefaultLogin(t *testing.T){
	ctx := context.Background()
	credentials, error := google.FindDefaultCredentials(ctx, storage.ScopeReadOnly)
	if error != nil {
		fmt.Println(error)
	}
	fmt.Printf("Project ID: %s\n", credentials.ProjectID)
	fmt.Printf(string(credentials.JSON))
}


func TestGetCredentials_FromEnvironment(t *testing.T){
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", os.ExpandEnv(viper.GetString("GOOGLE_APPLICATION_CREDENTIALS")))

	ctx := context.Background()
	credentials, error := google.FindDefaultCredentials(ctx, storage.ScopeReadOnly)
	if error != nil {
		fmt.Println(error)
	}
	fmt.Printf("Project ID: %s\n", credentials.ProjectID)
	fmt.Printf("JSON: %s\n", string(credentials.JSON))
}