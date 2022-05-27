package main

import (
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"
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

	_, err = tree.Tree("pkg/engine")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
