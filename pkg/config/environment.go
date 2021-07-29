package config

import (
	"os"
	"strconv"
)

func SetEnvironment() bool {
	_ = os.Setenv("PROD_ENVIRONMENT", "false")
	env, _ := strconv.ParseBool(os.Getenv("PROD_ENVIRONMENT"))
	return env
}
