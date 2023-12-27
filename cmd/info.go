// since I might add support for more platforms,
// this file contains the stuff that should work
// everywhere

package cmd

import (
	"os"
	"os/user"
)

// returns the user's username
func Username() string {
	user, err := user.Current()
	if err != nil {
		return "unknown"
	}

	return user.Username
}

// returns the pc's hostname
func Hostname() string {
	hostname, err := os.Hostname()
	if err != nil {
		return "unknown"
	}

	return hostname
}
