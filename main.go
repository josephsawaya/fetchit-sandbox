package main

import (
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

func main() {
	repo, err := git.PlainOpen("/home/joe/work/fetchit")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	ref, err := repo.Head()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	hash := ref.Hash()

	commit, err := repo.CommitObject(hash)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	tree, err := commit.Tree()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	empty := object.Tree{}

	changes, err := empty.Diff(tree)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, change := range changes {
		fmt.Printf("From: %s, To: %s\n", change.From.Name, change.To.Name)
	}

}
