package cmd

import (
	"fmt"
	"os"
)

// DieIf prints the error and exits with an error code
func DieIf(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
