package main

import (
	"fmt"
	"log"

	"github.com/bagyr/fucit/internal/issue"
	git "github.com/go-git/go-git/v5"
)

func main() {
	repo, err := git.PlainOpen("")
	if err != nil {
		log.Fatal(err)
	}

	remote, err := repo.Remote()
	remote.Fetch()

	worktree, err := repo.Worktree()
	if err != nil {
		log.Fatal(err)
	}

	iter, err := repo.Branches()
	if err != nil {
		log.Fatal(err)
	}

	for {
		r, err := iter.Next()
		if err != nil {
			log.Println(err)
			break
		}

		iss, err := issue.FromRefName(r.Name().Short())
		if err != nil {
			log.Println(err)
			continue
		}
		fmt.Println(iss)

	}
}
