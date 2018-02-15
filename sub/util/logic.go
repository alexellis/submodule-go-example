package util

import (
	"os"
)

// GetHostname finds the hostname of the runtime
func GetHostname() string {
	name, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	return name
}
