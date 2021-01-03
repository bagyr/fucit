package main

import (
	"log"
	"os"

	git "github.com/go-git/go-git/v5"
)

func main() {
	_, err := git.PlainClone("", false, &git.CloneOptions{
		URL:      "https://github.com/bagyr/fucit",
		Progress: os.Stdout,
	})

	if err != nil {
		log.Fatal(err)
	}
}
