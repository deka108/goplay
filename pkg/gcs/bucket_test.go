// +build gcs_test bucket_test

package gcs

import (
	"bytes"
	"cloud.google.com/go/storage"
	"context"
	"fmt"
	. "github.com/deka108/goplay/pkg/general"
	"github.com/deka108/goplay/pkg/testutil"
	"github.com/spf13/viper"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/iterator"
	"io"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"testing"
	"time"
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
		log.Fatal(error)
	}
	fmt.Printf("Project ID: %s\n", credentials.ProjectID)
	fmt.Printf("JSON: %s\n", string(credentials.JSON))
}

func TestListObjects(t *testing.T){
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	bkt := client.Bucket(testConfig.bucket)

	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	//// with timeout
	//ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	//defer cancel()

	var names []string

	// no prefix
	query := &storage.Query{Prefix: ""}
	it := bkt.Objects(ctx, query)
	for {
		attrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		names = append(names, attrs.Name)
	}
}

func writeRandomStringToGCS(l int, objDir string){
	f, err := CreateTmpFile(l)

	//GetFileSize(f)
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(f.Name())
	defer f.Close()

	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// Upload an object with storage.Writer.
	_, fn := filepath.Split(f.Name())
	if objDir != "" {
		fn = filepath.Join(objDir, fn)
	}

	wc := client.Bucket(testConfig.bucket).Object(fn).NewWriter(ctx)
	f.Seek(0, io.SeekStart)
	if _, err = io.Copy(wc, f); err != nil {
		log.Fatal(err)
	}
	if err := wc.Close(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Blob %s uploaded.\n", fn)
}

func TestWriteLargeFile_ToGCS(t *testing.T){
	defer TrackTime(time.Now(), "Write Large to GCS")
	writeRandomStringToGCS(100*1024*1024, "dir1")
}

func TestWriteToGCS(t *testing.T){
	randGen := IntRange{0,10*1024*1024, rand.New(rand.NewSource(42))}

	objects := []struct{
		l int
		objDir string
	}{
		{
			randGen.NextRandom(), "",
		},
		{
			randGen.NextRandom(), "dir1",
		},
		{
			randGen.NextRandom(), "dir1/dir2",
		},
	}

	for _, obj := range objects {
		writeRandomStringToGCS(obj.l, obj.objDir)
	}
}