package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var password_store_dir = ""

func initPassStore() (string, error) {
	currentDir, err := os.Getwd()
	os.Setenv("PASSWORD_STORE_DIR", currentDir)
	fmt.Printf("passwordStore initialized to %s", currentDir)
	return currentDir, err
}

func listPassStore() {
	err := filepath.Walk(password_store_dir, func(path string, info os.FileInfo, err error) error {
		// Indent based on the depth of the file/directory
		indent := strings.Repeat("  ", strings.Count(path, "/"))

		// Print the file/directory name
		fmt.Println(indent + info.Name())

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
}

func helper() {
	fmt.Println("Usage: pass <command>\n\nAvailable commands:\n\n" +
		"    pass init\n" +
		"    pass list\n" +
		"    pass help <command>")
}
func main() {
	wordPtr := flag.String("word", "foo", "a string")
	//do things based on os args
	password_store_dir = os.Getenv("PASSWORD_STORE_DIR")
	flag.Parse()
	command := os.Args[1]
	switch command {
	case "init":
		initPassStore()
	case "list":
		listPassStore()
	case "help":
		helper()
	default:
		helper()
	}

	fmt.Println("word:", *wordPtr)
}
