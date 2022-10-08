package utils

import (
	"bufio"
	"fmt"
	"os"
)

func GetUsernames(filename string) ([]string, error) {
	file, err := os.Open(filename)

	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		os.Exit(1)
	}

	defer file.Close()
	var usernames []string

	read := bufio.NewScanner(file)
	for read.Scan() {
		usernames = append(usernames, read.Text())
	}
	return usernames, read.Err()
}
