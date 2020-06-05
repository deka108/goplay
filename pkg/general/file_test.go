// +build file_test

package general

import (
	"github.com/deka108/goplay/pkg/testutil"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"path/filepath"
	"time"

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

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789`~!@#$%^&*()_+-=[]{}|;:',./<>?"
func randStringBytesRmndr(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63() % int64(len(letterBytes))]
	}
	return string(b)
}

func TestRandStringBytesRmndr(t *testing.T){
	defer TrackTime(time.Now(), "RandStringBytesRmndr")
	res := randStringBytesRmndr(10 * 1024)
	fmt.Println(res)
}

func TestRandomBase16String(t *testing.T){
	defer TrackTime(time.Now(), "RandomBase16String")
	fmt.Println(RandomBase16String(10 * 1024))
}

func TestTempFile(t *testing.T){
	// Open tmp file
	f, err := ioutil.TempFile("", "tmpfile")
	check(err)
	fmt.Printf("Temp file name: %s\n", f.Name())
	defer os.Remove(f.Name())

	// Check stat before write
	getFileSize(err, f)

	// Write to tmp file
	if _, err := f.Write([]byte(RandomBase16String(10 * 1024))); err != nil {
		log.Fatal(err)
	}

	// Check stat after write
	getFileSize(err, f)

	// Close tmp file
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}

func getFileSize(err error, f *os.File) {
	fi, err := f.Stat()
	check(err)
	fmt.Printf("The file is %d bytes long\n", fi.Size())
}

func TestFileJoin(t *testing.T){
	tt := assert.New(t)

	testCases := []struct {
		baseDir, file, expected string
	}{
		{"dir", "file1", "dir/file1"},
		{"dir/", "file1", "dir/file1"},
		{"dir/", "/file1", "dir/file1"},
		{"dir1/dir2", "file1", "dir1/dir2/file1"},
	}

	for _, tc := range testCases{
		tt.Equal(tc.expected, filepath.Join(tc.baseDir, tc.file))
	}
}