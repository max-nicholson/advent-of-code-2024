package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

const (
	BASE_URL        = "https://adventofcode.com"
	COOKIE_ENV_NAME = "ADVENT_OF_CODE_SESSION_COOKIE"
)

func main() {
	d := flag.Int("day", 0, "The day to fetch")
	flag.Parse()
	day := *d
	if day == 0 {
		log.Fatalf("--day is required")
	}
	cookie := os.Getenv(COOKIE_ENV_NAME)
	if cookie == "" {
		log.Fatalf("%s environment variable not set", COOKIE_ENV_NAME)
	}

	url := fmt.Sprintf("%s/2023/day/%d/input", BASE_URL, day)
	log.Printf("about to fetch %s", url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("unable to create request: %v", err)
	}

	req.AddCookie(&http.Cookie{Name: "session", Value: cookie})
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("failed to fetch input: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("non-OK response, got %d", resp.StatusCode)
	}

	log.Printf("successfully fetched %s", url)

	path := fmt.Sprintf("pkg/%02d/input.txt", day)
	f, err := os.Create(path)
	if err != nil {
		log.Fatalf("failed to create file at %s: %v", path, err)
	}

	defer func() {
		if err := f.Close(); err != nil {
			log.Fatalf("failed to close file at %s: %v", path, err)
		}

		log.Printf("written day %02d to %s", day, path)
	}()

	_, err = f.ReadFrom(resp.Body)
	if err != nil {
		log.Fatalf("failed to write response to file at %s: %v", path, err)
	}
}
