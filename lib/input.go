package lib

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadLines(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return []string{}, fmt.Errorf("failed to read input file: %w", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	return lines, scanner.Err()
}

func ReadFile(path string) (string, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	return strings.TrimRight(string(b), "\n"), nil
}
