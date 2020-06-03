package gcs

import (
	"github.com/joho/godotenv"

	// "bytes"
	// "fmt"
	// "io/ioutil"
	"log"
	"testing"
)

func TestGetBucketMetadata(t *testing.T){
	t.Skip()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading env-databricks.yml file")
	}
	// buf := new(bytes.Buffer)
	// bucketName := ''
	// if _, err := getBucketMetadata(buf, bucketName); err != nil {
	// 	t.Errorf("getBucketMetadata: %#v", err)
	// }
}