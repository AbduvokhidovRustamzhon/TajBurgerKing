package services

import "os"

func FlagOrEnv(flag string, envName string) (string, bool) {
	if flag != "" {
		return flag, true
	}
	return os.LookupEnv(envName)
}