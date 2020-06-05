package general

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"math/rand"
	"os"
)

func RandomBase16String(l int) string {
	buff := make([]byte, int(math.Round(float64(l)/2)))
	rand.Read(buff)
	str := hex.EncodeToString(buff)
	return str[:l] // strip 1 extra character we get from odd length results
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// CreateTmpFile creates temporary file of size l and returns a File.
// Remove and Close the file after use if you want the tmpfile to be removed.
func CreateTmpFile(l int) (*os.File, error) {
	// Open tmp file
	f, err := ioutil.TempFile("", "tmpfile")
	if err != nil {
		return f, err
	}

	// Write to tmp file
	if _, err := f.Write([]byte(RandomBase16String(l))); err != nil {
		return f, err
	}

	return f, nil
}

// RemoveFile deletes a File.
func RemoveFile(f *os.File) error {
	return os.Remove(f.Name())
}

// CloseFile closes a File.
func CloseFile(f *os.File) error {
	//if err := f.Close(); err != nil {
	//	log.Fatal(err)
	//}
	return f.Close()
}

func GetFileSize(f *os.File) {
	fi, err := f.Stat()
	check(err)
	fmt.Printf("The file is %d bytes long\n", fi.Size())
}