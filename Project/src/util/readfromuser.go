package util

import (
	"bufio"
	"os"
	"strings"
)

//reads user input
func ReadFromInput() (string, error) {
	//read
	reader := bufio.NewReader(os.Stdin)
	s, err := reader.ReadString('\n')

	return strings.TrimSpace(s), err
}
