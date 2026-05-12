package errors

import (
	"fmt"
	"os"
)

func HandleError(err error) {
	if err != nil {
		fmt.Printf("error: %v", err)
		os.Exit(1)
	}
}
