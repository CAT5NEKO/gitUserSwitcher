package logger

import (
	"errors"
	"fmt"
	"os"

	"github.com/fatih/color"
)

var (
	errExecutingMode = errors.New("something went wrong executing this mode 😭")
	errReadingInput  = errors.New("couldn't understand your input 🤯")
)

func PrintErrorExecutingMode() {
	printError("Error executing mode", errExecutingMode)
}

func PrintErrorReadingInput() {
	printError("Error reading input", errReadingInput)
}

func PrintError(err error) {
	printError("Error", err)
}

func PrintErrorWithMessage(err error, message string) {
	printError(message, err)
}

func printError(message string, err error) {
	fmt.Fprintf(os.Stderr, color.RedString("error: %s - %v \n", message, err))
}
