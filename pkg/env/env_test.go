package env

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)


func TestEnv(t *testing.T){
	assert := assert.New(t)

	testCases := []struct{
		testFn func(string) string
		input string
		expected string
	}{
		{
			func(envKey string) string {return os.Getenv(envKey)},
			"CONFIG_FILE",
			"$HOME/configs/config-test.yml",
		},
		{
			func(envKey string) string {return os.ExpandEnv(os.Getenv(envKey))},
			"CONFIG_FILE",
			"/Users/dekauliya/configs/config-test.yml",
		},
		{
			func(envKey string) string {return GetEnv(envKey, false)},
			"CONFIG_FILE",
			"$HOME/configs/config-test.yml",
		},
		{
			func(envKey string) string {return GetEnv(envKey, true)},
			"CONFIG_FILE",
			"/Users/dekauliya/configs/config-test.yml",
		},
	}
	//fmt.Println(os.Getenv("CONFIG_FILE"))
	//fmt.Println(os.ExpandEnv(os.Getenv("CONFIG_FILE")))
	for _, tc := range testCases{
		output := tc.testFn(tc.input)
		assert.Equal(output, tc.expected)
		fmt.Println(output)
	}
}
