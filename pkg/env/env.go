package env

import "os"

func GetEnv(envKey string, expand bool) string {
	if expand {
		return os.ExpandEnv(os.Getenv(envKey))
	}
	return os.Getenv(envKey)
}
