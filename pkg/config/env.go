package config

import "os"

func Env() {
	os.Setenv("ENVIRONMENT", "development")
}
