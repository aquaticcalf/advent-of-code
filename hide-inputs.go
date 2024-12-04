package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	root_dir, err := os.Getwd()
	if err != nil {
		fmt.Println("error getting current working directory:", err)
		return
	}

	gitignore_path := filepath.Join(root_dir, ".gitignore")

	err = os.WriteFile(gitignore_path, []byte{}, 0644)
	if err != nil {
		fmt.Println("Error clearing .gitignore file:", err)
		return
	}

	gitignore_file, err := os.OpenFile(gitignore_path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("error opening .gitignore file:", err)
		return
	}
	defer gitignore_file.Close()

	err = filepath.Walk(root_dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println("error accessing path:", err)
			return err
		}

		if info.Mode().IsRegular() && info.Name() == "input.txt" {
			rel_path, err := filepath.Rel(root_dir, path)
			if err != nil {
				fmt.Println("error getting relative path:", err)
				return err
			}

			_, err = gitignore_file.WriteString(rel_path + "\n")
			if err != nil {
				fmt.Println("error writing to .gitignore file:", err)
				return err
			}

			fmt.Println("added to .gitignore:", rel_path)
		}
		return nil
	})

	if err != nil {
		fmt.Println("error walking through directories:", err)
	}
}