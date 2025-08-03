// Package cli implements logic which enables user to interact with the categorizer using command line interface.
package cli

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func Cli() (input string, err error) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter your input: ")
	if !scanner.Scan() {
		err = errors.New(fmt.Sprintf("failed to read input: %v", scanner.Err()))

	}
	input = scanner.Text()
	if strings.TrimSpace(input) == "" {
		err = errors.New("user exited")
	}
	return input, err

}
